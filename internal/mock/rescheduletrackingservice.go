package mock

import (
	"context"

	"weavelab.xyz/insys-onboarding-service/internal/app"
	"weavelab.xyz/monorail/shared/wlib/uuid"
)

type RescheduleTrackingService struct {
	CreateOrUpdateFn func(ctx context.Context, locationID uuid.UUID, count int, eventType string) (*app.RescheduleTracking, error)
}

func (rts *RescheduleTrackingService) CreateOrUpdate(ctx context.Context, locationID uuid.UUID, count int, eventType string) (*app.RescheduleTracking, error) {
	return rts.CreateOrUpdateFn(ctx, locationID, count, eventType)
}
