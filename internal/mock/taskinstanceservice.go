package mock

import (
	"context"

	"weavelab.xyz/insys-onboarding/internal/app"
	"weavelab.xyz/protorepo/dist/go/enums/insysenums"
	"weavelab.xyz/wlib/uuid"
)

type TaskInstanceService struct {
	ByLocationIDFn      func(ctx context.Context, locationID uuid.UUID) ([]app.TaskInstance, error)
	CreateFromTasksFn   func(ctx context.Context, locationID uuid.UUID) ([]app.TaskInstance, error)
	UpdateFn            func(ctx context.Context, id uuid.UUID, status insysenums.OnboardingTaskStatus, statusUpdatedBy string) (*app.TaskInstance, error)
	UpdateExplanationFn func(ctx context.Context, id uuid.UUID, explanation string) (*app.TaskInstance, error)
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

func (s *TaskInstanceService) UpdateExplanation(ctx context.Context, id uuid.UUID, explanation string) (*app.TaskInstance, error) {
	return s.UpdateExplanationFn(ctx, id, explanation)
}
