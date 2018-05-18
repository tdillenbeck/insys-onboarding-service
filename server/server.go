package server

import (
	"context"

	"github.com/golang/protobuf/ptypes"
	"weavelab.xyz/insys-onboarding/db"
	"weavelab.xyz/protorepo/dist/go/enums/insysenums"
	"weavelab.xyz/protorepo/dist/go/messages/insys/onboardingproto"
	"weavelab.xyz/protorepo/dist/go/messages/sharedproto"
	"weavelab.xyz/wlib/werror"
	"weavelab.xyz/wlib/wgrpc"
)

type OnboardingService struct{}

func (s *OnboardingService) Category(ctx context.Context, req *onboardingproto.CategoryRequest) (*onboardingproto.CategoryResponse, error) {
	categoryUUID, err := req.ID.UUID()
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.Wrap(err, "could not parse request category id into a uuid").Add("req.ID", req.ID))
	}

	onboardingCategory, err := db.Category(ctx, categoryUUID)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.Wrap(err, "error retrieving categories from database"))
	}

	category, err := convertToCategoryProto(onboardingCategory)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.Wrap(err, "could not convert database category to protobuf category"))
	}

	return &onboardingproto.CategoryResponse{
		Category: category,
	}, nil
}

func (s *OnboardingService) TaskInstances(ctx context.Context, req *onboardingproto.TaskInstancesRequest) (*onboardingproto.TaskInstancesResponse, error) {
	locationUUID, err := req.LocationID.UUID()
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.Wrap(err, "could not parse request location id into a uuid").Add("req.LocationID", req.LocationID))
	}

	onboardingTasks, err := db.TaskInstances(ctx, locationUUID)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.Wrap(err, "error retrieving tasks from database"))
	}

	tasks, err := convertToTasksProto(onboardingTasks)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.Wrap(err, "could not convert database tasks to protobuf tasks"))
	}

	return &onboardingproto.TaskInstancesResponse{
		TaskInstances: tasks,
	}, nil
}

func (s *OnboardingService) UpdateTaskInstance(ctx context.Context, req *onboardingproto.UpdateTaskInstanceRequest) (*onboardingproto.UpdateTaskInstanceResponse, error) {
	taskInstanceUUID, err := req.ID.UUID()
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.Wrap(err, "could not parse request task instance id into a uuid").Add("req.ID", req.ID))
	}

	onboardingTaskInstance, err := db.UpdateTaskInstance(ctx, taskInstanceUUID, req.Status, req.StatusUpdatedBy)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.Wrap(err, "error updating task database record"))
	}

	task, err := convertToTaskProto(onboardingTaskInstance)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.Wrap(err, "could not convert database task record to protobuf task format"))
	}

	return &onboardingproto.UpdateTaskInstanceResponse{
		TaskInstance: task,
	}, nil
}

func convertToCategoryProto(oc db.OnboardingCategory) (*onboardingproto.Category, error) {
	createdAt, err := ptypes.TimestampProto(oc.CreatedAt)
	if err != nil {
		return nil, werror.Wrap(err, "could not convert category created at time")
	}
	updatedAt, err := ptypes.TimestampProto(oc.UpdatedAt)
	if err != nil {
		return nil, werror.Wrap(err, "could not convert category updated at time")
	}
	return &onboardingproto.Category{
		ID:           sharedproto.UUIDToProto(oc.ID),
		DisplayText:  oc.DisplayText,
		DisplayOrder: int32(oc.DisplayOrder),
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}, nil
}

func convertToTaskProto(t db.OnboardingTaskInstance) (*onboardingproto.TaskInstance, error) {

	completedAt, err := ptypes.TimestampProto(t.CompletedAt.Time)
	if err != nil {
		return nil, werror.Wrap(err, "could not convert task completed at time")
	}
	verifiedAt, err := ptypes.TimestampProto(t.VerifiedAt.Time)
	if err != nil {
		return nil, werror.Wrap(err, "could not convert task completed at time")
	}
	statusUpdatedAt, err := ptypes.TimestampProto(t.StatusUpdatedAt)
	if err != nil {
		return nil, werror.Wrap(err, "could not convert task status updated at time")
	}
	createdAt, err := ptypes.TimestampProto(t.CreatedAt)
	if err != nil {
		return nil, werror.Wrap(err, "could not convert task created at time")
	}
	updatedAt, err := ptypes.TimestampProto(t.UpdatedAt)
	if err != nil {
		return nil, werror.Wrap(err, "could not convert task updated at time")
	}

	task := &onboardingproto.TaskInstance{
		ID:         sharedproto.UUIDToProto(t.ID),
		LocationID: sharedproto.UUIDToProto(t.LocationID),
		CategoryID: sharedproto.UUIDToProto(t.CategoryID),
		TaskID:     sharedproto.UUIDToProto(t.TaskID),

		CompletedAt:     completedAt,
		CompletedBy:     t.CompletedBy.String(),
		VerifiedAt:      verifiedAt,
		VerifiedBy:      t.VerifiedBy.String(),
		Content:         t.Content,
		DisplayOrder:    int32(t.DisplayOrder),
		Status:          insysenums.OnboardingTaskStatus(t.Status),
		StatusUpdatedAt: statusUpdatedAt,
		Title:           t.Title,

		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	return task, nil
}

func convertToTasksProto(onboardingTasks []db.OnboardingTaskInstance) ([]*onboardingproto.TaskInstance, error) {
	var tasks []*onboardingproto.TaskInstance

	for _, t := range onboardingTasks {
		task, err := convertToTaskProto(t)
		if err != nil {
			return tasks, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}
