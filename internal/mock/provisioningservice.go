package mock

import (
	"context"

	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/insysproto"
)

type ProvisioningService struct {
	PreProvisionsByLocationIDFn   func(ctx context.Context, req *insysproto.PreProvisionsByLocationIDRequest) (*insysproto.PreProvisionsByLocationIDResponse, error)
	PreProvisionByOpportunityIDFn func(ctx context.Context, req *insysproto.PreProvisionByOpportunityIDRequest) (*insysproto.PreProvisionByOpportunityIDResponse, error)
	CreateOrUpdatePreProvisionFn  func(ctx context.Context, req *insysproto.CreateOrUpdatePreProvisionRequest) (*insysproto.CreateOrUpdatePreProvisionResponse, error)
	InitialProvisionFn            func(ctx context.Context, req *insysproto.InitialProvisionRequest) (*insysproto.InitialProvisionResponse, error)
	ProvisionUserFn               func(ctx context.Context, req *insysproto.ProvisionUserRequest) (*insysproto.ProvisionUserResponse, error)
}

func (s *ProvisioningService) PreProvisionsByLocationID(ctx context.Context, req *insysproto.PreProvisionsByLocationIDRequest) (*insysproto.PreProvisionsByLocationIDResponse, error) {
	return s.PreProvisionsByLocationIDFn(ctx, req)
}

func (s *ProvisioningService) PreProvisionByOpportunityID(ctx context.Context, req *insysproto.PreProvisionByOpportunityIDRequest) (*insysproto.PreProvisionByOpportunityIDResponse, error) {
	return s.PreProvisionByOpportunityIDFn(ctx, req)
}
func (s *ProvisioningService) CreateOrUpdatePreProvision(ctx context.Context, req *insysproto.CreateOrUpdatePreProvisionRequest) (*insysproto.CreateOrUpdatePreProvisionResponse, error) {
	return s.CreateOrUpdatePreProvision(ctx, req)
}
func (s *ProvisioningService) InitialProvision(ctx context.Context, req *insysproto.InitialProvisionRequest) (*insysproto.InitialProvisionResponse, error) {
	return s.InitialProvision(ctx, req)
}
func (s *ProvisioningService) ProvisionUser(ctx context.Context, req *insysproto.ProvisionUserRequest) (*insysproto.ProvisionUserResponse, error) {
	return s.ProvisionUser(ctx, req)
}
