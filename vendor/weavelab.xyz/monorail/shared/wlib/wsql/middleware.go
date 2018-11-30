package wsql

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"strings"
	"time"

	"weavelab.xyz/monorail/shared/wlib/wlog"
	"weavelab.xyz/monorail/shared/wlib/wlog/tag"
	"weavelab.xyz/monorail/shared/wlib/wmetrics"
)

const (
	dbCallStackMetric = "DBCallStackMetric"
)

func init() {
	wmetrics.SetLabels(dbCallStackMetric, "query", "path")

	rand.Seed(time.Now().UnixNano())
}

func (p *PG) logQuery(q string) {
	if p.LogQueries {
		wlog.Info("query", tag.String("query", q))
	}
}

func (p *PG) logQueryParameters(q string, parameters ...interface{}) {
	if p.LogQueries {
		wlog.Info("query", tag.String("query", q), tag.String("paramaters", fmt.Sprintf("%#v", parameters)))
	}
}

func (p *PG) log(caller string, q string, parameters ...interface{}) {
	if p.LogQueries {
		wlog.Info("query", tag.String("caller", caller), tag.String("query", q), tag.String("parameters", fmt.Sprintf("%#v", parameters)))
	}
}

func findStackAndStartTimer() func(...string) {
	var callerName string

	stack := make([]uintptr, 1)
	_ = runtime.Callers(4, stack)

	for _, v := range stack {
		f := runtime.FuncForPC(v - 1)
		if f != nil {
			callerName = clean(f.Name())
		}
	}

	return wmetrics.StartTimer(dbCallStackMetric, "query", callerName)
}

func (p *PG) middleware(ctx context.Context, query string, parameters ...interface{}) func(...string) {

	var callerName = CallerName(ctx)

	if callerName == "" {
		stack := make([]uintptr, 1)
		_ = runtime.Callers(4, stack)

		for _, v := range stack {
			f := runtime.FuncForPC(v - 1)
			if f != nil {
				callerName = clean(f.Name())
				break
			}
		}
	}

	// log
	p.log(callerName, query, parameters)

	// trace
	t := p.openTracingInterceptor(ctx, callerName, query)

	// metrics
	m := wmetrics.StartTimer(dbCallStackMetric, "query", callerName)

	return func(args ...string) {
		m(args...)
		t()
	}

}

//turns weavelab.xyz/data-service/services/auth.(*authSQL).GetUser --> /services/auth_authSQL_GetUser
var metricCleaner = strings.NewReplacer(
	"weavelab.xyz/", "",
	"weavelab.xyx/", "",
	"*", "",
	"(", "",
	")", "",
	".", "_",
)

func clean(s string) string {
	return metricCleaner.Replace(s)
}

var callerNameKey = "callerNameKey"

func CallerName(ctx context.Context) string {
	v := ctx.Value(callerNameKey)

	name, ok := v.(string)
	if ok {
		return name
	}

	return ""
}

func SetCallerName(ctx context.Context, name string) context.Context {
	name = clean(name)

	ctx = context.WithValue(ctx, callerNameKey, name)

	return ctx
}
