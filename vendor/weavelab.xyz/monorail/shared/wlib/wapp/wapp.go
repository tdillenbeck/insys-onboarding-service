package wapp

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"weavelab.xyz/monorail/shared/wlib/config"
	"weavelab.xyz/monorail/shared/wlib/werror"
	"weavelab.xyz/monorail/shared/wlib/wlog"
	"weavelab.xyz/monorail/shared/wlib/wlog/tag"
)

const (
	wappDebugLog = "wapp-debug-log"
)

var (
	ProbesAddr = ":45334"

	ShutdownTimeout = time.Minute
)

func init() {
	config.Add(wappDebugLog, "false", "set to true to turn on debug logging")
}

func Exit(err error) {
	wlog.WError(werror.Wrap(err, "exiting Weave app with error"))
	time.Sleep(time.Second * 2)
	os.Exit(1)
}

type SimpleStopFunc func()

func WrapSimpleStopFunc(f SimpleStopFunc) StopFunc {
	return func(_ context.Context) error {
		f()

		return nil
	}
}

type StopFunc func(context.Context) error

func Wait(ctx context.Context, fs ...StopFunc) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	select {
	case <-ctx.Done():
	case <-c:
	}

	// call the stop function in the reverse order they were added in
	for i := len(fs) - 1; i >= 0; i-- {
		v := fs[i]

		wlog.Info("shutting down", tag.Int("func", i))

		if v == nil {
			continue
		}

		ctx, done := context.WithTimeout(ctx, ShutdownTimeout)

		err := v(ctx)
		if err != nil {
			wlog.WErrorC(ctx, werror.Wrap(err, "error calling stop func"))
		}

		done()
	}

}

type Starter interface {
	Start() (StopFunc, error)
}

type StartFunc func() (StopFunc, error)

func (s StartFunc) Start() (StopFunc, error) {
	if s == nil {
		return nil, nil
	}
	return s()
}

// Up executes all Starters then waits until the app is killed with an os.Signal. If any Starter returns an error then the error is logged and exit is called.
func Up(ctx context.Context, ws ...Starter) {
	debugLog, _ := config.GetBool(wappDebugLog, false)
	if debugLog {
		wlog.SetDebugLogging(debugLog)
	}

	ws = append(ws, Probes(ProbesAddr))

	var allFS []StopFunc
	for _, w := range ws {
		if w == nil {
			continue
		}

		fs, err := w.Start()
		if err != nil {
			Exit(werror.Wrap(err, "error starting app"))
		}

		allFS = append(allFS, fs)
	}

	Wait(ctx, allFS...)
}
