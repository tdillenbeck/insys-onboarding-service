package mock

import (
	"context"

	"google.golang.org/grpc"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/insysproto"
)

type PortingDataClient struct {
	CreateFn       func(ctx context.Context, in *insysproto.PortingDataCreateRequest, opts []grpc.CallOption) (*insysproto.PortingDataCreateResponse, error)
	UpdateFn       func(ctx context.Context, in *insysproto.PortingDataUpdateRequest, opts []grpc.CallOption) (*insysproto.PortingDataUpdateResponse, error)
	ByLocationIDFn func(ctx context.Context, in *insysproto.PortingDataByLocationIDRequest, opts []grpc.CallOption) (*insysproto.PortingDataByLocationIDResponse, error)
	PortInStatusFn func(ctx context.Context, in *insysproto.PortInStatusRequest, opts []grpc.CallOption) (*insysproto.PortInStatusResponse, error)
	PortInCreateFn func(ctx context.Context, in *insysproto.PortInCreateRequest, opts []grpc.CallOption) (*insysproto.PortInCreateResponse, error)
}

func (pdc *PortingDataClient) Create(ctx context.Context, in *insysproto.PortingDataCreateRequest, opts ...grpc.CallOption) (*insysproto.PortingDataCreateResponse, error) {
	return pdc.CreateFn(ctx, in, opts)
}

func (pdc *PortingDataClient) Update(ctx context.Context, in *insysproto.PortingDataUpdateRequest, opts ...grpc.CallOption) (*insysproto.PortingDataUpdateResponse, error) {
	return pdc.UpdateFn(ctx, in, opts)
}

func (pdc *PortingDataClient) ByLocationID(ctx context.Context, in *insysproto.PortingDataByLocationIDRequest, opts ...grpc.CallOption) (*insysproto.PortingDataByLocationIDResponse, error) {
	return pdc.ByLocationIDFn(ctx, in, opts)
}

func (pdc *PortingDataClient) PortInStatus(ctx context.Context, in *insysproto.PortInStatusRequest, opts ...grpc.CallOption) (*insysproto.PortInStatusResponse, error) {
	return pdc.PortInStatusFn(ctx, in, opts)
}

func (pdc *PortingDataClient) PortInCreate(ctx context.Context, in *insysproto.PortInCreateRequest, opts ...grpc.CallOption) (*insysproto.PortInCreateResponse, error) {
	return pdc.PortInCreateFn(ctx, in, opts)
}
