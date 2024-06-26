package gutil_test

import (
	"context"
	"sync"
	"testing"

	"github.com/xrcn/cg/container/garray"
	"github.com/xrcn/cg/test/gtest"
	"github.com/xrcn/cg/util/gutil"
)

func Test_Go(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			wg    = sync.WaitGroup{}
			array = garray.NewArray(true)
		)
		wg.Add(1)
		gutil.Go(ctx, func(ctx context.Context) {
			defer wg.Done()
			array.Append(1)
		}, nil)
		wg.Wait()
		t.Assert(array.Len(), 1)
	})
	// recover
	gtest.C(t, func(t *gtest.T) {
		var (
			wg    = sync.WaitGroup{}
			array = garray.NewArray(true)
		)
		wg.Add(1)
		gutil.Go(ctx, func(ctx context.Context) {
			defer wg.Done()
			panic("error")
			array.Append(1)
		}, nil)
		wg.Wait()
		t.Assert(array.Len(), 0)
	})
	gtest.C(t, func(t *gtest.T) {
		var (
			wg    = sync.WaitGroup{}
			array = garray.NewArray(true)
		)
		wg.Add(1)
		gutil.Go(ctx, func(ctx context.Context) {
			panic("error")
		}, func(ctx context.Context, exception error) {
			defer wg.Done()
			array.Append(exception)
		})
		wg.Wait()
		t.Assert(array.Len(), 1)
		t.Assert(array.At(0), "error")
	})
}
