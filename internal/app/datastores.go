package app

import (
	"context"

	"weavelab.xyz/protorepo/dist/go/enums/insysenums"
	"weavelab.xyz/wlib/uuid"
)

// CategoryService defines the actions for the database related to Categories
type CategoryService interface {
	ByID(ctx context.Context, id uuid.UUID) (*Category, error)
}

// TaskInstanceService defines the actions for the database related to TaskInstances
type TaskInstanceService interface {
	ByLocationID(ctx context.Context, locationID uuid.UUID) ([]TaskInstance, error)
	CreateFromTasks(ctx context.Context, locationID uuid.UUID) ([]TaskInstance, error)
	Update(ctx context.Context, id uuid.UUID, status insysenums.OnboardingTaskStatus, statusUpdatedBy string) (*TaskInstance, error)
	UpdateExplanation(ctx context.Context, id uuid.UUID, explanation string) (*TaskInstance, error)
}

type OnboarderService interface {
	CreateOrUpdate(ctx context.Context, onboarder *Onboarder) (*Onboarder, error)
	ReadByUserID(ctx context.Context, userID uuid.UUID) (*Onboarder, error)
}

type OnboardersLocationService interface {
	CreateOrUpdate(ctx context.Context, onboardersLocation *OnboardersLocation) (*OnboardersLocation, error)
	ReadByLocationID(ctx context.Context, locationID uuid.UUID) (*OnboardersLocation, error)
}
