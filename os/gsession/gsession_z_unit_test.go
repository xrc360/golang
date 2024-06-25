package gsession

import (
	"testing"

	"github.com/xrc360/golang/test/gtest"
)

func Test_NewSessionId(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		id1 := NewSessionId()
		id2 := NewSessionId()
		t.AssertNE(id1, id2)
		t.Assert(len(id1), 32)
	})
}
