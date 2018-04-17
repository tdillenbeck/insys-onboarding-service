package wgoji

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"weavelab.xyz/wlib/wcontext"
)

func LogMiddleware(inner http.Handler) http.Handler {

	mw := func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()
		id := wcontext.RequestID(r.Context())

		w2 := newResponseWriter(w)

		ip, _, _ := net.SplitHostPort(r.RemoteAddr)
		forwardedIPs := r.Header["X-Forwarded-For"]

		Logger(fmt.Sprintf("t=[%s] id=[%s] method=[%s] path=[%s] useragent=[%s] ip=[%s] remoteip=%v", time.Now(), id, r.Method, r.URL.Path, r.Header.Get("User-Agent"), ip, forwardedIPs))

		inner.ServeHTTP(w2, r)

		ptrn, _ := Pattern(r.Context())
		Logger(fmt.Sprintf("t=[%s] id=[%s] status=[%d] bytes=[%d] duration=[%s] pattern=[%s]", time.Now(), id, w2.statusCode(), w2.bytesWritten, time.Since(start), ptrn))
	}

	return http.HandlerFunc(mw)
}

var Logger = func(msg string) { fmt.Println(msg) }
