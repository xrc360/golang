package gtimer_test

import (
	"context"
	"fmt"
	"time"

	"github.com/xrcn/cg/os/gtimer"
)

func ExampleAdd() {
	var (
		ctx      = context.Background()
		now      = time.Now()
		interval = 1400 * time.Millisecond
	)
	gtimer.Add(ctx, interval, func(ctx context.Context) {
		fmt.Println(time.Now(), time.Duration(time.Now().UnixNano()-now.UnixNano()))
		now = time.Now()
	})

	select {}
}
