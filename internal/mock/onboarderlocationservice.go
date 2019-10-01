package mock

import (
	"context"

	"weavelab.xyz/insys-onboarding-service/internal/app"
	"weavelab.xyz/monorail/shared/wlib/uuid"
)

type OnboarderLocationService struct {
	CreateOrUpdateFn   func(ctx context.Context, onboardersLocation *app.OnboardersLocation) (*app.OnboardersLocation, error)
	ReadByLocationIDFn func(ctx context.Context, locationID uuid.UUID) (*app.OnboardersLocation, error)
	RecordFirstLoginFn func(ctx context.Context, locationID uuid.UUID) error
}

func (o *OnboarderLocationService) CreateOrUpdate(ctx context.Context, onboardersLocation *app.OnboardersLocation) (*app.OnboardersLocation, error) {
	return o.CreateOrUpdateFn(ctx, onboardersLocation)
}

func (o *OnboarderLocationService) ReadByLocationID(ctx context.Context, locationID uuid.UUID) (*app.OnboardersLocation, error) {
	return o.ReadByLocationIDFn(ctx, locationID)
}

func (o *OnboarderLocationService) RecordFirstLogin(ctx context.Context, locationID uuid.UUID) error {
	return o.RecordFirstLoginFn(ctx, locationID)
}
