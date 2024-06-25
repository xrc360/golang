package gvalid_test

import (
	"testing"

	"github.com/xrc360/golang/frame/g"
	"github.com/xrc360/golang/test/gtest"
)

func Test_CI(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		err := g.Validator().Data("id").Rules("in:Id,Name").Run(ctx)
		t.AssertNE(err, nil)
	})
	gtest.C(t, func(t *gtest.T) {
		err := g.Validator().Data("id").Rules("ci|in:Id,Name").Run(ctx)
		t.AssertNil(err)
	})
	gtest.C(t, func(t *gtest.T) {
		err := g.Validator().Ci().Rules("in:Id,Name").Data("id").Run(ctx)
		t.AssertNil(err)
	})
}
