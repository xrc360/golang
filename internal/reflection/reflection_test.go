package reflection_test

import (
	"reflect"
	"testing"

	"github.com/xrc360/golang/internal/reflection"
	"github.com/xrc360/golang/test/gtest"
)

func Test_OriginValueAndKind(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var s = "s"
		out := reflection.OriginValueAndKind(s)
		t.Assert(out.InputKind, reflect.String)
		t.Assert(out.OriginKind, reflect.String)
	})
	gtest.C(t, func(t *gtest.T) {
		var s = "s"
		out := reflection.OriginValueAndKind(&s)
		t.Assert(out.InputKind, reflect.Ptr)
		t.Assert(out.OriginKind, reflect.String)
	})
	gtest.C(t, func(t *gtest.T) {
		var s []int
		out := reflection.OriginValueAndKind(s)
		t.Assert(out.InputKind, reflect.Slice)
		t.Assert(out.OriginKind, reflect.Slice)
	})
	gtest.C(t, func(t *gtest.T) {
		var s []int
		out := reflection.OriginValueAndKind(&s)
		t.Assert(out.InputKind, reflect.Ptr)
		t.Assert(out.OriginKind, reflect.Slice)
	})
}

func Test_OriginTypeAndKind(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var s = "s"
		out := reflection.OriginTypeAndKind(s)
		t.Assert(out.InputKind, reflect.String)
		t.Assert(out.OriginKind, reflect.String)
	})
	gtest.C(t, func(t *gtest.T) {
		var s = "s"
		out := reflection.OriginTypeAndKind(&s)
		t.Assert(out.InputKind, reflect.Ptr)
		t.Assert(out.OriginKind, reflect.String)
	})
	gtest.C(t, func(t *gtest.T) {
		var s []int
		out := reflection.OriginTypeAndKind(s)
		t.Assert(out.InputKind, reflect.Slice)
		t.Assert(out.OriginKind, reflect.Slice)
	})
	gtest.C(t, func(t *gtest.T) {
		var s []int
		out := reflection.OriginTypeAndKind(&s)
		t.Assert(out.InputKind, reflect.Ptr)
		t.Assert(out.OriginKind, reflect.Slice)
	})
}
