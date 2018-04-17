package httpwapp

import (
	"net/http"
	"time"

	"weavelab.xyz/wlib/wapp"
	"weavelab.xyz/wlib/werror"
	"weavelab.xyz/wlib/wlog"
	"weavelab.xyz/wlib/wlog/tag"
	"weavelab.xyz/wlib/wmetrics"
)

const (
	defaultIdleTimeout = time.Second * 90
)

// Starter returns a starter which will run an HTTP server with the given handler and handle graceful shutdown for you
func Starter(addr string, h http.Handler) wapp.StartFunc {

	return func() ([]wapp.StopFunc, error) {
		wmetrics.Incr(1, "wapp", "httpwapp")

		idleTimeout := defaultIdleTimeout

		srv := &http.Server{
			Addr:        addr,
			Handler:     h,
			IdleTimeout: idleTimeout,
		}

		go func() {
			wlog.Info("starting HTTP server", tag.String("addr", addr))
			err := srv.ListenAndServe()
			if err != nil {
				wlog.WError(werror.Wrap(err, "listen and serve ended with err"))
			}
		}()

		return []wapp.StopFunc{srv.Shutdown}, nil
	}
}
