package gclient

import (
	"context"
	"fmt"
	"github.com/xrcn/cg"
	"io"
	"net/http"
	"net/http/httptrace"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"

	"github.com/xrcn/cg/internal/httputil"
	"github.com/xrcn/cg/internal/utils"
	"github.com/xrcn/cg/net/gtrace"
	"github.com/xrcn/cg/os/gctx"
	"github.com/xrcn/cg/text/gstr"
	"github.com/xrcn/cg/util/gconv"
)

const (
	tracingInstrumentName                       = "github.com/xrcn/cg/net/gclient.Client"
	tracingAttrHttpAddressRemote                = "http.address.remote"
	tracingAttrHttpAddressLocal                 = "http.address.local"
	tracingAttrHttpDnsStart                     = "http.dns.start"
	tracingAttrHttpDnsDone                      = "http.dns.done"
	tracingAttrHttpConnectStart                 = "http.connect.start"
	tracingAttrHttpConnectDone                  = "http.connect.done"
	tracingEventHttpRequest                     = "http.request"
	tracingEventHttpRequestHeaders              = "http.request.headers"
	tracingEventHttpRequestBaggage              = "http.request.baggage"
	tracingEventHttpRequestBody                 = "http.request.body"
	tracingEventHttpResponse                    = "http.response"
	tracingEventHttpResponseHeaders             = "http.response.headers"
	tracingEventHttpResponseBody                = "http.response.body"
	tracingMiddlewareHandled        gctx.StrKey = `MiddlewareClientTracingHandled`
)

// internalMiddlewareTracing is a client middleware that enables tracing feature using standards of OpenTelemetry.
func internalMiddlewareTracing(c *Client, r *http.Request) (response *Response, err error) {
	var ctx = r.Context()
	// Mark this request is handled by server tracing middleware,
	// to avoid repeated handling by the same middleware.
	if ctx.Value(tracingMiddlewareHandled) != nil {
		return c.Next(r)
	}

	ctx = context.WithValue(ctx, tracingMiddlewareHandled, 1)
	tr := otel.GetTracerProvider().Tracer(
		tracingInstrumentName,
		trace.WithInstrumentationVersion(cg.VERSION),
	)
	ctx, span := tr.Start(ctx, r.URL.String(), trace.WithSpanKind(trace.SpanKindClient))
	defer span.End()

	span.SetAttributes(gtrace.CommonLabels()...)

	// Inject tracing content into http header.
	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(r.Header))

	// If it is now using default trace provider, it then does no complex tracing jobs.
	if gtrace.IsUsingDefaultProvider() {
		response, err = c.Next(r)
		return
	}

	// Continue client handler executing.
	response, err = c.Next(
		r.WithContext(
			httptrace.WithClientTrace(
				ctx, newClientTrace(ctx, span, r),
			),
		),
	)
	if err != nil {
		span.SetStatus(codes.Error, fmt.Sprintf(`%+v`, err))
	}
	if response == nil || response.Response == nil {
		return
	}

	reqBodyContentBytes, _ := io.ReadAll(response.Body)
	response.Body = utils.NewReadCloser(reqBodyContentBytes, false)

	span.AddEvent(tracingEventHttpResponse, trace.WithAttributes(
		attribute.String(tracingEventHttpResponseHeaders, gconv.String(httputil.HeaderToMap(response.Header))),
		attribute.String(tracingEventHttpResponseBody, gstr.StrLimit(
			string(reqBodyContentBytes),
			gtrace.MaxContentLogSize(),
			"...",
		)),
	))
	return
}
