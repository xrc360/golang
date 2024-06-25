package glog_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/xrc360/golang/container/garray"
	"github.com/xrc360/golang/os/glog"
	"github.com/xrc360/golang/test/gtest"
	"github.com/xrc360/golang/text/gstr"
)

var arrayForHandlerTest1 = garray.NewStrArray()

func customHandler1(ctx context.Context, input *glog.HandlerInput) {
	arrayForHandlerTest1.Append(input.String(false))
	input.Next(ctx)
}

func TestLogger_SetHandlers1(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		w := bytes.NewBuffer(nil)
		l := glog.NewWithWriter(w)
		l.SetHandlers(customHandler1)
		l.SetCtxKeys("Trace-Id", "Span-Id", "Test")
		ctx := context.WithValue(context.Background(), "Trace-Id", "1234567890")
		ctx = context.WithValue(ctx, "Span-Id", "abcdefg")

		l.Print(ctx, 1, 2, 3)
		t.Assert(gstr.Count(w.String(), "1234567890"), 1)
		t.Assert(gstr.Count(w.String(), "abcdefg"), 1)
		t.Assert(gstr.Count(w.String(), "1 2 3"), 1)

		t.Assert(arrayForHandlerTest1.Len(), 1)
		t.Assert(gstr.Count(arrayForHandlerTest1.At(0), "1234567890"), 1)
		t.Assert(gstr.Count(arrayForHandlerTest1.At(0), "abcdefg"), 1)
		t.Assert(gstr.Count(arrayForHandlerTest1.At(0), "1 2 3"), 1)
	})
}

var arrayForHandlerTest2 = garray.NewStrArray()

func customHandler2(ctx context.Context, input *glog.HandlerInput) {
	arrayForHandlerTest2.Append(input.String(false))
}

func TestLogger_SetHandlers2(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		w := bytes.NewBuffer(nil)
		l := glog.NewWithWriter(w)
		l.SetHandlers(customHandler2)
		l.SetCtxKeys("Trace-Id", "Span-Id", "Test")
		ctx := context.WithValue(context.Background(), "Trace-Id", "1234567890")
		ctx = context.WithValue(ctx, "Span-Id", "abcdefg")

		l.Print(ctx, 1, 2, 3)
		t.Assert(gstr.Count(w.String(), "1234567890"), 0)
		t.Assert(gstr.Count(w.String(), "abcdefg"), 0)
		t.Assert(gstr.Count(w.String(), "1 2 3"), 0)

		t.Assert(arrayForHandlerTest2.Len(), 1)
		t.Assert(gstr.Count(arrayForHandlerTest2.At(0), "1234567890"), 1)
		t.Assert(gstr.Count(arrayForHandlerTest2.At(0), "abcdefg"), 1)
		t.Assert(gstr.Count(arrayForHandlerTest2.At(0), "1 2 3"), 1)
	})
}

func TestLogger_SetHandlers_HandlerJson(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		w := bytes.NewBuffer(nil)
		l := glog.NewWithWriter(w)
		l.SetHandlers(glog.HandlerJson)
		l.SetCtxKeys("Trace-Id", "Span-Id", "Test")
		ctx := context.WithValue(context.Background(), "Trace-Id", "1234567890")
		ctx = context.WithValue(ctx, "Span-Id", "abcdefg")

		l.Debug(ctx, 1, 2, 3)
		t.Assert(gstr.Count(w.String(), "1234567890"), 1)
		t.Assert(gstr.Count(w.String(), "abcdefg"), 1)
		t.Assert(gstr.Count(w.String(), `"1 2 3"`), 1)
		t.Assert(gstr.Count(w.String(), `"DEBU"`), 1)
	})
}

func TestLogger_SetHandlers_HandlerStructure(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		w := bytes.NewBuffer(nil)
		l := glog.NewWithWriter(w)
		l.SetHandlers(glog.HandlerStructure)
		l.SetCtxKeys("Trace-Id", "Span-Id", "Test")
		ctx := context.WithValue(context.Background(), "Trace-Id", "1234567890")
		ctx = context.WithValue(ctx, "Span-Id", "abcdefg")

		l.Debug(ctx, "debug", "uid", 1000)
		l.Info(ctx, "info", "' '", `"\n`)

		t.Assert(gstr.Count(w.String(), "uid=1000"), 1)
		t.Assert(gstr.Count(w.String(), "Content=debug"), 1)
		t.Assert(gstr.Count(w.String(), `"' '"="\"\\n"`), 1)
		t.Assert(gstr.Count(w.String(), `CtxStr="1234567890, abcdefg"`), 2)
	})
}

func Test_SetDefaultHandler(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		oldHandler := glog.GetDefaultHandler()
		glog.SetDefaultHandler(func(ctx context.Context, in *glog.HandlerInput) {
			glog.HandlerJson(ctx, in)
		})
		defer glog.SetDefaultHandler(oldHandler)

		w := bytes.NewBuffer(nil)
		l := glog.NewWithWriter(w)
		l.SetCtxKeys("Trace-Id", "Span-Id", "Test")
		ctx := context.WithValue(context.Background(), "Trace-Id", "1234567890")
		ctx = context.WithValue(ctx, "Span-Id", "abcdefg")

		l.Debug(ctx, 1, 2, 3)
		t.Assert(gstr.Count(w.String(), "1234567890"), 1)
		t.Assert(gstr.Count(w.String(), "abcdefg"), 1)
		t.Assert(gstr.Count(w.String(), `"1 2 3"`), 1)
		t.Assert(gstr.Count(w.String(), `"DEBU"`), 1)
	})
}
