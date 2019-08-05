package mock

import (
	"context"

	"weavelab.xyz/insys-onboarding-service/internal/app"
	"weavelab.xyz/monorail/shared/wlib/uuid"
)

type OnboarderService struct {
	CreateOrUpdateFn func(ctx context.Context, onboarder *app.Onboarder) (*app.Onboarder, error)
	DeleteFn         func(ctx context.Context, id uuid.UUID) error
	ListFn           func(ctx context.Context) ([]app.Onboarder, error)
	ReadByUserIDFn   func(ctx context.Context, userID uuid.UUID) (*app.Onboarder, error)
}

func (o *OnboarderService) CreateOrUpdate(ctx context.Context, onboarder *app.Onboarder) (*app.Onboarder, error) {
	return o.CreateOrUpdateFn(ctx, onboarder)
}

func (o *OnboarderService) Delete(ctx context.Context, id uuid.UUID) error {
	return o.DeleteFn(ctx, id)
}

func (o *OnboarderService) List(ctx context.Context) ([]app.Onboarder, error) {
	return o.ListFn(ctx)
}

func (o *OnboarderService) ReadByUserID(ctx context.Context, userID uuid.UUID) (*app.Onboarder, error) {
	return o.ReadByUserIDFn(ctx, userID)
}
