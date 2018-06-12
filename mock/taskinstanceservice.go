package mock

import (
	"context"

	app "weavelab.xyz/insys-onboarding"

	"weavelab.xyz/protorepo/dist/go/enums/insysenums"
	"weavelab.xyz/wlib/uuid"
)

type TaskInstanceService struct {
	ByLocationIDFn    func(ctx context.Context, locationID uuid.UUID) ([]app.TaskInstance, error)
	CreateFromTasksFn func(ctx context.Context, locationID uuid.UUID) ([]app.TaskInstance, error)
	UpdateFn          func(ctx context.Context, id uuid.UUID, status insysenums.OnboardingTaskStatus, statusUpdatedBy string) (*app.TaskInstance, error)
}

func (s *TaskInstanceService) ByLocationID(ctx context.Context, locationID uuid.UUID) ([]app.TaskInstance, error) {
	return s.ByLocationIDFn(ctx, locationID)
}

func (s *TaskInstanceService) CreateFromTasks(ctx context.Context, locationID uuid.UUID) ([]app.TaskInstance, error) {
	return s.CreateFromTasksFn(ctx, locationID)
}

func (s *TaskInstanceService) Update(ctx context.Context, id uuid.UUID, status insysenums.OnboardingTaskStatus, statusUpdatedBy string) (*app.TaskInstance, error) {
	return s.UpdateFn(ctx, id, status, statusUpdatedBy)
}
