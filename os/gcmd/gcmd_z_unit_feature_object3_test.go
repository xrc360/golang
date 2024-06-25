package gcmd_test

import (
	"context"
	"os"
	"testing"

	"github.com/xrc360/golang/frame/g"
	"github.com/xrc360/golang/os/gcmd"
	"github.com/xrc360/golang/os/gctx"
	"github.com/xrc360/golang/test/gtest"
)

type TestParamsCase struct {
	g.Meta `name:"root" root:"root"`
}

type TestParamsCaseRootInput struct {
	g.Meta `name:"root"`
	Name   string
}

type TestParamsCaseRootOutput struct {
	Content string
}

func (c *TestParamsCase) Root(ctx context.Context, in TestParamsCaseRootInput) (out *TestParamsCaseRootOutput, err error) {
	out = &TestParamsCaseRootOutput{
		Content: in.Name,
	}
	return
}

func Test_Command_ParamsCase(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var ctx = gctx.New()
		cmd, err := gcmd.NewFromObject(TestParamsCase{})
		t.AssertNil(err)

		os.Args = []string{"root", "-name=john"}
		value, err := cmd.RunWithValueError(ctx)
		t.AssertNil(err)
		t.Assert(value, `{"Content":"john"}`)
	})
}
