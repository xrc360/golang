// go test *.go -bench=".*" -benchmem

package gmap_test

import (
	"sync"
	"testing"

	"github.com/xrcn/cg/container/gmap"
)

var gm = gmap.NewIntIntMap(true)

var sm = sync.Map{}

func Benchmark_GMapSet(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			gm.Set(i, i)
			i++
		}
	})
}

func Benchmark_SyncMapSet(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			sm.Store(i, i)
			i++
		}
	})
}

func Benchmark_GMapGet(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			gm.Get(i)
			i++
		}
	})
}

func Benchmark_SyncMapGet(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			sm.Load(i)
			i++
		}
	})
}

func Benchmark_GMapRemove(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			gm.Remove(i)
			i++
		}
	})
}

func Benchmark_SyncMapRmove(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			sm.Delete(i)
			i++
		}
	})
}
