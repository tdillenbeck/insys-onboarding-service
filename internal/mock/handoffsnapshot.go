package mock

import (
	"context"

	"weavelab.xyz/monorail/shared/wlib/uuid"

	"weavelab.xyz/insys-onboarding-service/internal/app"
)

type HandoffSnapshotService struct {
	CreateOrUpdateFn             func(ctx context.Context, snapshot app.HandoffSnapshot) (app.HandoffSnapshot, error)
	ReadByOnboardersLocationIDFn func(ctx context.Context, onboardersLocationID uuid.UUID) (app.HandoffSnapshot, error)
	SubmitCSATFn                 func(ctx context.Context, onboardersLocationID uuid.UUID, csatRecipientUserEmail string) (app.HandoffSnapshot, error)
	SubmitHandoffFn              func(ctx context.Context, onboardersLocationID uuid.UUID) (app.HandoffSnapshot, error)
}

func (h *HandoffSnapshotService) CreateOrUpdate(ctx context.Context, snapshot app.HandoffSnapshot) (app.HandoffSnapshot, error) {
	return h.CreateOrUpdateFn(ctx, snapshot)
}

func (h *HandoffSnapshotService) ReadByOnboardersLocationID(ctx context.Context, onboardersLocationID uuid.UUID) (app.HandoffSnapshot, error) {
	return h.ReadByOnboardersLocationIDFn(ctx, onboardersLocationID)
}

func (h *HandoffSnapshotService) SubmitCSAT(ctx context.Context, onboardersLocationID uuid.UUID, csatRecipientUserEmail string) (app.HandoffSnapshot, error) {
	return h.SubmitCSATFn(ctx, onboardersLocationID, csatRecipientUserEmail)
}

func (h *HandoffSnapshotService) SubmitHandoff(ctx context.Context, onboardersLocationID uuid.UUID) (app.HandoffSnapshot, error) {
	return h.SubmitHandoffFn(ctx, onboardersLocationID)
}
