package gvar_test

import (
	"testing"

	"github.com/xrc360/golang/container/gvar"
	"github.com/xrc360/golang/frame/g"
	"github.com/xrc360/golang/test/gtest"
	"github.com/xrc360/golang/util/gconv"
)

func TestVar_Struct(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		type StTest struct {
			Test int
		}

		Kv := make(map[string]int, 1)
		Kv["Test"] = 100

		testObj := &StTest{}

		objOne := gvar.New(Kv, true)

		objOne.Struct(testObj)

		t.Assert(testObj.Test, Kv["Test"])
	})
	gtest.C(t, func(t *gtest.T) {
		type StTest struct {
			Test int8
		}
		o := &StTest{}
		v := gvar.New(g.Slice{"Test", "-25"})
		v.Struct(o)
		t.Assert(o.Test, -25)
	})
}

func TestVar_Var_Attribute_Struct(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Uid  int
			Name string
		}
		user := new(User)
		err := gconv.Struct(
			g.Map{
				"uid":  gvar.New(1),
				"name": gvar.New("john"),
			}, user)
		t.AssertNil(err)
		t.Assert(user.Uid, 1)
		t.Assert(user.Name, "john")
	})
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Uid  int
			Name string
		}
		var user *User
		err := gconv.Struct(
			g.Map{
				"uid":  gvar.New(1),
				"name": gvar.New("john"),
			}, &user)
		t.AssertNil(err)
		t.Assert(user.Uid, 1)
		t.Assert(user.Name, "john")
	})
}
