// go test *.go -bench=".*" -benchmem

package gproc_test

import (
	"testing"

	"github.com/xrcn/cg/os/gctx"
	"github.com/xrcn/cg/os/gproc"
	"github.com/xrcn/cg/test/gtest"
)

func Test_ShellExec(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s, err := gproc.ShellExec(gctx.New(), `echo 123`)
		t.AssertNil(err)
		t.Assert(s, "123\n")
	})
	// error
	gtest.C(t, func(t *gtest.T) {
		_, err := gproc.ShellExec(gctx.New(), `NoneExistCommandCall`)
		t.AssertNE(err, nil)
	})
}
