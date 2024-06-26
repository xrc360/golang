package gcron_test

import (
	"context"
	"time"

	"github.com/xrcn/cg/os/gcron"
	"github.com/xrcn/cg/os/glog"
)

func ExampleCronAddSingleton() {
	gcron.AddSingleton(ctx, "* * * * * *", func(ctx context.Context) {
		glog.Print(context.TODO(), "doing")
		time.Sleep(2 * time.Second)
	})
	select {}
}
