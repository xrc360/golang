package gproc

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/xrcn/cg/internal/intlog"
	"github.com/xrcn/cg/util/gutil"
)

// SigHandler defines a function type for signal handling.
type SigHandler func(sig os.Signal)

var (
	// Use internal variable to guarantee concurrent safety
	// when multiple Listen happen.
	signalChan        = make(chan os.Signal, 1)
	signalHandlerMu   sync.Mutex
	signalHandlerMap  = make(map[os.Signal][]SigHandler)
	shutdownSignalMap = map[os.Signal]struct{}{
		syscall.SIGINT:  {},
		syscall.SIGQUIT: {},
		syscall.SIGKILL: {},
		syscall.SIGTERM: {},
		syscall.SIGABRT: {},
	}
)

func init() {
	for sig := range shutdownSignalMap {
		signalHandlerMap[sig] = make([]SigHandler, 0)
	}
}

// AddSigHandler adds custom signal handler for custom one or more signals.
func AddSigHandler(handler SigHandler, signals ...os.Signal) {
	signalHandlerMu.Lock()
	defer signalHandlerMu.Unlock()
	for _, sig := range signals {
		signalHandlerMap[sig] = append(signalHandlerMap[sig], handler)
	}
}

// AddSigHandlerShutdown adds custom signal handler for shutdown signals:
// syscall.SIGINT,
// syscall.SIGQUIT,
// syscall.SIGKILL,
// syscall.SIGTERM,
// syscall.SIGABRT.
func AddSigHandlerShutdown(handler ...SigHandler) {
	signalHandlerMu.Lock()
	defer signalHandlerMu.Unlock()
	for _, h := range handler {
		for sig := range shutdownSignalMap {
			signalHandlerMap[sig] = append(signalHandlerMap[sig], h)
		}
	}
}

// Listen blocks and does signal listening and handling.
func Listen() {
	var (
		signals = getHandlerSignals()
		ctx     = context.Background()
		wg      = sync.WaitGroup{}
		sig     os.Signal
	)
	signal.Notify(signalChan, signals...)
	for {
		sig = <-signalChan
		intlog.Printf(ctx, `signal received: %s`, sig.String())
		if handlers := getHandlersBySignal(sig); len(handlers) > 0 {
			for _, handler := range handlers {
				wg.Add(1)
				var (
					currentHandler = handler
					currentSig     = sig
				)
				gutil.TryCatch(ctx, func(ctx context.Context) {
					defer wg.Done()
					currentHandler(currentSig)
				}, func(ctx context.Context, exception error) {
					intlog.Errorf(ctx, `execute signal handler failed: %+v`, exception)
				})
			}
		}
		// If it is shutdown signal, it exits this signal listening.
		if _, ok := shutdownSignalMap[sig]; ok {
			intlog.Printf(
				ctx,
				`receive shutdown signal "%s", waiting all signal handler done`,
				sig.String(),
			)
			// Wait until signal handlers done.
			wg.Wait()
			intlog.Print(ctx, `all signal handler done, exit process`)
			return
		}
	}
}

func getHandlerSignals() []os.Signal {
	signalHandlerMu.Lock()
	defer signalHandlerMu.Unlock()
	var signals = make([]os.Signal, 0)
	for s := range signalHandlerMap {
		signals = append(signals, s)
	}
	return signals
}

func getHandlersBySignal(sig os.Signal) []SigHandler {
	signalHandlerMu.Lock()
	defer signalHandlerMu.Unlock()
	return signalHandlerMap[sig]
}
