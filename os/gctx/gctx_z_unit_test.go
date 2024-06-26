package gctx_test

import (
	"context"
	"testing"

	"github.com/xrcn/cg/os/gctx"
	"github.com/xrcn/cg/test/gtest"
)

func Test_New(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		ctx := gctx.New()
		t.AssertNE(ctx, nil)
		t.AssertNE(gctx.CtxId(ctx), "")
	})
}

func Test_WithCtx(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		ctx := context.WithValue(context.TODO(), "TEST", 1)
		ctx = gctx.WithCtx(ctx)
		t.AssertNE(gctx.CtxId(ctx), "")
		t.Assert(ctx.Value("TEST"), 1)
	})
}

func Test_SetInitCtx(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		ctx := context.WithValue(context.TODO(), "TEST", 1)
		gctx.SetInitCtx(ctx)
		t.AssertNE(gctx.GetInitCtx(), "")
		t.Assert(gctx.GetInitCtx().Value("TEST"), 1)
	})
}
