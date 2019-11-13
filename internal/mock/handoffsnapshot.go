package mock

import (
	"context"

	"weavelab.xyz/insys-onboarding-service/internal/app"
)

type HandOffSnapshotService struct {
	CreateOrUpdateFn func(ctx context.Context, snapshot app.HandOffSnapshot) (app.HandOffSnapshot, error)
}

func (h *HandOffSnapshotService) CreateOrUpdate(ctx context.Context, snapshot app.HandOffSnapshot) (app.HandOffSnapshot, error) {
	return h.CreateOrUpdateFn(ctx, snapshot)
}
