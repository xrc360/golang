// go test *.go -bench=".*"

package grpool_test

import (
	"context"
	"testing"

	"github.com/xrcn/cg/os/grpool"
)

var (
	ctx = context.TODO()
	n   = 500000
)

func increment(ctx context.Context) {
	for i := 0; i < 1000000; i++ {
	}
}

func BenchmarkGrpool_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		grpool.Add(ctx, increment)
	}
}

func BenchmarkGoroutine_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go increment(ctx)
	}
}

func BenchmarkGrpool2(b *testing.B) {
	b.N = n
	for i := 0; i < b.N; i++ {
		grpool.Add(ctx, increment)
	}
}

func BenchmarkGoroutine2(b *testing.B) {
	b.N = n
	for i := 0; i < b.N; i++ {
		go increment(ctx)
	}
}
