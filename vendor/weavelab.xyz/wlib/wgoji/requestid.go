package wgoji

import (
	"net/http"

	"weavelab.xyz/wlib/wcontext"
)

const requestIDHeader = "X-Weave-Request-ID"

func RequestIDMiddleware(inner http.Handler) http.Handler {

	mw := func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get(requestIDHeader)
		ctx := wcontext.NewWithRequestID(r.Context(), id)

		inner.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(mw)
}
