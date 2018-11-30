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

//go:generate go get -u github.com/golang/protobuf/proto
//go:generate go get -u github.com/golang/protobuf/protoc-gen-go
//go:generate protoc --go_out=import_path=wlogproto:proto -I $GOPATH/src/weavelab.xyz/monorail/shared/wlogd/proto/ $GOPATH/src/weavelab.xyz/monorail/shared/wlogd/wlogproto/wlog.proto

// DefaultLogger is where all global wlog.Info etc. forwards to
var defaultLogger = NewWLogger(WlogdLogger)

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
)

func (m LogMsgType) String() string {
	switch m {
	case INFO:
		return "INFO"
	case ERROR:
		return "ERROR"
	case DEBUG:
		return "DEBUG"
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
	case INFO:
		prefix = "INFO "
	case ERROR:
		prefix = "ERROR"
	case DEBUG:
		prefix = "DEBUG"
	}

	parts := strings.Split(file, "/")
	if len(parts) > 0 {
		file = parts[len(parts)-1]
	}

	// WError messages
	if mtype == ERROR && msg == "" && len(tags) >= 1 && tags[0].Type == tag.WErrorType && tags[0].Key == "" {
		werr := tags[0].WErrorVal

		msg = werr.Error()
	}

	StdoutPrintHeader(prefix, time.Now(), file, line, msg)

	for _, t := range tags {
		StdoutPrintTag(t)
	}
	fmt.Println()
}

func SetDebugLogging(l bool) {
	defaultLogger.SetDebugLogging(l)
}

func SetLogHandler(h LogHandlerFunc) {
	defaultLogger.SetLogHandler(h)
}

func AddMiddleware(fn LogMiddlewareFunc) {
	defaultLogger.AddMiddleware(fn)
}

// Deprecated: should use InfoC
func Info(msg string, tags ...tag.Tag) {
	defaultLogger.logMessage(nil, INFO, msg, tags)
}

// Deprecated: should use DebugC
func Debug(msg string, tags ...tag.Tag) {
	if defaultLogger.debugLogging == 1 {
		defaultLogger.logMessage(nil, DEBUG, msg, tags)
	}
}

// Deprecated, should use WErrorC
func WError(werr *werror.Error) {
	defaultLogger.logMessage(nil, ERROR, "", []tag.Tag{tag.WError("", werr)})
}
