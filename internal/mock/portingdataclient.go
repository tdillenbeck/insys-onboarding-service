package mock

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/insysproto"
)

type PortingDataClient struct {
	BatchCreateFn                 func(ctx context.Context, in *insysproto.PortingDataBatchCreateRequest, opts []grpc.CallOption) (*insysproto.PortingDataBatchCreateResponse, error)
	ByLocationIDFn                func(ctx context.Context, in *insysproto.PortingDataByLocationIDRequest, opts []grpc.CallOption) (*insysproto.PortingDataByLocationIDResponse, error)
	CreateFn                      func(ctx context.Context, in *insysproto.PortingDataCreateRequest, opts []grpc.CallOption) (*insysproto.PortingDataCreateResponse, error)
	DeleteRequestFn               func(ctx context.Context, in *insysproto.PortingRequestDeleteRequest, opts []grpc.CallOption) (*empty.Empty, error)
	FirstAvailableFOCDateFn       func(ctx context.Context, in *insysproto.FirstAvailableFOCDateRequest, opts []grpc.CallOption) (*insysproto.FirstAvailableFOCDateResponse, error)
	PortInCreateByPortingDataIDFn func(ctx context.Context, in *insysproto.PortInCreateByPortingDataIDRequest, opts []grpc.CallOption) (*insysproto.PortInCreateByPortingDataIDResponse, error)
	PortInCreateFn                func(ctx context.Context, in *insysproto.PortInCreateRequest, opts []grpc.CallOption) (*insysproto.PortInCreateResponse, error)
	PortInStatusByPortingDataIDFn func(ctx context.Context, in *insysproto.PortInStatusByPortingDataIDRequest, opts []grpc.CallOption) (*insysproto.PortInStatusByPortingDataIDResponse, error)
	PortInStatusFn                func(ctx context.Context, in *insysproto.PortInStatusRequest, opts []grpc.CallOption) (*insysproto.PortInStatusResponse, error)
	PortabilityCheckFn            func(ctx context.Context, in *insysproto.PortabilityCheckRequest, opts []grpc.CallOption) (*insysproto.PortabilityCheckResponse, error)
	UpdateFn                      func(ctx context.Context, in *insysproto.PortingDataUpdateRequest, opts []grpc.CallOption) (*insysproto.PortingDataUpdateResponse, error)
}

func (pdc *PortingDataClient) ByLocationID(ctx context.Context, in *insysproto.PortingDataByLocationIDRequest, opts ...grpc.CallOption) (*insysproto.PortingDataByLocationIDResponse, error) {
	return pdc.ByLocationIDFn(ctx, in, opts)
}

func (pdc *PortingDataClient) BatchCreate(ctx context.Context, in *insysproto.PortingDataBatchCreateRequest, opts ...grpc.CallOption) (*insysproto.PortingDataBatchCreateResponse, error) {
	return pdc.BatchCreateFn(ctx, in, opts)
}

func (pdc *PortingDataClient) Create(ctx context.Context, in *insysproto.PortingDataCreateRequest, opts ...grpc.CallOption) (*insysproto.PortingDataCreateResponse, error) {
	return pdc.CreateFn(ctx, in, opts)
}

func (pdc *PortingDataClient) DeleteRequest(ctx context.Context, in *insysproto.PortingRequestDeleteRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	return pdc.DeleteRequestFn(ctx, in, opts)
}

func (pdc *PortingDataClient) FirstAvailableFOCDate(ctx context.Context, in *insysproto.FirstAvailableFOCDateRequest, opts ...grpc.CallOption) (*insysproto.FirstAvailableFOCDateResponse, error) {
	return pdc.FirstAvailableFOCDateFn(ctx, in, opts)
}

func (pdc *PortingDataClient) PortInCreateByPortingDataID(ctx context.Context, in *insysproto.PortInCreateByPortingDataIDRequest, opts ...grpc.CallOption) (*insysproto.PortInCreateByPortingDataIDResponse, error) {
	return pdc.PortInCreateByPortingDataIDFn(ctx, in, opts)
}

func (pdc *PortingDataClient) PortInCreate(ctx context.Context, in *insysproto.PortInCreateRequest, opts ...grpc.CallOption) (*insysproto.PortInCreateResponse, error) {
	return pdc.PortInCreateFn(ctx, in, opts)
}

func (pdc *PortingDataClient) PortInStatusByPortingDataID(ctx context.Context, in *insysproto.PortInStatusByPortingDataIDRequest, opts ...grpc.CallOption) (*insysproto.PortInStatusByPortingDataIDResponse, error) {
	return pdc.PortInStatusByPortingDataIDFn(ctx, in, opts)
}

func (pdc *PortingDataClient) PortInStatus(ctx context.Context, in *insysproto.PortInStatusRequest, opts ...grpc.CallOption) (*insysproto.PortInStatusResponse, error) {
	return pdc.PortInStatusFn(ctx, in, opts)
}

func (pdc *PortingDataClient) PortabilityCheck(ctx context.Context, in *insysproto.PortabilityCheckRequest, opts ...grpc.CallOption) (*insysproto.PortabilityCheckResponse, error) {
	return pdc.PortabilityCheckFn(ctx, in, opts)
}

func (pdc *PortingDataClient) Update(ctx context.Context, in *insysproto.PortingDataUpdateRequest, opts ...grpc.CallOption) (*insysproto.PortingDataUpdateResponse, error) {
	return pdc.UpdateFn(ctx, in, opts)
}
