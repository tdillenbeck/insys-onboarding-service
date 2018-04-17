package wgoji

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"weavelab.xyz/wlib/wmetrics"
	"weavelab.xyz/wlib/wmetrics/client"
)

const httpStatsPrefix = "http"

var (
	httpLabels     = []string{"method", "pattern", "code"}
	WMetricsClient client.WMetrics
)

func init() {
	//WMetricsClient that the stats middleware will use to send stats
	WMetricsClient = wmetrics.DefaultClient

	WMetricsClient.SetLabels(httpStatsPrefix, httpLabels...)
}

// StatsMiddleware collects standard information about the request and dumps
// it into statsdaemon. It tracks request count, request time, response codes,
// and response size (if content-size header exists) by matched endpoint.
func StatsMiddleware(inner http.Handler) http.Handler {
	mw := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		statsW := newResponseWriter(w)

		ctx := withPattern(r.Context())
		inner.ServeHTTP(statsW, r.WithContext(ctx))

		s := stats(ctx, start, statsW)
		sendStats(r.Method, s)
	}

	return http.HandlerFunc(mw)
}

func sendStats(method string, s responseStats) {
	basePath := fmt.Sprintf("%s.%s.%s", method, s.urlPattern, strconv.Itoa(s.code))

	WMetricsClient.Time(s.duration, httpStatsPrefix, basePath)
}

type responseStats struct {
	duration   time.Duration
	urlPattern string

	code int
	size int
}

func stats(ctx context.Context, start time.Time, w *responseWriter) responseStats {

	var s responseStats

	s.duration = time.Since(start)
	s.urlPattern = matchedPattern(ctx)

	s.code = w.statusCode()
	s.size, _ = strconv.Atoi(w.Header().Get("Content-Length"))
	if s.size == 0 {
		s.size = w.bytesWritten
	}

	return s
}

var patternReplacer = strings.NewReplacer(
	".", "_",
)

func matchedPattern(ctx context.Context) string {
	ptrn, _ := Pattern(ctx)

	ptrn = patternReplacer.Replace(ptrn)

	return ptrn
}
