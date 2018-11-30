package wrapstream

import (
	"context"
	"google.golang.org/grpc"
)

//Wraps the server stream that is used for sending/receiving messages
type WrappedServerStream struct {
	grpc.ServerStream
	WrappedContext    context.Context
	recvMsgDispatch   StreamHandler
	sendMsgDispatch   StreamHandler
	recvMsgMiddleware []func(StreamHandler) StreamHandler
	sendMsgMiddleware []func(StreamHandler) StreamHandler
}

// WrapServerStream returns a ServerStream that has the ability to overwrite context.
func WrapServerStream(stream grpc.ServerStream) *WrappedServerStream {
	//If this is already a wrapped stream, just return it
	if existing, ok := stream.(*WrappedServerStream); ok {
		return existing
	}

	ctx := context.Background()
	if stream != nil {
		ctx = stream.Context()
	}

	return &WrappedServerStream{
		ServerStream:   stream,
		WrappedContext: ctx,
	}
}

//RegisterRecvMiddleware registers a middleware that will be called before each individual message is received
func (w *WrappedServerStream) RegisterRecvMiddleware(middleware func(StreamHandler) StreamHandler) {
	w.recvMsgMiddleware = append(w.recvMsgMiddleware, func(handler StreamHandler) StreamHandler {
		return outerBridge{middleware, handler}
	})

	w.recvMsgDispatch = buildChain(w.ServerStream.RecvMsg, w.recvMsgMiddleware)
}

//RegisterSendMiddleware registers a middleware that will be called before each individual message is sent
func (w *WrappedServerStream) RegisterSendMiddleware(middleware func(StreamHandler) StreamHandler) {
	w.sendMsgMiddleware = append(w.sendMsgMiddleware, func(handler StreamHandler) StreamHandler {
		return outerBridge{middleware, handler}
	})

	w.sendMsgDispatch = buildChain(w.ServerStream.SendMsg, w.sendMsgMiddleware)
}

//RegisterSendRecvMiddleware registers the middleware with the send and receive message functions
func (w *WrappedServerStream) RegisterSendRecvMiddleware(middleware func(StreamHandler) StreamHandler) {
	w.RegisterRecvMiddleware(middleware)
	w.RegisterSendMiddleware(middleware)
}

//builds the stream middleware chain
func buildChain(root func(interface{}) error, middleware []func(StreamHandler) StreamHandler) StreamHandler {
	var handler StreamHandler
	handler = StreamFunc(root)

	for i := len(middleware) - 1; i >= 0; i-- {
		handler = middleware[i](handler)
	}

	return handler
}

// Context returns the wrapper's WrappedContext, overwriting the nested grpc.ServerStream.Context()
func (w *WrappedServerStream) Context() context.Context {
	return w.WrappedContext
}

//SendMsg calls SendMsg on the underlying grpc.ServerStream but allows for middleware
func (w *WrappedServerStream) SendMsg(m interface{}) error {
	if w.sendMsgDispatch != nil {
		return w.sendMsgDispatch.Stream(m)
	}
	return w.ServerStream.SendMsg(m)
}

//RecvMsg calls RecvMsg on the underlying grpc.ServerStream but allows for middleware
func (w *WrappedServerStream) RecvMsg(m interface{}) error {
	if w.recvMsgDispatch != nil {
		return w.recvMsgDispatch.Stream(m)
	}
	return w.ServerStream.RecvMsg(m)
}

//Middleware BS

//StreamFunc implements Stream
type StreamFunc func(m interface{}) error

//Stream is the interface for Send/Recv message
func (s StreamFunc) Stream(m interface{}) error {
	return s(m)
}

//StreamHandler can stream data
type StreamHandler interface {
	Stream(m interface{}) error
}

//outerBridge that calls an inner StreamHandler, implements Stream
type outerBridge struct {
	mware func(StreamHandler) StreamHandler
	inner StreamHandler
}

func (b outerBridge) Stream(m interface{}) error {
	return b.mware(innerBridge{b.inner}).Stream(m)
}

//innerBridge that executes a stream, implements Stream
type innerBridge struct {
	inner StreamHandler
}

func (b innerBridge) Stream(m interface{}) error {
	return b.inner.Stream(m)
}
