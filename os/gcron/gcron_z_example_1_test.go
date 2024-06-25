package gcron_test

import (
	"context"
	"time"

	"github.com/xrc360/golang/os/gcron"
	"github.com/xrc360/golang/os/glog"
)

func ExampleCronAddSingleton() {
	gcron.AddSingleton(ctx, "* * * * * *", func(ctx context.Context) {
		glog.Print(context.TODO(), "doing")
		time.Sleep(2 * time.Second)
	})
	select {}
}
