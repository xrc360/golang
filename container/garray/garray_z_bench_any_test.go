package garray_test

import (
	"testing"

	"github.com/xrcn/cg/container/garray"
)

type anySortedArrayItem struct {
	priority int64
	value    interface{}
}

var (
	anyArray       = garray.NewArray()
	anySortedArray = garray.NewSortedArray(func(a, b interface{}) int {
		return int(a.(anySortedArrayItem).priority - b.(anySortedArrayItem).priority)
	})
)

func Benchmark_AnyArray_Add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		anyArray.Append(i)
	}
}

func Benchmark_AnySortedArray_Add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		anySortedArray.Add(anySortedArrayItem{
			priority: int64(i),
			value:    i,
		})
	}
}
