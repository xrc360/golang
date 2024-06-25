package gconv_test

import (
	"testing"

	"github.com/xrc360/golang/test/gtest"
	"github.com/xrc360/golang/util/gconv"
)

func Test_Unsafe(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := "I love tt"
		t.AssertEQ(gconv.UnsafeStrToBytes(s), []byte(s))
	})

	gtest.C(t, func(t *gtest.T) {
		b := []byte("I love tt")
		t.AssertEQ(gconv.UnsafeBytesToStr(b), string(b))
	})
}
