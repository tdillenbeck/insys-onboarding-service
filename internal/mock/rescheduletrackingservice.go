package mock

import (
	"context"

	"weavelab.xyz/monorail/shared/wlib/uuid"
)

type RescheduleTrackingService struct {
	CreateOrUpdateFn func(ctx context.Context, locationID uuid.UUID, count int, eventType string) error
}

func (rts *RescheduleTrackingService) CreateOrUpdate(ctx context.Context, locationID uuid.UUID, count int, eventType string) error {
	return rts.CreateOrUpdateFn(ctx, locationID, count, eventType)
}
