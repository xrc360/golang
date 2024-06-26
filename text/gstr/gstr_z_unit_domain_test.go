// go test *.go -bench=".*"

package gstr_test

import (
	"testing"

	"github.com/xrcn/cg/test/gtest"
	"github.com/xrcn/cg/text/gstr"
)

func Test_IsSubDomain(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		main := "goxrc.org"
		t.Assert(gstr.IsSubDomain("goxrc.org", main), true)
		t.Assert(gstr.IsSubDomain("s.goxrc.org", main), true)
		t.Assert(gstr.IsSubDomain("s.s.goxrc.org", main), true)
		t.Assert(gstr.IsSubDomain("s.s.goxrc.org:8080", main), true)
		t.Assert(gstr.IsSubDomain("johng.cn", main), false)
		t.Assert(gstr.IsSubDomain("s.johng.cn", main), false)
		t.Assert(gstr.IsSubDomain("s.s.johng.cn", main), false)
	})
	gtest.C(t, func(t *gtest.T) {
		main := "*.goxrc.org"
		t.Assert(gstr.IsSubDomain("goxrc.org", main), true)
		t.Assert(gstr.IsSubDomain("s.goxrc.org", main), true)
		t.Assert(gstr.IsSubDomain("s.goxrc.org:80", main), true)
		t.Assert(gstr.IsSubDomain("s.s.goxrc.org", main), false)
		t.Assert(gstr.IsSubDomain("johng.cn", main), false)
		t.Assert(gstr.IsSubDomain("s.johng.cn", main), false)
		t.Assert(gstr.IsSubDomain("s.s.johng.cn", main), false)
	})
	gtest.C(t, func(t *gtest.T) {
		main := "*.*.goxrc.org"
		t.Assert(gstr.IsSubDomain("goxrc.org", main), true)
		t.Assert(gstr.IsSubDomain("s.goxrc.org", main), true)
		t.Assert(gstr.IsSubDomain("s.s.goxrc.org", main), true)
		t.Assert(gstr.IsSubDomain("s.s.goxrc.org:8000", main), true)
		t.Assert(gstr.IsSubDomain("s.s.s.goxrc.org", main), false)
		t.Assert(gstr.IsSubDomain("johng.cn", main), false)
		t.Assert(gstr.IsSubDomain("s.johng.cn", main), false)
		t.Assert(gstr.IsSubDomain("s.s.johng.cn", main), false)
	})
	gtest.C(t, func(t *gtest.T) {
		main := "*.*.goxrc.org:8080"
		t.Assert(gstr.IsSubDomain("goxrc.org", main), true)
		t.Assert(gstr.IsSubDomain("s.goxrc.org", main), true)
		t.Assert(gstr.IsSubDomain("s.s.goxrc.org", main), true)
		t.Assert(gstr.IsSubDomain("s.s.goxrc.org:8000", main), true)
		t.Assert(gstr.IsSubDomain("s.s.s.goxrc.org", main), false)
		t.Assert(gstr.IsSubDomain("johng.cn", main), false)
		t.Assert(gstr.IsSubDomain("s.johng.cn", main), false)
		t.Assert(gstr.IsSubDomain("s.s.johng.cn", main), false)
	})

	gtest.C(t, func(t *gtest.T) {
		main := "*.*.goxrc.org:8080"
		t.Assert(gstr.IsSubDomain("goxrc.org", main), true)
		t.Assert(gstr.IsSubDomain("s.goxrc.org", main), true)
		t.Assert(gstr.IsSubDomain("s.s.goxrc.org", main), true)
		t.Assert(gstr.IsSubDomain("s.s.goxrc.org:8000", main), true)
		t.Assert(gstr.IsSubDomain("s.s.s.goxrc.org", main), false)
		t.Assert(gstr.IsSubDomain("johng.cn", main), false)
		t.Assert(gstr.IsSubDomain("s.johng.cn", main), false)
		t.Assert(gstr.IsSubDomain("s.s.johng.cn", main), false)
	})
	gtest.C(t, func(t *gtest.T) {
		main := "s.goxrc.org"
		t.Assert(gstr.IsSubDomain("goxrc.org", main), false)
	})
}
