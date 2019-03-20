package wnsq

import (
	"weavelab.xyz/monorail/shared/wlib/werror"
	"weavelab.xyz/monorail/shared/wlib/wlog"
	"weavelab.xyz/monorail/shared/wlib/wlog/tag"
)

const (
	logInfo = iota
	logWarning
	logError
)

var (
	infoLogger    = newLogger(logInfo)
	warningLogger = newLogger(logWarning)
	errorLogger   = newLogger(logError)
)

type logger struct {
	f func(string, ...tag.Tag)
}

func newLogger(level int) *logger {

	l := logger{}

	switch level {
	case logInfo, logWarning:
		l.f = wlog.Info
	case logError:
		l.f = func(msg string, tags ...tag.Tag) {
			err := werror.New(msg)
			for _, v := range tags {
				// we know that only string tags are added...
				err = err.Add(v.Key, v.StringVal)
			}
			wlog.WError(err)
		}
	}

	return &l
}

func (l *logger) Output(calldepth int, s string) error {

	l.f(s, tag.String("source", "NSQ client"))

	return nil
}
