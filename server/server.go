package server

import (
	"context"
	"time"

	"weavelab.xyz/insys-onboarding/exampleproto"
	"weavelab.xyz/wlib/werror"
	"weavelab.xyz/wlib/wgrpc"
	"weavelab.xyz/wlib/wlog"
	"weavelab.xyz/wlib/wlog/tag"
)

type ServerImpl struct {
	Delay time.Duration
}

func (e *ServerImpl) ExampleRequest(ctx context.Context, in *exampleproto.ExampleRequestMessage) (*exampleproto.ExampleResponseMessage, error) {
	// convert wgrpcprotouuid to uuid.UUID
	someID, err := in.SomeID.UUID()
	if err != nil {
		// use wgrpc.Error to return wgrpc errors
		return nil, wgrpc.Error(wgrpc.CodeInvalidArgument, werror.Wrap(err, "error converting someID to UUID"))
	}

	wlog.Info("got example request for id", tag.String("someID", someID.String()))

	time.Sleep(e.Delay)

	return &exampleproto.ExampleResponseMessage{
		Message: "hey",
	}, nil
}
