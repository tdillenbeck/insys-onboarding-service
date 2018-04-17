package wgoji

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
)

// compile time check that responseWriter is a http.ResponseWriter
var _ = http.ResponseWriter(&responseWriter{})

// responseWriter's purpose is to capture the response status
// code and also the number of bytes written
type responseWriter struct {
	w            http.ResponseWriter
	bytesWritten int
	_statusCode  int
}

func newResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{
		w: w,
	}
}

func (s *responseWriter) Header() http.Header {
	return s.w.Header()
}

func (s *responseWriter) Write(b []byte) (int, error) {
	n, err := s.w.Write(b)
	s.bytesWritten += n

	return n, err
}

func (s *responseWriter) WriteHeader(c int) {
	s._statusCode = c
	s.w.WriteHeader(c)
}

func (s *responseWriter) statusCode() int {
	//When the status code is not explicitly set then the http package adds an http.StatusOK by default, so we do the same
	if s._statusCode == 0 {
		return http.StatusOK
	}
	return s._statusCode
}

func (s *responseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {

	var h http.Hijacker
	var ok bool
	if h, ok = s.w.(http.Hijacker); ok == false {
		return nil, nil, fmt.Errorf("response writer is not hijackable")
	}

	return h.Hijack()

}
