package mock

import (
	"context"

	"google.golang.org/grpc"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/insysproto"
)

type PortingDataClient struct {
	ByLocationIDFn                    func(ctx context.Context, in *insysproto.PortingDataByLocationIDRequest, opts []grpc.CallOption) (*insysproto.PortingDataByLocationIDResponse, error)
	CreateFn                          func(ctx context.Context, in *insysproto.PortingDataCreateRequest, opts []grpc.CallOption) (*insysproto.PortingDataCreateResponse, error)
	PortInCreateBySalesforceBatchIDFn func(ctx context.Context, in *insysproto.PortInCreateBySalesforceBatchIDRequest, opts []grpc.CallOption) (*insysproto.PortInCreateBySalesforceBatchIDResponse, error)
	PortInCreateFn                    func(ctx context.Context, in *insysproto.PortInCreateRequest, opts []grpc.CallOption) (*insysproto.PortInCreateResponse, error)
	PortInStatusBySalesforceBatchIDFn func(ctx context.Context, in *insysproto.PortInStatusBySalesforceBatchIDRequest, opts []grpc.CallOption) (*insysproto.PortInStatusBySalesforceBatchIDResponse, error)
	PortInStatusFn                    func(ctx context.Context, in *insysproto.PortInStatusRequest, opts []grpc.CallOption) (*insysproto.PortInStatusResponse, error)
	UpdateFn                          func(ctx context.Context, in *insysproto.PortingDataUpdateRequest, opts []grpc.CallOption) (*insysproto.PortingDataUpdateResponse, error)
}

func (pdc *PortingDataClient) ByLocationID(ctx context.Context, in *insysproto.PortingDataByLocationIDRequest, opts ...grpc.CallOption) (*insysproto.PortingDataByLocationIDResponse, error) {
	return pdc.ByLocationIDFn(ctx, in, opts)
}

func (pdc *PortingDataClient) Create(ctx context.Context, in *insysproto.PortingDataCreateRequest, opts ...grpc.CallOption) (*insysproto.PortingDataCreateResponse, error) {
	return pdc.CreateFn(ctx, in, opts)
}

func (pdc *PortingDataClient) PortInCreateBySalesforceBatchID(ctx context.Context, in *insysproto.PortInCreateBySalesforceBatchIDRequest, opts ...grpc.CallOption) (*insysproto.PortInCreateBySalesforceBatchIDResponse, error) {
	return pdc.PortInCreateBySalesforceBatchIDFn(ctx, in, opts)
}

func (pdc *PortingDataClient) PortInCreate(ctx context.Context, in *insysproto.PortInCreateRequest, opts ...grpc.CallOption) (*insysproto.PortInCreateResponse, error) {
	return pdc.PortInCreateFn(ctx, in, opts)
}

func (pdc *PortingDataClient) PortInStatusBySalesforceBatchID(ctx context.Context, in *insysproto.PortInStatusBySalesforceBatchIDRequest, opts ...grpc.CallOption) (*insysproto.PortInStatusBySalesforceBatchIDResponse, error) {
	return pdc.PortInStatusBySalesforceBatchIDFn(ctx, in, opts)
}

func (pdc *PortingDataClient) PortInStatus(ctx context.Context, in *insysproto.PortInStatusRequest, opts ...grpc.CallOption) (*insysproto.PortInStatusResponse, error) {
	return pdc.PortInStatusFn(ctx, in, opts)
}

func (pdc *PortingDataClient) Update(ctx context.Context, in *insysproto.PortingDataUpdateRequest, opts ...grpc.CallOption) (*insysproto.PortingDataUpdateResponse, error) {
	return pdc.UpdateFn(ctx, in, opts)
}
