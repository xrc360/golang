package gvalid_test

import (
	"context"
	"testing"

	"github.com/xrcn/cg/errors/gerror"
	"github.com/xrcn/cg/frame/g"
	"github.com/xrcn/cg/test/gtest"
	"github.com/xrcn/cg/util/gvalid"
)

type UserCreateReq struct {
	g.Meta `v:"UserCreateReq"`
	Name   string
	Pass   string
}

func RuleUserCreateReq(ctx context.Context, in gvalid.RuleFuncInput) error {
	var req *UserCreateReq
	if err := in.Data.Scan(&req); err != nil {
		return gerror.Wrap(err, `Scan data to UserCreateReq failed`)
	}
	return gerror.Newf(`The name "%s" is already token by others`, req.Name)
}

func Test_Meta(t *testing.T) {
	var user = &UserCreateReq{
		Name: "john",
		Pass: "123456",
	}

	gtest.C(t, func(t *gtest.T) {
		err := g.Validator().RuleFunc("UserCreateReq", RuleUserCreateReq).
			Data(user).
			Assoc(g.Map{
				"Name": "john smith",
			}).Run(ctx)
		t.Assert(err.String(), `The name "john smith" is already token by others`)
	})
}
