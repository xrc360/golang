package gcron_test

import (
	"context"
	"testing"

	"github.com/xrc360/golang/os/gcron"
)

func Benchmark_Add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gcron.Add(ctx, "1 1 1 1 1 1", func(ctx context.Context) {

		})
	}
}
