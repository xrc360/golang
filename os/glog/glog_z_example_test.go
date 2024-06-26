package glog_test

import (
	"context"

	"github.com/xrcn/cg/frame/g"
)

func ExampleContext() {
	ctx := context.WithValue(context.Background(), "Trace-Id", "123456789")
	g.Log().Error(ctx, "runtime error")

	// May Output:
	// 2020-06-08 20:17:03.630 [ERRO] {Trace-Id: 123456789} runtime error
	// Stack:
	// ...
}
