// This file is for logging with a Google Context
package wlog

import (
	"context"

	"weavelab.xyz/monorail/shared/wlib/werror"
	"weavelab.xyz/monorail/shared/wlib/wlog/tag"
)

type key int

const debugKey key = 1

func WithDebugPrint(c context.Context, printDebug bool) context.Context {
	return context.WithValue(c, debugKey, printDebug)
}

func TraceC(c context.Context, msg string, tags ...tag.Tag) {
	shouldPrint, ok := c.Value(debugKey).(bool)
	if currentLogger.debugLogging == 1 || (ok && shouldPrint) {
		currentLogger.logMessage(c, TRACE, msg, tags)
	}
}

func DebugC(c context.Context, msg string, tags ...tag.Tag) {
	shouldPrint, ok := c.Value(debugKey).(bool)
	if currentLogger.debugLogging == 1 || (ok && shouldPrint) {
		currentLogger.logMessage(c, DEBUG, msg, tags)
	}
}

func InfoC(c context.Context, msg string, tags ...tag.Tag) {
	currentLogger.logMessage(c, INFO, msg, tags)
}

func WarnC(c context.Context, msg string, tags ...tag.Tag) {
	currentLogger.logMessage(c, WARN, msg, tags)
}

func ErrorC(c context.Context, msg string, tags ...tag.Tag) {
	currentLogger.logMessage(c, ERROR, msg, tags)
}

func WErrorC(c context.Context, werr *werror.Error) {
	currentLogger.logMessage(c, ERROR, "", []tag.Tag{tag.WError("", werr)})
}

//

func (w *WLogger) TraceC(c context.Context, msg string, tags ...tag.Tag) {
	shouldPrint, ok := c.Value(debugKey).(bool)
	if w.debugLogging == 1 || (ok && shouldPrint) {
		w.logMessage(c, TRACE, msg, append(w.flattenTags(), tags...))
	}
}

func (w *WLogger) DebugC(c context.Context, msg string, tags ...tag.Tag) {
	shouldPrint, ok := c.Value(debugKey).(bool)
	if w.debugLogging == 1 || (ok && shouldPrint) {
		w.logMessage(c, DEBUG, msg, append(w.flattenTags(), tags...))
	}
}

func (w *WLogger) InfoC(c context.Context, msg string, tags ...tag.Tag) {
	w.logMessage(c, INFO, msg, append(w.flattenTags(), tags...))
}

func (w *WLogger) WarnC(c context.Context, msg string, tags ...tag.Tag) {
	w.logMessage(c, WARN, msg, append(w.flattenTags(), tags...))
}

func (w *WLogger) ErrorC(c context.Context, msg string, tags ...tag.Tag) {
	w.logMessage(c, ERROR, msg, append(w.flattenTags(), tags...))
}

func (w *WLogger) WErrorC(c context.Context, werr *werror.Error) {
	w.logMessage(c, ERROR, "", append([]tag.Tag{tag.WError("", werr)}, w.flattenTags()...))
}
