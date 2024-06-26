//go:build !windows
// +build !windows

package ghttp

import (
	"context"
	"os"
	"syscall"

	"github.com/xrcn/cg/internal/intlog"
	"github.com/xrcn/cg/os/glog"
	"github.com/xrcn/cg/os/gproc"
)

// handleProcessSignal handles all signals from system in blocking way.
func handleProcessSignal() {
	var ctx = context.TODO()
	gproc.AddSigHandlerShutdown(func(sig os.Signal) {
		shutdownWebServersGracefully(ctx, sig)
	})
	gproc.AddSigHandler(func(sig os.Signal) {
		// If the graceful restart feature is not enabled,
		// it does nothing except printing a warning log.
		if !gracefulEnabled {
			glog.Warning(ctx, "graceful reload feature is disabled")
			return
		}
		if err := restartWebServers(ctx, sig, ""); err != nil {
			intlog.Errorf(ctx, `%+v`, err)
		}
	}, syscall.SIGUSR1)

	gproc.Listen()
}
