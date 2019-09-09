package wlog

import (
	"context"
	"fmt"
	"runtime"
	"strings"
	"time"

	"weavelab.xyz/monorail/shared/wlib/werror"
	"weavelab.xyz/monorail/shared/wlib/wlog/tag"
)

// DefaultLogger is where all global wlog.Info etc. forwards to
var currentLogger = NewWLogger(WlogdLogger)

func SetLogger(logger *WLogger) {
	currentLogger = logger
}

func Logger() *WLogger {
	return currentLogger
}

// LogHandlerFunc gets all log messages and can do what it likes with them.
// The messages are untouched at this point, so it is this function's responsibility to add
// Timestamp, Caller, HTTP Context, etc.
// The Context may be nil.
type LogHandlerFunc func(c context.Context, mtype LogMsgType, msg string, tags []tag.Tag)

type LogMiddlewareFunc func(c *context.Context, mtype *LogMsgType, msg *string, tags *[]tag.Tag)

type LogMsgType int

const (
	INFO  = LogMsgType(1)
	ERROR = LogMsgType(2)
	DEBUG = LogMsgType(3)
	TRACE = LogMsgType(4)
	WARN  = LogMsgType(5)
)

func (m LogMsgType) String() string {
	switch m {
	case ERROR:
		return "ERROR"
	case WARN:
		return "WARN"
	case INFO:
		return "INFO"
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	}
	return "UNKNOWN"
}

func StdoutLogger(c context.Context, mtype LogMsgType, msg string, tags []tag.Tag) {

	_, file, line, ok := runtime.Caller(3)
	if !ok {
		file = "???"
		line = 0
	}

	stdoutLogger(c, mtype, msg, tags, file, line)
}

// var so it can be overriden
var StdoutPrintHeader = func(prefix string, t time.Time, file string, line int, msg string) {
	hour, min, sec := t.Clock()
	fmt.Printf("%s %02d:%02d:%02d %s:%d: %s", prefix, hour, min, sec, file, line, msg)
}

var StdoutPrintTag = func(t tag.Tag) {
	if t.Type == tag.WErrorType {
		werr := t.WErrorVal

		if t.Key == "" {
			return // already handled in stdoutLogger
		}

		fmt.Printf(" %s=%s", t.Key, werr.Message())

		tags := werr.Tags()
		for key, val := range tags {
			fmt.Printf(" %s=[%v]", key, val)
		}
	} else {
		fmt.Print(" ", t.String())
	}
}

func stdoutLogger(c context.Context, mtype LogMsgType, msg string, tags []tag.Tag, file string, line int) {
	var prefix string
	switch mtype {
	case ERROR:
		prefix = "ERROR"
	case WARN:
		prefix = "WARN "
	case INFO:
		prefix = "INFO "
	case DEBUG:
		prefix = "DEBUG"
	case TRACE:
		prefix = "TRACE"
	}

	parts := strings.Split(file, "/")
	if len(parts) > 0 {
		file = parts[len(parts)-1]
	}

	// WError messages
	if mtype == ERROR && msg == "" && len(tags) >= 1 && tags[0].Type == tag.WErrorType && tags[0].Key == "" {
		werr := tags[0].WErrorVal

		m := werr.Message()
		if m != "" {
			// add extra message if it is set
			msg = m + " " + werr.Error()
		}

	}

	StdoutPrintHeader(prefix, time.Now(), file, line, msg)

	for _, t := range tags {
		StdoutPrintTag(t)
	}
	fmt.Println()
}

func SetDebugLogging(l bool) {
	currentLogger.SetDebugLogging(l)
}

func SetLogHandler(h LogHandlerFunc) {
	currentLogger.SetLogHandler(h)
}

func AddMiddleware(fn LogMiddlewareFunc) {
	currentLogger.AddMiddleware(fn)
}

// Deprecated: should use InfoC
func Info(msg string, tags ...tag.Tag) {
	currentLogger.logMessage(nil, INFO, msg, append(currentLogger.flattenTags(), tags...))
}

// Deprecated: should use DebugC
func Debug(msg string, tags ...tag.Tag) {
	if currentLogger.debugLogging == 1 {
		currentLogger.logMessage(nil, DEBUG, msg, append(currentLogger.flattenTags(), tags...))
	}
}

// Deprecated, should use WErrorC
func WError(werr *werror.Error) {
	currentLogger.logMessage(nil, ERROR, "", append([]tag.Tag{tag.WError("", werr)}, currentLogger.flattenTags()...))
}
