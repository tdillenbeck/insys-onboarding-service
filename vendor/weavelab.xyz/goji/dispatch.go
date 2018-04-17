package goji

import (
	"net/http"

	"weavelab.xyz/goji/internal"
)

type dispatch struct{}

func (d dispatch) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h := r.Context().Value(internal.Handler)
	if h == nil {
		http.NotFound(w, r)
	} else {
		h.(http.Handler).ServeHTTP(w, r)
	}
}
