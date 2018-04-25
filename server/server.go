package server

import (
	"context"
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes"
	"weavelab.xyz/protorepo/dist/go/messages/sharedproto"

	"weavelab.xyz/protorepo/dist/go/messages/insys/onboardingproto"
	"weavelab.xyz/wlib/uuid"
	"weavelab.xyz/wlib/werror"
	"weavelab.xyz/wlib/wgrpc"
	"weavelab.xyz/wlib/wlog"
)

type OnboardingService struct{}

func (s *OnboardingService) Categories(ctx context.Context, req *onboardingproto.CategoriesRequest) (*onboardingproto.CategoriesResponse, error) {
	wlog.Info("received categories request")

	id, err := uuid.Parse("6ba7b8149dad11d180b400c04fd430c8")
	if err != nil {
		return nil, fmt.Errorf("could not generate uuid")
	}
	protoID := sharedproto.UUIDToProto(id)
	createdAt, err := ptypes.TimestampProto(time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC))
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.Wrap(err, "could not convert created at time from time to timestamp"))
	}
	updatedAt, err := ptypes.TimestampProto(time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC))
	if err != nil {
		return nil, wgrpc.Error(wgrpc.CodeInternal, werror.Wrap(err, "could not convert updated at time from time to timestamp"))
	}

	categories := []*onboardingproto.Category{
		&onboardingproto.Category{
			ID:           protoID,
			DisplayText:  "testing categories",
			DisplayOrder: int32(1),
			CreatedAt:    createdAt,
			UpdatedAt:    updatedAt,
		},
	}
	return &onboardingproto.CategoriesResponse{
		Categories: categories,
	}, nil
}

func (s *OnboardingService) Tasks(ctx context.Context, req *onboardingproto.TasksRequest) (*onboardingproto.TasksResponse, error) {
	return nil, wgrpc.Error(wgrpc.CodeUnimplemented, werror.New("not implemented"))
}

func (s *OnboardingService) UpdateTask(ctx context.Context, req *onboardingproto.UpdateTaskRequest) (*onboardingproto.UpdateTaskResponse, error) {
	return nil, wgrpc.Error(wgrpc.CodeUnimplemented, werror.New("not implemented"))
}
