package gutil_test

import (
	"testing"

	"github.com/xrcn/cg/frame/g"
	"github.com/xrcn/cg/test/gtest"
	"github.com/xrcn/cg/util/gutil"
)

func Test_Copy(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gutil.Copy(0), 0)
		t.Assert(gutil.Copy(1), 1)
		t.Assert(gutil.Copy("a"), "a")
		t.Assert(gutil.Copy(nil), nil)
	})
	gtest.C(t, func(t *gtest.T) {
		src := g.Map{
			"k1": "v1",
			"k2": "v2",
		}
		dst := gutil.Copy(src)
		t.Assert(dst, src)

		dst.(g.Map)["k3"] = "v3"
		t.Assert(src, g.Map{
			"k1": "v1",
			"k2": "v2",
		})
		t.Assert(dst, g.Map{
			"k1": "v1",
			"k2": "v2",
			"k3": "v3",
		})
	})
}
