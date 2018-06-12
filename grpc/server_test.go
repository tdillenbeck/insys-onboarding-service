package grpc

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes"
	"weavelab.xyz/go-utilities/null"
	app "weavelab.xyz/insys-onboarding"
	"weavelab.xyz/insys-onboarding/mock"
	"weavelab.xyz/protorepo/dist/go/messages/insys/onboardingproto"
	"weavelab.xyz/protorepo/dist/go/messages/sharedproto"
	"weavelab.xyz/wlib/uuid"
)

func TestOnboardingServer_Category(t *testing.T) {
	testUUID := uuid.NewV4()
	createdAt := time.Now()
	updatedAt := time.Now()
	displayText := "Test Display Text"
	displayOrder := 0

	protoUUID := sharedproto.UUIDToProto(testUUID)
	expectedCreatedAt, _ := ptypes.TimestampProto(createdAt)
	expectedUpdatedAt, _ := ptypes.TimestampProto(updatedAt)

	// setup the mock category service
	var cs mock.CategoryService
	cs.ByIDFn = func(ctx context.Context, id uuid.UUID) (*app.Category, error) {
		return &app.Category{
			ID:           id,
			DisplayText:  displayText,
			DisplayOrder: displayOrder,
			CreatedAt:    createdAt,
			UpdatedAt:    updatedAt,
		}, nil
	}

	type fields struct {
		categoryService     app.CategoryService
		taskInstanceService app.TaskInstanceService
	}
	type args struct {
		ctx context.Context
		req *onboardingproto.CategoryRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *onboardingproto.CategoryResponse
		wantErr bool
	}{
		{
			"retrieve category",
			fields{
				categoryService:     &cs,
				taskInstanceService: nil,
			},
			args{
				ctx: context.Background(),
				req: &onboardingproto.CategoryRequest{ID: protoUUID},
			},
			&onboardingproto.CategoryResponse{
				Category: &onboardingproto.Category{
					ID:           protoUUID,
					DisplayText:  "Test Display Text",
					DisplayOrder: int32(0),
					CreatedAt:    expectedCreatedAt,
					UpdatedAt:    expectedUpdatedAt,
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OnboardingServer{
				categoryService:     tt.fields.categoryService,
				taskInstanceService: tt.fields.taskInstanceService,
			}
			got, err := s.Category(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("OnboardingServer.Category() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OnboardingServer.Category() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOnboardingServer_CreateTaskInstancesFromTasks(t *testing.T) {
	testUUID := uuid.NewV4()
	protoUUID := sharedproto.UUIDToProto(testUUID)
	now := time.Now()

	expectedCreatedAt, _ := ptypes.TimestampProto(now)
	expectedUpdatedAt, _ := ptypes.TimestampProto(now)
	expectedCompletedAt, _ := ptypes.TimestampProto(now)
	expectedVerifiedAt, _ := ptypes.TimestampProto(null.Time{Time: time.Time{}, Valid: false}.Time)
	expectedStatusUpdatedAt, _ := ptypes.TimestampProto(now)

	taskInstance := app.TaskInstance{
		ID:         testUUID,
		LocationID: testUUID,
		CategoryID: testUUID,
		TaskID:     testUUID,

		ButtonContent:     null.NewString("Click here"),
		ButtonExternalURL: null.NewString("www.google.com"),
		CompletedAt:       null.NewTime(now),
		CompletedBy:       null.NewString("Donald Duck"),
		Content:           "Testing content",
		DisplayOrder:      0,
		Status:            2,
		StatusUpdatedAt:   now,
		StatusUpdatedBy:   "Donald Duck",
		Title:             "Test Title",
		VerifiedAt:        null.Time{Time: time.Time{}, Valid: false},
		VerifiedBy:        null.String{Str: "", Valid: false},

		CreatedAt: now,
		UpdatedAt: now,
	}

	var tis mock.TaskInstanceService
	tis.CreateFromTasksFn = func(ctx context.Context, locationID uuid.UUID) ([]app.TaskInstance, error) {
		return []app.TaskInstance{taskInstance}, nil
	}

	type fields struct {
		categoryService     app.CategoryService
		taskInstanceService app.TaskInstanceService
	}
	type args struct {
		ctx context.Context
		req *onboardingproto.CreateTaskInstancesFromTasksRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *onboardingproto.TaskInstancesResponse
		wantErr bool
	}{
		{
			"copy tasks into task instances",
			fields{
				categoryService:     nil,
				taskInstanceService: &tis,
			},
			args{
				ctx: context.Background(),
				req: &onboardingproto.CreateTaskInstancesFromTasksRequest{LocationID: protoUUID},
			},
			&onboardingproto.TaskInstancesResponse{
				TaskInstances: []*onboardingproto.TaskInstance{
					{
						ID:         protoUUID,
						LocationID: protoUUID,
						CategoryID: protoUUID,
						TaskID:     protoUUID,

						ButtonContent:     "Click here",
						ButtonExternalURL: "www.google.com",
						CompletedAt:       expectedCompletedAt,
						CompletedBy:       "Donald Duck",
						Content:           "Testing content",
						DisplayOrder:      int32(0),
						Status:            2,
						StatusUpdatedAt:   expectedStatusUpdatedAt,
						StatusUpdatedBy:   "Donald Duck",
						Title:             "Test Title",
						VerifiedAt:        expectedVerifiedAt,
						VerifiedBy:        "",

						CreatedAt: expectedCreatedAt,
						UpdatedAt: expectedUpdatedAt,
					},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OnboardingServer{
				categoryService:     tt.fields.categoryService,
				taskInstanceService: tt.fields.taskInstanceService,
			}
			got, err := s.CreateTaskInstancesFromTasks(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("OnboardingServer.CreateTaskInstancesFromTasks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OnboardingServer.CreateTaskInstancesFromTasks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOnboardingServer_TaskInstances(t *testing.T) {
	testUUID := uuid.NewV4()
	protoUUID := sharedproto.UUIDToProto(testUUID)
	now := time.Now()

	expectedCreatedAt, _ := ptypes.TimestampProto(now)
	expectedUpdatedAt, _ := ptypes.TimestampProto(now)
	expectedCompletedAt, _ := ptypes.TimestampProto(now)
	expectedVerifiedAt, _ := ptypes.TimestampProto(null.Time{Time: time.Time{}, Valid: false}.Time)
	expectedStatusUpdatedAt, _ := ptypes.TimestampProto(now)

	taskInstance := app.TaskInstance{
		ID:         testUUID,
		LocationID: testUUID,
		CategoryID: testUUID,
		TaskID:     testUUID,

		ButtonContent:     null.NewString("Click here"),
		ButtonExternalURL: null.NewString("www.google.com"),
		CompletedAt:       null.NewTime(now),
		CompletedBy:       null.NewString("Donald Duck"),
		Content:           "Testing content",
		DisplayOrder:      0,
		Status:            2,
		StatusUpdatedAt:   now,
		StatusUpdatedBy:   "Donald Duck",
		Title:             "Test Title",
		VerifiedAt:        null.Time{Time: time.Time{}, Valid: false},
		VerifiedBy:        null.String{Str: "", Valid: false},

		CreatedAt: now,
		UpdatedAt: now,
	}

	var tis mock.TaskInstanceService
	tis.ByLocationIDFn = func(ctx context.Context, locationID uuid.UUID) ([]app.TaskInstance, error) {
		return []app.TaskInstance{taskInstance}, nil
	}

	type fields struct {
		categoryService     app.CategoryService
		taskInstanceService app.TaskInstanceService
	}
	type args struct {
		ctx context.Context
		req *onboardingproto.TaskInstancesRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *onboardingproto.TaskInstancesResponse
		wantErr bool
	}{
		{
			"look up task instances by location id",
			fields{
				categoryService:     nil,
				taskInstanceService: &tis,
			},
			args{
				ctx: context.Background(),
				req: &onboardingproto.TaskInstancesRequest{LocationID: protoUUID},
			},
			&onboardingproto.TaskInstancesResponse{
				TaskInstances: []*onboardingproto.TaskInstance{
					{
						ID:         protoUUID,
						LocationID: protoUUID,
						CategoryID: protoUUID,
						TaskID:     protoUUID,

						ButtonContent:     "Click here",
						ButtonExternalURL: "www.google.com",
						CompletedAt:       expectedCompletedAt,
						CompletedBy:       "Donald Duck",
						Content:           "Testing content",
						DisplayOrder:      int32(0),
						Status:            2,
						StatusUpdatedAt:   expectedStatusUpdatedAt,
						StatusUpdatedBy:   "Donald Duck",
						Title:             "Test Title",
						VerifiedAt:        expectedVerifiedAt,
						VerifiedBy:        "",

						CreatedAt: expectedCreatedAt,
						UpdatedAt: expectedUpdatedAt,
					},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OnboardingServer{
				categoryService:     tt.fields.categoryService,
				taskInstanceService: tt.fields.taskInstanceService,
			}
			got, err := s.TaskInstances(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("OnboardingServer.TaskInstances() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OnboardingServer.TaskInstances() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOnboardingServer_UpdateTaskInstance(t *testing.T) {
	t.Skip("Not implemented yet")

	type fields struct {
		categoryService     app.CategoryService
		taskInstanceService app.TaskInstanceService
	}
	type args struct {
		ctx context.Context
		req *onboardingproto.UpdateTaskInstanceRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *onboardingproto.UpdateTaskInstanceResponse
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OnboardingServer{
				categoryService:     tt.fields.categoryService,
				taskInstanceService: tt.fields.taskInstanceService,
			}
			got, err := s.UpdateTaskInstance(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("OnboardingServer.UpdateTaskInstance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OnboardingServer.UpdateTaskInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}
