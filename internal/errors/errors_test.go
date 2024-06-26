package errors_test

import (
	"testing"

	"github.com/xrcn/cg/internal/errors"
	"github.com/xrcn/cg/test/gtest"
)

func Test_IsStackModeBrief(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(errors.IsStackModeBrief(), true)
	})
}
