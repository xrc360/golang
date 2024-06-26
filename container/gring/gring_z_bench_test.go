// go test *.go -bench=".*"

package gring_test

import (
	"testing"

	"github.com/xrcn/cg/container/gring"
)

var length = 10000

var ringObject = gring.New(length, true)

func BenchmarkRing_Put(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			ringObject.Put(i)
			i++
		}
	})
}

func BenchmarkRing_Next(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			ringObject.Next()
			i++
		}
	})
}

func BenchmarkRing_Set(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			ringObject.Set(i)
			i++
		}
	})
}

func BenchmarkRing_Len(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			ringObject.Len()
			i++
		}
	})
}

func BenchmarkRing_Cap(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			ringObject.Cap()
			i++
		}
	})
}
