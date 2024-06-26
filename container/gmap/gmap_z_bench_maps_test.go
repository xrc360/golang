// go test *.go -bench=".*" -benchmem

package gmap_test

import (
	"testing"

	"github.com/xrcn/cg/container/gmap"
	"github.com/xrcn/cg/util/gutil"
)

var hashMap = gmap.New(true)

var listMap = gmap.NewListMap(true)

var treeMap = gmap.NewTreeMap(gutil.ComparatorInt, true)

func Benchmark_HashMap_Set(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			hashMap.Set(i, i)
			i++
		}
	})
}

func Benchmark_ListMap_Set(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			listMap.Set(i, i)
			i++
		}
	})
}

func Benchmark_TreeMap_Set(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			treeMap.Set(i, i)
			i++
		}
	})
}

func Benchmark_HashMap_Get(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			hashMap.Get(i)
			i++
		}
	})
}

func Benchmark_ListMap_Get(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			listMap.Get(i)
			i++
		}
	})
}

func Benchmark_TreeMap_Get(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			treeMap.Get(i)
			i++
		}
	})
}
