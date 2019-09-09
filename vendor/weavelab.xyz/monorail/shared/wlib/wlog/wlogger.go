package wlog

import (
	"context"
	"sync/atomic"
	"unsafe"

	"weavelab.xyz/monorail/shared/wlib/werror"
	"weavelab.xyz/monorail/shared/wlib/wlog/tag"
)

type WLogger struct {
	debugLogging int32
	logHandler   LogHandlerFunc

	tags map[string]tag.Tag

	// this is * because AddMiddleware uses atomic.StorePointer to set it
	logMiddleware *[]LogMiddlewareFunc
}

func NewWLogger(handler LogHandlerFunc) *WLogger {
	l := WLogger{
		logHandler:    handler,
		logMiddleware: new([]LogMiddlewareFunc),
	}

	l.AddMiddleware(TracingMiddleware)

	return &l
}

func (w *WLogger) Extend(tags ...tag.Tag) *WLogger {
	newTags := make(map[string]tag.Tag)
	for k, v := range w.tags {
		newTags[k] = v
	}

	for _, tag := range tags {
		newTags[tag.Key] = tag
	}

	sub := WLogger{
		debugLogging:  w.debugLogging,
		logHandler:    w.logHandler,
		logMiddleware: w.logMiddleware,
		tags:          newTags,
	}

	return &sub
}

func (w *WLogger) SetDebugLogging(l bool) {
	v := int32(0)
	if l {
		v = 1
	}
	atomic.StoreInt32(&w.debugLogging, v)
}

func (w *WLogger) SetLogHandler(h LogHandlerFunc) {
	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&w.logHandler)), *(*unsafe.Pointer)(unsafe.Pointer(&h)))
}

func (w *WLogger) AddMiddleware(fn LogMiddlewareFunc) {
	newMiddleware := new([]LogMiddlewareFunc)
	*newMiddleware = append(*w.logMiddleware, fn)
	// atomics are super ugly
	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&w.logMiddleware)), *(*unsafe.Pointer)(unsafe.Pointer(&newMiddleware)))
}

func (w *WLogger) callMiddlewares(c *context.Context, mtype *LogMsgType, msg *string, tags *[]tag.Tag) {
	for _, m := range *w.logMiddleware {
		m(c, mtype, msg, tags)
	}
}

func (w *WLogger) logMessage(c context.Context, mtype LogMsgType, msg string, tags []tag.Tag) {
	w.callMiddlewares(&c, &mtype, &msg, &tags)
	w.logHandler(c, mtype, msg, tags)
}

func (w *WLogger) Info(msg string, tags ...tag.Tag) {
	w.logMessage(nil, INFO, msg, append(w.flattenTags(), tags...))
}

func (w *WLogger) Debug(msg string, tags ...tag.Tag) {
	if w.debugLogging == 1 {
		w.logMessage(nil, DEBUG, msg, append(w.flattenTags(), tags...))
	}
}

func (w *WLogger) Error(msg string, tags ...tag.Tag) {
	w.logMessage(nil, ERROR, msg, append(w.flattenTags(), tags...))
}

func (w *WLogger) WError(werr *werror.Error) {
	w.logMessage(nil, ERROR, "", append([]tag.Tag{tag.WError("", werr)}, w.flattenTags()...)) // Some methods expect the error to be first
}

func (w *WLogger) flattenTags() []tag.Tag {
	flattened := make([]tag.Tag, len(w.tags))
	count := 0
	for _, v := range w.tags {
		flattened[count] = v
		count++
	}

	return flattened
}
