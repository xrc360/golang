package gconv_test

import (
	"testing"

	"github.com/xrc360/golang/test/gtest"
	"github.com/xrc360/golang/util/gconv"
)

type boolStruct struct{}

func Test_Bool(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var any interface{} = nil
		t.AssertEQ(gconv.Bool(any), false)
		t.AssertEQ(gconv.Bool(false), false)
		t.AssertEQ(gconv.Bool(nil), false)
		t.AssertEQ(gconv.Bool(0), false)
		t.AssertEQ(gconv.Bool("0"), false)
		t.AssertEQ(gconv.Bool(""), false)
		t.AssertEQ(gconv.Bool("false"), false)
		t.AssertEQ(gconv.Bool("off"), false)
		t.AssertEQ(gconv.Bool([]byte{}), false)
		t.AssertEQ(gconv.Bool([]string{}), false)
		t.AssertEQ(gconv.Bool([]interface{}{}), false)
		t.AssertEQ(gconv.Bool([]map[int]int{}), false)

		t.AssertEQ(gconv.Bool("1"), true)
		t.AssertEQ(gconv.Bool("on"), true)
		t.AssertEQ(gconv.Bool(1), true)
		t.AssertEQ(gconv.Bool(123.456), true)
		t.AssertEQ(gconv.Bool(boolStruct{}), true)
		t.AssertEQ(gconv.Bool(&boolStruct{}), true)
	})
}
