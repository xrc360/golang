package gtrace_test

import (
	"context"
	"testing"

	"github.com/xrcn/cg/net/gtrace"
	"github.com/xrcn/cg/test/gtest"
	"github.com/xrcn/cg/text/gstr"
)

func TestWithTraceID(t *testing.T) {
	var (
		ctx  = context.Background()
		uuid = `a323f910-f690-11ec-963d-79c0b7fcf119`
	)
	gtest.C(t, func(t *gtest.T) {
		newCtx, err := gtrace.WithTraceID(ctx, uuid)
		t.AssertNE(err, nil)
		t.Assert(newCtx, ctx)
	})
	gtest.C(t, func(t *gtest.T) {
		var traceId = gstr.Replace(uuid, "-", "")
		newCtx, err := gtrace.WithTraceID(ctx, traceId)
		t.AssertNil(err)
		t.AssertNE(newCtx, ctx)
		t.Assert(gtrace.GetTraceID(ctx), "")
		t.Assert(gtrace.GetTraceID(newCtx), traceId)
	})
}

func TestWithUUID(t *testing.T) {
	var (
		ctx  = context.Background()
		uuid = `a323f910-f690-11ec-963d-79c0b7fcf119`
	)
	gtest.C(t, func(t *gtest.T) {
		newCtx, err := gtrace.WithTraceID(ctx, uuid)
		t.AssertNE(err, nil)
		t.Assert(newCtx, ctx)
	})
	gtest.C(t, func(t *gtest.T) {
		newCtx, err := gtrace.WithUUID(ctx, uuid)
		t.AssertNil(err)
		t.AssertNE(newCtx, ctx)
		t.Assert(gtrace.GetTraceID(ctx), "")
		t.Assert(gtrace.GetTraceID(newCtx), gstr.Replace(uuid, "-", ""))
	})
}
