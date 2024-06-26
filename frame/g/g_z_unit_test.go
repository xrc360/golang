package g_test

import (
	"context"
	"os"
	"sync"
	"testing"

	"github.com/xrcn/cg/container/garray"
	"github.com/xrcn/cg/frame/g"
	"github.com/xrcn/cg/test/gtest"
	"github.com/xrcn/cg/util/gutil"
)

var (
	ctx = context.TODO()
)

func Test_NewVar(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(g.NewVar(1).Int(), 1)
		t.Assert(g.NewVar(1, true).Int(), 1)
	})
}

func Test_Dump(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		g.Dump("GoXrc")
	})
}

func Test_DumpTo(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		g.DumpTo(os.Stdout, "GoXrc", gutil.DumpOption{})
	})
}

func Test_DumpWithType(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		g.DumpWithType("GoXrc", 123)
	})
}

func Test_DumpWithOption(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		g.DumpWithOption("GoXrc", gutil.DumpOption{})
	})
}

func Test_Try(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		g.Try(ctx, func(ctx context.Context) {
			g.Dump("GoXrc")
		})
	})
}

func Test_TryCatch(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		g.TryCatch(ctx, func(ctx context.Context) {
			g.Dump("GoXrc")
		}, func(ctx context.Context, exception error) {
			g.Dump(exception)
		})
	})
	gtest.C(t, func(t *gtest.T) {
		g.TryCatch(ctx, func(ctx context.Context) {
			g.Throw("GoXrc")
		}, func(ctx context.Context, exception error) {
			t.Assert(exception.Error(), "GoXrc")
		})
	})
}

func Test_IsNil(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(g.IsNil(nil), true)
		t.Assert(g.IsNil(0), false)
		t.Assert(g.IsNil("GoXrc"), false)
	})
}

func Test_IsEmpty(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(g.IsEmpty(nil), true)
		t.Assert(g.IsEmpty(0), true)
		t.Assert(g.IsEmpty("GoXrc"), false)
	})
}

func Test_SetDebug(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		g.SetDebug(true)
	})
}

func Test_Object(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertNE(g.Client(), nil)
		t.AssertNE(g.Server(), nil)
		t.AssertNE(g.TCPServer(), nil)
		t.AssertNE(g.UDPServer(), nil)
		t.AssertNE(g.View(), nil)
		t.AssertNE(g.Config(), nil)
		t.AssertNE(g.Cfg(), nil)
		t.AssertNE(g.Resource(), nil)
		t.AssertNE(g.I18n(), nil)
		t.AssertNE(g.Res(), nil)
		t.AssertNE(g.Log(), nil)
		t.AssertNE(g.Validator(), nil)
	})
}

func Test_Go(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			wg    = sync.WaitGroup{}
			array = garray.NewArray(true)
		)
		wg.Add(1)
		g.Go(context.Background(), func(ctx context.Context) {
			defer wg.Done()
			array.Append(1)
		}, nil)
		wg.Wait()
		t.Assert(array.Len(), 1)
	})
}
