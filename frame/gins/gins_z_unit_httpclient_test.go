package gins_test

import (
	"fmt"
	"testing"

	"github.com/xrc360/golang/frame/gins"
	"github.com/xrc360/golang/test/gtest"
)

func Test_Client(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			c  = gins.HttpClient()
			c1 = gins.HttpClient("c1")
			c2 = gins.HttpClient("c2")
		)
		c.SetAgent("test1")
		c.SetAgent("test2")
		t.AssertNE(fmt.Sprintf(`%p`, c), fmt.Sprintf(`%p`, c1))
		t.AssertNE(fmt.Sprintf(`%p`, c), fmt.Sprintf(`%p`, c2))
		t.AssertNE(fmt.Sprintf(`%p`, c1), fmt.Sprintf(`%p`, c2))
	})
}
