package mock

import (
	"context"

	"weavelab.xyz/insys-onboarding-service/internal/app"
)

type HandoffSnapshotService struct {
	CreateOrUpdateFn func(ctx context.Context, snapshot app.HandoffSnapshot) (app.HandoffSnapshot, error)
}

func (h *HandoffSnapshotService) CreateOrUpdate(ctx context.Context, snapshot app.HandoffSnapshot) (app.HandoffSnapshot, error) {
	return h.CreateOrUpdateFn(ctx, snapshot)
}
