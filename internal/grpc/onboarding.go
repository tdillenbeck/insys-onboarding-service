package grpc

import (
	"context"

	"github.com/golang/protobuf/ptypes"

	"weavelab.xyz/insys-onboarding-service/internal/app"

	"weavelab.xyz/monorail/shared/protorepo/dist/go/enums/insysenums"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/insysproto"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/sharedproto"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/services/insys"
	"weavelab.xyz/monorail/shared/wlib/werror"
	"weavelab.xyz/monorail/shared/wlib/wgrpc"
)

// verify that the OnboardingService struct implements all methods required in the proto definition
var _ insys.OnboardingServer = &OnboardingServer{}

type OnboardingServer struct {
	categoryService     app.CategoryService
	taskInstanceService app.TaskInstanceService
}

func NewOnboardingServer(cs app.CategoryService, tis app.TaskInstanceService) *OnboardingServer {
	return &OnboardingServer{
		categoryService:     cs,
		taskInstanceService: tis,
	}
}

// CreateTaskInstancesFromTasks is the grpc method to handle creating task instances from the tasks database table.
func (s *OnboardingServer) CreateTaskInstancesFromTasks(ctx context.Context, req *insysproto.CreateTaskInstancesFromTasksRequest) (*insysproto.TaskInstancesResponse, error) {
	locationUUID, err := req.LocationID.UUID()
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.Wrap(err, "could not parse request location id into a uuid").Add("req.LocationID", req.LocationID))
	}

	onboardingTasks, err := s.taskInstanceService.CreateFromTasks(ctx, locationUUID)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.Wrap(err, "error creating task instances from tasks in the database"))
	}

	taskInstances, err := convertToTaskInstancesProto(onboardingTasks)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.Wrap(err, "could not convert database tasks to protobuf tasks"))
	}

	return &insysproto.TaskInstancesResponse{
		TaskInstances: taskInstances,
	}, nil
}

// Category is the grpc method to retrieve a category from the database
func (s *OnboardingServer) Category(ctx context.Context, req *insysproto.CategoryRequest) (*insysproto.CategoryResponse, error) {
	categoryUUID, err := req.ID.UUID()
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.Wrap(err, "could not parse request category id into a uuid").Add("req.ID", req.ID))
	}

	onboardingCategory, err := s.categoryService.ByID(ctx, categoryUUID)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.Wrap(err, "error retrieving categories from database"))
	}

	category, err := convertToCategoryProto(onboardingCategory)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.Wrap(err, "could not convert database category to protobuf category"))
	}

	return &insysproto.CategoryResponse{
		Category: category,
	}, nil
}

func (s *OnboardingServer) TaskInstances(ctx context.Context, req *insysproto.TaskInstancesRequest) (*insysproto.TaskInstancesResponse, error) {
	locationUUID, err := req.LocationID.UUID()
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.Wrap(err, "could not parse request location id into a uuid").Add("req.LocationID", req.LocationID))
	}

	onboardingTasks, err := s.taskInstanceService.ByLocationID(ctx, locationUUID)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.Wrap(err, "error retrieving tasks from database"))
	}

	taskInstances, err := convertToTaskInstancesProto(onboardingTasks)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.Wrap(err, "could not convert database tasks to protobuf tasks"))
	}

	return &insysproto.TaskInstancesResponse{
		TaskInstances: taskInstances,
	}, nil
}

func (s *OnboardingServer) UpdateTaskInstance(ctx context.Context, req *insysproto.UpdateTaskInstanceRequest) (*insysproto.UpdateTaskInstanceResponse, error) {
	taskInstanceUUID, err := req.ID.UUID()
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.Wrap(err, "could not parse request task instance id into a uuid").Add("req.ID", req.ID))
	}

	onboardingTaskInstance, err := s.taskInstanceService.Update(ctx, taskInstanceUUID, req.Status, req.StatusUpdatedBy)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.Wrap(err, "error updating task database record"))
	}

	task, err := convertToTaskInstanceProto(*onboardingTaskInstance)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.Wrap(err, "could not convert database task record to protobuf task format"))
	}

	return &insysproto.UpdateTaskInstanceResponse{
		TaskInstance: task,
	}, nil
}

func (s *OnboardingServer) UpdateTaskInstanceExplanation(ctx context.Context, req *insysproto.UpdateTaskInstanceExplanationRequest) (*insysproto.UpdateTaskInstanceResponse, error) {
	taskInstanceUUID, err := req.ID.UUID()
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.Wrap(err, "could not parse request task instance id into a uuid").Add("req.ID", req.ID))
	}

	onboardingTaskInstance, err := s.taskInstanceService.UpdateExplanation(ctx, taskInstanceUUID, req.Explanation)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.Wrap(err, "error updating task database record"))
	}

	task, err := convertToTaskInstanceProto(*onboardingTaskInstance)
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.Wrap(err, "could not convert database task record to protobuf task format"))
	}

	return &insysproto.UpdateTaskInstanceResponse{
		TaskInstance: task,
	}, nil
}

func convertToCategoryProto(oc *app.Category) (*insysproto.Category, error) {
	createdAt, err := ptypes.TimestampProto(oc.CreatedAt)
	if err != nil {
		return nil, werror.Wrap(err, "could not convert category created at time")
	}
	updatedAt, err := ptypes.TimestampProto(oc.UpdatedAt)
	if err != nil {
		return nil, werror.Wrap(err, "could not convert category updated at time")
	}
	return &insysproto.Category{
		ID:           sharedproto.UUIDToProto(oc.ID),
		DisplayText:  oc.DisplayText,
		DisplayOrder: int32(oc.DisplayOrder),
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}, nil
}

func convertToTaskInstanceProto(t app.TaskInstance) (*insysproto.TaskInstance, error) {

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

	taskInstance := &insysproto.TaskInstance{
		ID:         sharedproto.UUIDToProto(t.ID),
		LocationID: sharedproto.UUIDToProto(t.LocationID),
		CategoryID: sharedproto.UUIDToProto(t.CategoryID),
		TaskID:     sharedproto.UUIDToProto(t.TaskID),

		CompletedAt:       completedAt,
		CompletedBy:       t.CompletedBy.String(),
		VerifiedAt:        verifiedAt,
		VerifiedBy:        t.VerifiedBy.String(),
		ButtonContent:     t.ButtonContent.String(),
		ButtonExternalURL: t.ButtonExternalURL.String(),
		ButtonInternalURL: t.ButtonInternalURL.String(),
		Content:           t.Content,
		DisplayOrder:      int32(t.DisplayOrder),
		Status:            insysenums.OnboardingTaskStatus(t.Status),
		StatusUpdatedAt:   statusUpdatedAt,
		StatusUpdatedBy:   t.StatusUpdatedBy.String(),
		Title:             t.Title,
		Explanation:       t.Explanation.String(),

		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	return taskInstance, nil
}

func convertToTaskInstancesProto(onboardingTasks []app.TaskInstance) ([]*insysproto.TaskInstance, error) {
	var taskInstances []*insysproto.TaskInstance

	for _, t := range onboardingTasks {
		taskInstance, err := convertToTaskInstanceProto(t)
		if err != nil {
			return nil, err
		}
		taskInstances = append(taskInstances, taskInstance)
	}

	return taskInstances, nil
}
