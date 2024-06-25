package gfile_test

import (
	"os"
	"testing"
	"time"

	"github.com/xrc360/golang/os/gfile"
	"github.com/xrc360/golang/test/gtest"
)

func Test_MTime(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {

		var (
			file1   = "/testfile_t1.txt"
			err     error
			fileobj os.FileInfo
		)

		createTestFile(file1, "")
		defer delTestFiles(file1)
		fileobj, err = os.Stat(testpath() + file1)
		t.AssertNil(err)

		t.Assert(gfile.MTime(testpath()+file1), fileobj.ModTime())
		t.Assert(gfile.MTime(""), "")
	})
}

func Test_MTimeMillisecond(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			file1   = "/testfile_t1.txt"
			err     error
			fileobj os.FileInfo
		)

		createTestFile(file1, "")
		defer delTestFiles(file1)
		fileobj, err = os.Stat(testpath() + file1)
		t.AssertNil(err)

		time.Sleep(time.Millisecond * 100)
		t.AssertGE(
			gfile.MTimestampMilli(testpath()+file1),
			fileobj.ModTime().UnixNano()/1000000,
		)
		t.Assert(gfile.MTimestampMilli(""), -1)
	})
}
