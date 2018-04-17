package wgoji

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"weavelab.xyz/go-utilities/responsewrapper"
	"weavelab.xyz/wlib/wgoji/httpheaders"
)

func DeadlineMiddleware(defaultDeadline time.Duration, serviceRequiredTime time.Duration) func(inner http.Handler) http.Handler {

	return func(inner http.Handler) http.Handler {
		mw := func(w http.ResponseWriter, r *http.Request) {

			// get the deadline from the request
			deadline := requestDeadline(defaultDeadline, serviceRequiredTime, r)

			// add the deadline to the context
			// ignore the returned cancel function
			ctx, cancel := context.WithDeadline(r.Context(), deadline)
			defer cancel()

			// make sure we can service the request before the deadline
			// see if we have enough time to complete the request
			select {
			// The returned context's Done channel is closed
			// when the deadline expires
			case <-ctx.Done():
				rw := responsewrapper.New(w, r)
				rw.Status(http.StatusRequestTimeout)
				rw.Error(fmt.Errorf("unable to complete request due to deadline"))
				return
			default:
			}

			inner.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(mw)
	}
}

func requestDeadline(d time.Duration, buffer time.Duration, r *http.Request) time.Time {
	deadline := time.Now().Add(d)

	s := r.Header.Get(httpheaders.DeadlineHeader)
	if s != "" {
		t, err := time.Parse(time.RFC3339, s)
		if err != nil {
			// TODO: report the error
		} else {
			deadline = t
		}
	}

	deadline = deadline.Add(-buffer)

	return deadline
}
