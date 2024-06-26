package gcron_test

import (
	"context"
	"testing"

	"github.com/xrcn/cg/os/gcron"
)

func Benchmark_Add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gcron.Add(ctx, "1 1 1 1 1 1", func(ctx context.Context) {

		})
	}
}
