package errors_test

import (
	"testing"

	"github.com/xrc360/golang/internal/errors"
	"github.com/xrc360/golang/test/gtest"
)

func Test_IsStackModeBrief(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(errors.IsStackModeBrief(), true)
	})
}
