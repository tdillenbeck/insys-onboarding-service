package wsql

import (
	"context"
	"fmt"

	"weavelab.xyz/monorail/shared/wlib/wlog"
	"weavelab.xyz/monorail/shared/wlib/wlog/tag"
)

var wsqlLogger *wlog.WLogger
func SetWsqlLogger(l *wlog.WLogger) {
	if l == nil {
		return
	}
	wsqlLogger = l
}

func init() {
	wsqlLogger = wlog.Logger()
}

type LoggerFunc func(context.Context, string, string, ...interface{})

func (p *PG) AddLogger(l LoggerFunc) {
	p.loggers = append(p.loggers, l)
}

func (p *PG) log(ctx context.Context, caller string, q string, parameters ...interface{}) {
	if p.LogQueries {
		wsqlLogger.Info("query", tag.String("caller", caller), tag.String("query", q), tag.String("parameters", fmt.Sprintf("%#v", parameters)))
	}

	for _, v := range p.loggers {
		v(ctx, caller, q, parameters)
	}

}
