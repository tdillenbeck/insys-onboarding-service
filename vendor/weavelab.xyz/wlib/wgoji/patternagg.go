package wgoji

import (
	"context"
	"net/http"
	"strings"

	"weavelab.xyz/goji/middleware"
	"weavelab.xyz/goji/pat"
)

const (
	patternUnmatched = "/unmatched"
	patternKey       = "concatPattern"
)

//PatternAggMiddleware should be added to all root and submuxes, which makes the entire matched pattern available to the last submux.
func PatternAggMiddleware(inner http.Handler) http.Handler {

	mw := func(w http.ResponseWriter, r *http.Request) {

		ctx := withPattern(r.Context())
		inner.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(mw)

}

//withPattern adds current pattern to context so it can be used to create the whole pattern
func withPattern(ctx context.Context) context.Context {
	currPtrn := patternUnmatched

	gojiPtrn := gojiPattern(ctx)
	if gojiPtrn != "" {
		currPtrn = gojiPtrn
	}

	concatPtrn := currPtrn

	prev := ctx.Value(patternKey)
	if prevPtrn, ok := prev.(string); ok {
		concatPtrn = strings.Replace(prevPtrn, "/*", currPtrn, -1)
	}

	ctx = context.WithValue(ctx, patternKey, concatPtrn)
	return ctx
}

//gojiPattern returns what the pattern is according to goji
func gojiPattern(ctx context.Context) string {
	var currPtrn string
	patt := middleware.Pattern(ctx)
	if patt, ok := patt.(*pat.Pattern); ok {
		currPtrn = patt.String()
	}

	return currPtrn
}

//Pattern returns the aggregated pattern from the context if it exists; if not, it returns the matched pattern according to goji.
//This means this Pattern func can be used even if the PatternAggMiddleware is not being used
func Pattern(ctx context.Context) (ptrnStr string, complete bool) {
	ptrn := ctx.Value(patternKey)
	if ptrn == nil {
		gojiPtrn := gojiPattern(ctx)
		if gojiPtrn != "" {
			return gojiPtrn, false
		}

		return patternUnmatched, false
	}

	ptrnStr, ok := ptrn.(string)
	if !ok {
		return patternUnmatched, false
	}

	if strings.Contains(ptrnStr, "*") {
		return ptrnStr, false
	}

	return ptrnStr, true
}
