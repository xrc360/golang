package gtype_test

import (
	"math"
	"sync"
	"testing"

	"github.com/xrcn/cg/container/gtype"
	"github.com/xrcn/cg/internal/json"
	"github.com/xrcn/cg/test/gtest"
	"github.com/xrcn/cg/util/gconv"
)

func Test_Uint32(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var wg sync.WaitGroup
		addTimes := 1000
		i := gtype.NewUint32(0)
		iClone := i.Clone()
		t.AssertEQ(iClone.Set(1), uint32(0))
		t.AssertEQ(iClone.Val(), uint32(1))
		for index := 0; index < addTimes; index++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				i.Add(1)
			}()
		}
		wg.Wait()
		t.AssertEQ(uint32(addTimes), i.Val())

		// empty param test
		i1 := gtype.NewUint32()
		t.AssertEQ(i1.Val(), uint32(0))

		i2 := gtype.NewUint32(11)
		t.AssertEQ(i2.Add(1), uint32(12))
		t.AssertEQ(i2.Cas(11, 13), false)
		t.AssertEQ(i2.Cas(12, 13), true)
		t.AssertEQ(i2.String(), "13")

		copyVal := i2.DeepCopy()
		i2.Set(14)
		t.AssertNE(copyVal, iClone.Val())
		i2 = nil
		copyVal = i2.DeepCopy()
		t.AssertNil(copyVal)
	})
}

func Test_Uint32_JSON(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		i := gtype.NewUint32(math.MaxUint32)
		b1, err1 := json.Marshal(i)
		b2, err2 := json.Marshal(i.Val())
		t.Assert(err1, nil)
		t.Assert(err2, nil)
		t.Assert(b1, b2)

		i2 := gtype.NewUint32()
		err := json.UnmarshalUseNumber(b2, &i2)
		t.AssertNil(err)
		t.Assert(i2.Val(), i)
	})
}

func Test_Uint32_UnmarshalValue(t *testing.T) {
	type V struct {
		Name string
		Var  *gtype.Uint32
	}
	gtest.C(t, func(t *gtest.T) {
		var v *V
		err := gconv.Struct(map[string]interface{}{
			"name": "john",
			"var":  "123",
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Var.Val(), "123")
	})
}
