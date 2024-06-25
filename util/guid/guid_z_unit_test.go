// go test *.go -bench=".*"

package guid_test

import (
	"testing"

	"github.com/xrc360/golang/container/gset"
	"github.com/xrc360/golang/test/gtest"
	"github.com/xrc360/golang/util/guid"
)

func Test_S(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		set := gset.NewStrSet()
		for i := 0; i < 1000000; i++ {
			s := guid.S()
			t.Assert(set.AddIfNotExist(s), true)
			t.Assert(len(s), 32)
		}
	})
}

func Test_S_Data(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(len(guid.S([]byte("123"))), 32)
	})
}
