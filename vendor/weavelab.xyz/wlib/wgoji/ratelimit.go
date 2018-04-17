package wgoji

import (
	"net/http"
	"time"

	"weavelab.xyz/go-utilities/responsewrapper"
	"weavelab.xyz/wlib/werror"
	"golang.org/x/time/rate"
)

//RateLimitMiddleware allows throttling. Requests that exceed the rate limit will receive a 429 error and the request will not be passed on
func RateLimitMiddleware(numRequests int, forEvery time.Duration) func(inner http.Handler) http.Handler {

	limiter := rate.NewLimiter(rate.Every(forEvery), numRequests)

	return func(inner http.Handler) http.Handler {
		mw := func(w http.ResponseWriter, r *http.Request) {
			if !limiter.Allow() {
				rw := responsewrapper.New(w, r)
				rw.Status(http.StatusTooManyRequests)
				rw.Error(werror.New("Exceeded rate limit"))
				return
			}

			inner.ServeHTTP(w, r)
		}

		return http.HandlerFunc(mw)
	}
}
