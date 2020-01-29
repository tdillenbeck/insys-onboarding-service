package mock

import (
	"context"

	"google.golang.org/grpc"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/insysproto"
)

type ProvisioningService struct {
	CreateOrUpdatePreProvisionFn     func(ctx context.Context, req *insysproto.CreateOrUpdatePreProvisionRequest, opts []grpc.CallOption) (*insysproto.CreateOrUpdatePreProvisionResponse, error)
	InitialProvisionFn               func(ctx context.Context, req *insysproto.InitialProvisionRequest, opts []grpc.CallOption) (*insysproto.InitialProvisionResponse, error)
	PortDayProvisionFn               func(ctx context.Context, req *insysproto.PortDayProvisionRequest, opts []grpc.CallOption) (*insysproto.PortDayProvisionResponse, error)
	PreProvisionByOpportunityIDFn    func(ctx context.Context, req *insysproto.PreProvisionByOpportunityIDRequest, opts []grpc.CallOption) (*insysproto.PreProvisionByOpportunityIDResponse, error)
	PreProvisionsByLocationIDFn      func(ctx context.Context, req *insysproto.PreProvisionsByLocationIDRequest, opts []grpc.CallOption) (*insysproto.PreProvisionsByLocationIDResponse, error)
	ProvisionUserFn                  func(ctx context.Context, req *insysproto.ProvisionUserRequest, opts []grpc.CallOption) (*insysproto.ProvisionUserResponse, error)
	PreProvisionsByParentAccountIDFn func(ctx context.Context, req *insysproto.PreProvisionsByParentAccountIDRequest, opts []grpc.CallOption) (*insysproto.PreProvisionsByParentAccountIDResponse, error)
}

func (s *ProvisioningService) CreateOrUpdatePreProvision(ctx context.Context, req *insysproto.CreateOrUpdatePreProvisionRequest, opts ...grpc.CallOption) (*insysproto.CreateOrUpdatePreProvisionResponse, error) {
	return s.CreateOrUpdatePreProvisionFn(ctx, req, opts)
}

func (s *ProvisioningService) InitialProvision(ctx context.Context, req *insysproto.InitialProvisionRequest, opts ...grpc.CallOption) (*insysproto.InitialProvisionResponse, error) {
	return s.InitialProvisionFn(ctx, req, opts)
}

func (s *ProvisioningService) PortDayProvision(ctx context.Context, req *insysproto.PortDayProvisionRequest, opts ...grpc.CallOption) (*insysproto.PortDayProvisionResponse, error) {
	return s.PortDayProvisionFn(ctx, req, opts)
}

func (s *ProvisioningService) PreProvisionsByLocationID(ctx context.Context, req *insysproto.PreProvisionsByLocationIDRequest, opts ...grpc.CallOption) (*insysproto.PreProvisionsByLocationIDResponse, error) {
	return s.PreProvisionsByLocationIDFn(ctx, req, opts)
}

func (s *ProvisioningService) PreProvisionByOpportunityID(ctx context.Context, req *insysproto.PreProvisionByOpportunityIDRequest, opts ...grpc.CallOption) (*insysproto.PreProvisionByOpportunityIDResponse, error) {
	return s.PreProvisionByOpportunityIDFn(ctx, req, opts)
}

func (s *ProvisioningService) ProvisionUser(ctx context.Context, req *insysproto.ProvisionUserRequest, opts ...grpc.CallOption) (*insysproto.ProvisionUserResponse, error) {
	return s.ProvisionUserFn(ctx, req, opts)
}

func (s *ProvisioningService) PreProvisionsByParentAccountID(ctx context.Context, req *insysproto.PreProvisionsByParentAccountIDRequest, opts ...grpc.CallOption) (*insysproto.PreProvisionsByParentAccountIDResponse, error) {
	return s.PreProvisionsByParentAccountIDFn(ctx, req, opts)
}
