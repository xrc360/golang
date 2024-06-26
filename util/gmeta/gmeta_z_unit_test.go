package gmeta_test

import (
	"testing"

	"github.com/xrcn/cg/internal/json"
	"github.com/xrcn/cg/test/gtest"
	"github.com/xrcn/cg/util/gconv"
	"github.com/xrcn/cg/util/gmeta"
)

func TestMeta_Basic(t *testing.T) {
	type A struct {
		gmeta.Meta `tag:"123" orm:"456"`
		Id         int
		Name       string
	}

	gtest.C(t, func(t *gtest.T) {
		a := &A{
			Id:   100,
			Name: "john",
		}
		t.Assert(len(gmeta.Data(a)), 2)
		t.AssertEQ(gmeta.Get(a, "tag").String(), "123")
		t.AssertEQ(gmeta.Get(a, "orm").String(), "456")
		t.AssertEQ(gmeta.Get(a, "none"), nil)

		b, err := json.Marshal(a)
		t.AssertNil(err)
		t.Assert(b, `{"Id":100,"Name":"john"}`)
	})
}

func TestMeta_Convert_Map(t *testing.T) {
	type A struct {
		gmeta.Meta `tag:"123" orm:"456"`
		Id         int
		Name       string
	}

	gtest.C(t, func(t *gtest.T) {
		a := &A{
			Id:   100,
			Name: "john",
		}
		m := gconv.Map(a)
		t.Assert(len(m), 2)
		t.Assert(m[`Meta`], nil)
	})
}

func TestMeta_Json(t *testing.T) {
	type A struct {
		gmeta.Meta `tag:"123" orm:"456"`
		Id         int
	}

	gtest.C(t, func(t *gtest.T) {
		a := &A{
			Id: 100,
		}
		b, err := json.Marshal(a)
		t.AssertNil(err)
		t.Assert(string(b), `{"Id":100}`)
	})
}
