package gvar_test

import (
	"math"
	"testing"

	"github.com/xrc360/golang/container/gvar"
	"github.com/xrc360/golang/internal/json"
	"github.com/xrc360/golang/test/gtest"
)

func TestVar_Json(t *testing.T) {
	// Marshal
	gtest.C(t, func(t *gtest.T) {
		s := "i love gf"
		v := gvar.New(s)
		b1, err1 := json.Marshal(v)
		b2, err2 := json.Marshal(s)
		t.Assert(err1, err2)
		t.Assert(b1, b2)
	})

	gtest.C(t, func(t *gtest.T) {
		s := int64(math.MaxInt64)
		v := gvar.New(s)
		b1, err1 := json.Marshal(v)
		b2, err2 := json.Marshal(s)
		t.Assert(err1, err2)
		t.Assert(b1, b2)
	})

	// Unmarshal
	gtest.C(t, func(t *gtest.T) {
		s := "i love gf"
		v := gvar.New(nil)
		b, err := json.Marshal(s)
		t.AssertNil(err)

		err = json.UnmarshalUseNumber(b, v)
		t.AssertNil(err)
		t.Assert(v.String(), s)
	})

	gtest.C(t, func(t *gtest.T) {
		var v gvar.Var
		s := "i love gf"
		b, err := json.Marshal(s)
		t.AssertNil(err)

		err = json.UnmarshalUseNumber(b, &v)
		t.AssertNil(err)
		t.Assert(v.String(), s)
	})
}
