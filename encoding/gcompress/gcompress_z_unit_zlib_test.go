package gcompress_test

import (
	"testing"

	"github.com/xrc360/golang/encoding/gcompress"
	"github.com/xrc360/golang/test/gtest"
)

func Test_Zlib_UnZlib(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		src := "hello, world\n"
		dst := []byte{120, 156, 202, 72, 205, 201, 201, 215, 81, 40, 207, 47, 202, 73, 225, 2, 4, 0, 0, 255, 255, 33, 231, 4, 147}
		data, _ := gcompress.Zlib([]byte(src))
		t.Assert(data, dst)

		data, _ = gcompress.UnZlib(dst)
		t.Assert(data, []byte(src))

		data, _ = gcompress.Zlib(nil)
		t.Assert(data, nil)
		data, _ = gcompress.UnZlib(nil)
		t.Assert(data, nil)

		data, _ = gcompress.UnZlib(dst[1:])
		t.Assert(data, nil)
	})
}
