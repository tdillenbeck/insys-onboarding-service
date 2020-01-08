package app

import (
	"context"

	"weavelab.xyz/monorail/shared/go-utilities/null"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/enums/insysenums"
	"weavelab.xyz/monorail/shared/wlib/uuid"
)

// CategoryService defines the actions for the database related to Categories
type CategoryService interface {
	ByID(ctx context.Context, id uuid.UUID) (*Category, error)
}

type ChiliPiperScheduleEventService interface {
	ByLocationID(ctx context.Context, locationID uuid.UUID) ([]ChiliPiperScheduleEvent, error)
	Cancel(ctx context.Context, eventID string) (*ChiliPiperScheduleEvent, error)
	Create(ctx context.Context, scheduleEvent *ChiliPiperScheduleEvent) (*ChiliPiperScheduleEvent, error)
	Update(ctx context.Context, eventID, assigneeID string, startAt, endAt null.Time) (*ChiliPiperScheduleEvent, error)
}

type OnboarderService interface {
	CreateOrUpdate(ctx context.Context, onboarder *Onboarder) (*Onboarder, error)
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context) ([]Onboarder, error)
	ReadBySalesforceUserID(ctx context.Context, salesforceUserID string) (*Onboarder, error)
	ReadByUserID(ctx context.Context, userID uuid.UUID) (*Onboarder, error)
}

type OnboardersLocationService interface {
	CreateOrUpdate(ctx context.Context, onboardersLocation *OnboardersLocation) (*OnboardersLocation, error)
	ReadByLocationID(ctx context.Context, locationID uuid.UUID) (*OnboardersLocation, error)
	RecordFirstLogin(ctx context.Context, locationID uuid.UUID) error
}

// TaskInstanceService defines the actions for the database related to TaskInstances
type TaskInstanceService interface {
	ByLocationID(ctx context.Context, locationID uuid.UUID) ([]TaskInstance, error)
	CreateFromTasks(ctx context.Context, locationID uuid.UUID) ([]TaskInstance, error)
	SyncTaskInstanceLinksFromOnboarderLinks(ctx context.Context, locationID uuid.UUID) error
	Update(ctx context.Context, id uuid.UUID, status insysenums.OnboardingTaskStatus, statusUpdatedBy string) (*TaskInstance, error)
	UpdateExplanation(ctx context.Context, id uuid.UUID, explanation string) (*TaskInstance, error)
}

type HandoffSnapshotService interface {
	CreateOrUpdate(ctx context.Context, snapshot HandoffSnapshot) (HandoffSnapshot, error)
	ReadByOnboardersLocationID(ctx context.Context, onboardersLocationId uuid.UUID) (HandoffSnapshot, error)
	SubmitCSAT(ctx context.Context, onboardersLocationId uuid.UUID, csatRecipientUserEmail string) (HandoffSnapshot, error)
	SubmitHandoff(ctx context.Context, handoffSnapshot uuid.UUID) (HandoffSnapshot, error)
}
