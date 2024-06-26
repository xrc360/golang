package gvar_test

import (
	"math"
	"testing"

	"github.com/xrcn/cg/container/gvar"
	"github.com/xrcn/cg/internal/json"
	"github.com/xrcn/cg/test/gtest"
)

func TestVar_Json(t *testing.T) {
	// Marshal
	gtest.C(t, func(t *gtest.T) {
		s := "i love cg"
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
		s := "i love cg"
		v := gvar.New(nil)
		b, err := json.Marshal(s)
		t.AssertNil(err)

		err = json.UnmarshalUseNumber(b, v)
		t.AssertNil(err)
		t.Assert(v.String(), s)
	})

	gtest.C(t, func(t *gtest.T) {
		var v gvar.Var
		s := "i love cg"
		b, err := json.Marshal(s)
		t.AssertNil(err)

		err = json.UnmarshalUseNumber(b, &v)
		t.AssertNil(err)
		t.Assert(v.String(), s)
	})
}
