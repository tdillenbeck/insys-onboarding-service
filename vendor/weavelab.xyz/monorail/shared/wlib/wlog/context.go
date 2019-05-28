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

func InfoC(c context.Context, msg string, tags ...tag.Tag) {
	defaultLogger.logMessage(c, INFO, msg, tags)
}

func DebugC(c context.Context, msg string, tags ...tag.Tag) {
	shouldPrint, ok := c.Value(debugKey).(bool)
	if defaultLogger.debugLogging == 1 || (ok && shouldPrint) {
		defaultLogger.logMessage(c, DEBUG, msg, tags)
	}
}

func ErrorC(c context.Context, msg string, tags ...tag.Tag) {
	defaultLogger.logMessage(c, ERROR, msg, tags)
}

func WErrorC(c context.Context, werr *werror.Error) {
	defaultLogger.logMessage(c, ERROR, "", []tag.Tag{tag.WError("", werr)})
}

//

func (w *WLogger) InfoC(c context.Context, msg string, tags ...tag.Tag) {
	w.logMessage(c, INFO, msg, tags)
}

func (w *WLogger) DebugC(c context.Context, msg string, tags ...tag.Tag) {
	shouldPrint, ok := c.Value(debugKey).(bool)
	if w.debugLogging == 1 || (ok && shouldPrint) {
		w.logMessage(c, DEBUG, msg, tags)
	}
}

func (w *WLogger) ErrorC(c context.Context, msg string, tags ...tag.Tag) {
	w.logMessage(c, ERROR, msg, tags)
}

func (w *WLogger) WErrorC(c context.Context, werr *werror.Error) {
	w.logMessage(c, ERROR, "", []tag.Tag{tag.WError("", werr)})
}
