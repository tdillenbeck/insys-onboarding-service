package mock

import (
	"context"

	app "weavelab.xyz/insys-onboarding"

	"weavelab.xyz/wlib/uuid"
)

type CategoryService struct {
	ByIDFn func(ctx context.Context, id uuid.UUID) (*app.Category, error)
}

func (s *CategoryService) ByID(ctx context.Context, id uuid.UUID) (*app.Category, error) {
	return s.ByIDFn(ctx, id)
}
