/*
	whttptracing adds middleware for distributed tracing. It adds the necessary information
	to the context, and at the end of the request, it sends the span to the tracing agent.
*/

// jaeger-debug-id: some-correlation-id

package whttptracing

import (
	"bufio"
	"net"
	"net/http"
	"net/http/httputil"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	"github.com/uber/jaeger-client-go"
	"weavelab.xyz/wlib/wcontext"
	"weavelab.xyz/wlib/werror"
	"weavelab.xyz/wlib/wgoji"
	"weavelab.xyz/wlib/wgoji/httpheaders"
	"weavelab.xyz/wlib/wtracer"
)

const (
	authorizationHeader = "Authorization"
	maxLogFieldSize     = 64000
)

func New(tracer opentracing.Tracer) (func(inner http.Handler) http.Handler, error) {

	if tracer == nil {
		var err error
		tracer, err = wtracer.DefaultTracer()
		if err != nil {
			return nil, werror.Wrap(err)
		}
	}

	m := func(inner http.Handler) http.Handler {

		mw := func(w http.ResponseWriter, r *http.Request) {

			// set the jaeger-debug header if the legacy weave debug header is set
			if d := r.Header.Get(httpheaders.DebugIDHeader); d != "" {
				r.Header.Set(jaeger.JaegerDebugHeader, d)
				r.Header.Set(wtracer.TraceBodyHeader, d)
			}

			name := "HTTP " + r.Method

			sp := opentracing.SpanFromContext(r.Context())
			if sp == nil {
				// only create a new span, iff it hasn't been created yet
				spCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
				sp = tracer.StartSpan(name, ext.RPCServerOption(spCtx))
				defer sp.Finish()

				ext.HTTPMethod.Set(sp, r.Method)
				ext.HTTPUrl.Set(sp, r.URL.String())
			}

			componentName, _ := wgoji.Pattern(r.Context())
			ext.Component.Set(sp, componentName)

			w = &statusCodeTracker{w, 200}

			r = r.WithContext(opentracing.ContextWithSpan(r.Context(), sp))

			logBodies := wtracer.ShouldLogBodies(sp.Context())
			if logBodies {

				a := r.Header.Get(authorizationHeader)
				r.Header.Del(authorizationHeader)

				includeBody := r.ContentLength < maxLogFieldSize
				out, _ := httputil.DumpRequest(r, includeBody)
				if len(out) < maxLogFieldSize {
					sp.LogFields(log.Object("HTTP.request", string(out)))
				}

				if a != "" {
					r.Header.Set(authorizationHeader, a)
				}
				// TODO: capture response body
			}

			// call the next HTTP handler in the chain
			inner.ServeHTTP(w, r)

			sp.SetTag(wtracer.RequestIDTag, wcontext.RequestID(r.Context()))

			ext.HTTPStatusCode.Set(sp, uint16(w.(*statusCodeTracker).status))

		}

		return http.HandlerFunc(mw)
	}

	return m, nil
}

type statusCodeTracker struct {
	http.ResponseWriter
	status int
}

func (w *statusCodeTracker) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *statusCodeTracker) Hijack() (net.Conn, *bufio.ReadWriter, error) {

	var h http.Hijacker
	var ok bool
	if h, ok = w.ResponseWriter.(http.Hijacker); ok == false {
		return nil, nil, werror.New("response writer is not hijackable")
	}

	return h.Hijack()

}
