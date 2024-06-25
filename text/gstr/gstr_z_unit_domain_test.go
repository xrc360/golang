// go test *.go -bench=".*"

package gstr_test

import (
	"testing"

	"github.com/xrc360/golang/test/gtest"
	"github.com/xrc360/golang/text/gstr"
)

func Test_IsSubDomain(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		main := "goframe.org"
		t.Assert(gstr.IsSubDomain("goframe.org", main), true)
		t.Assert(gstr.IsSubDomain("s.goframe.org", main), true)
		t.Assert(gstr.IsSubDomain("s.s.goframe.org", main), true)
		t.Assert(gstr.IsSubDomain("s.s.goframe.org:8080", main), true)
		t.Assert(gstr.IsSubDomain("johng.cn", main), false)
		t.Assert(gstr.IsSubDomain("s.johng.cn", main), false)
		t.Assert(gstr.IsSubDomain("s.s.johng.cn", main), false)
	})
	gtest.C(t, func(t *gtest.T) {
		main := "*.goframe.org"
		t.Assert(gstr.IsSubDomain("goframe.org", main), true)
		t.Assert(gstr.IsSubDomain("s.goframe.org", main), true)
		t.Assert(gstr.IsSubDomain("s.goframe.org:80", main), true)
		t.Assert(gstr.IsSubDomain("s.s.goframe.org", main), false)
		t.Assert(gstr.IsSubDomain("johng.cn", main), false)
		t.Assert(gstr.IsSubDomain("s.johng.cn", main), false)
		t.Assert(gstr.IsSubDomain("s.s.johng.cn", main), false)
	})
	gtest.C(t, func(t *gtest.T) {
		main := "*.*.goframe.org"
		t.Assert(gstr.IsSubDomain("goframe.org", main), true)
		t.Assert(gstr.IsSubDomain("s.goframe.org", main), true)
		t.Assert(gstr.IsSubDomain("s.s.goframe.org", main), true)
		t.Assert(gstr.IsSubDomain("s.s.goframe.org:8000", main), true)
		t.Assert(gstr.IsSubDomain("s.s.s.goframe.org", main), false)
		t.Assert(gstr.IsSubDomain("johng.cn", main), false)
		t.Assert(gstr.IsSubDomain("s.johng.cn", main), false)
		t.Assert(gstr.IsSubDomain("s.s.johng.cn", main), false)
	})
	gtest.C(t, func(t *gtest.T) {
		main := "*.*.goframe.org:8080"
		t.Assert(gstr.IsSubDomain("goframe.org", main), true)
		t.Assert(gstr.IsSubDomain("s.goframe.org", main), true)
		t.Assert(gstr.IsSubDomain("s.s.goframe.org", main), true)
		t.Assert(gstr.IsSubDomain("s.s.goframe.org:8000", main), true)
		t.Assert(gstr.IsSubDomain("s.s.s.goframe.org", main), false)
		t.Assert(gstr.IsSubDomain("johng.cn", main), false)
		t.Assert(gstr.IsSubDomain("s.johng.cn", main), false)
		t.Assert(gstr.IsSubDomain("s.s.johng.cn", main), false)
	})

	gtest.C(t, func(t *gtest.T) {
		main := "*.*.goframe.org:8080"
		t.Assert(gstr.IsSubDomain("goframe.org", main), true)
		t.Assert(gstr.IsSubDomain("s.goframe.org", main), true)
		t.Assert(gstr.IsSubDomain("s.s.goframe.org", main), true)
		t.Assert(gstr.IsSubDomain("s.s.goframe.org:8000", main), true)
		t.Assert(gstr.IsSubDomain("s.s.s.goframe.org", main), false)
		t.Assert(gstr.IsSubDomain("johng.cn", main), false)
		t.Assert(gstr.IsSubDomain("s.johng.cn", main), false)
		t.Assert(gstr.IsSubDomain("s.s.johng.cn", main), false)
	})
	gtest.C(t, func(t *gtest.T) {
		main := "s.goframe.org"
		t.Assert(gstr.IsSubDomain("goframe.org", main), false)
	})
}
