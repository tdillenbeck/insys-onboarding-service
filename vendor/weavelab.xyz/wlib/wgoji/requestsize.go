package wgoji

import "net/http"

//LimitRequestSizeMiddleware ensures that a request's body will not exceed the given number of bytes. If it does, an error will be returned when it is read.
func LimitRequestSizeMiddleware(numBytes int64) func(inner http.Handler) http.Handler {

	return func(inner http.Handler) http.Handler {
		mw := func(w http.ResponseWriter, r *http.Request) {
			//Replace the default reader with one that will return an error if it gets to the maximum number of bytes
			r.Body = http.MaxBytesReader(w, r.Body, numBytes)

			inner.ServeHTTP(w, r)
		}

		return http.HandlerFunc(mw)
	}
}
