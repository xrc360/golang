package gtrace_test

import (
	"context"
	"testing"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"

	"github.com/xrc360/golang/net/gtrace"
	"github.com/xrc360/golang/test/gtest"
)

const (
	traceIDStr = "4bf92f3577b34da6a3ce929d0e0e4736"
	spanIDStr  = "00f067aa0ba902b7"
)

var (
	traceID = mustTraceIDFromHex(traceIDStr)
	spanID  = mustSpanIDFromHex(spanIDStr)
)

func mustTraceIDFromHex(s string) (t trace.TraceID) {
	var err error
	t, err = trace.TraceIDFromHex(s)
	if err != nil {
		panic(err)
	}
	return
}

func mustSpanIDFromHex(s string) (t trace.SpanID) {
	var err error
	t, err = trace.SpanIDFromHex(s)
	if err != nil {
		panic(err)
	}
	return
}

func TestNewCarrier(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		ctx := trace.ContextWithRemoteSpanContext(context.Background(), trace.NewSpanContext(trace.SpanContextConfig{
			TraceID:    traceID,
			SpanID:     spanID,
			TraceFlags: trace.FlagsSampled,
		}))
		sc := trace.SpanContextFromContext(ctx)
		t.Assert(sc.TraceID().String(), traceID.String())
		t.Assert(sc.SpanID().String(), "00f067aa0ba902b7")

		ctx, _ = otel.Tracer("").Start(ctx, "inject")
		carrier1 := gtrace.NewCarrier()
		otel.GetTextMapPropagator().Inject(ctx, carrier1)

		ctx = otel.GetTextMapPropagator().Extract(ctx, carrier1)
		gotSc := trace.SpanContextFromContext(ctx)
		t.Assert(gotSc.TraceID().String(), traceID.String())
		// New span is created internally, so the SpanID is different.
	})
}
