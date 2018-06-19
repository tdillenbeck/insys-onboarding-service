package mock

import (
	"context"

	"weavelab.xyz/insys-onboarding/internal/app"
	"weavelab.xyz/wlib/uuid"
)

type CategoryService struct {
	ByIDFn func(ctx context.Context, id uuid.UUID) (*app.Category, error)
}

func (s *CategoryService) ByID(ctx context.Context, id uuid.UUID) (*app.Category, error) {
	return s.ByIDFn(ctx, id)
}
