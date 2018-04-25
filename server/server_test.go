package server

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes"
	"weavelab.xyz/protorepo/dist/go/messages/insys/onboardingproto"
	"weavelab.xyz/protorepo/dist/go/messages/sharedproto"
	"weavelab.xyz/wlib/uuid"
)

func TestOnboardingService_Categories(t *testing.T) {
	id, _ := uuid.Parse("6ba7b8149dad11d180b400c04fd430c8")
	expectedID := sharedproto.UUIDToProto(id)

	createdAt, _ := ptypes.TimestampProto(time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC))
	updatedAt, _ := ptypes.TimestampProto(time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC))

	type args struct {
		ctx context.Context
		req *onboardingproto.CategoriesRequest
	}
	tests := []struct {
		name    string
		s       *OnboardingService
		args    args
		want    *onboardingproto.CategoriesResponse
		wantErr bool
	}{
		{
			// Add test cases.
			"hello world",
			&OnboardingService{},
			args{
				ctx: context.Background(),
				req: &onboardingproto.CategoriesRequest{},
			},
			&onboardingproto.CategoriesResponse{
				Categories: []*onboardingproto.Category{
					&onboardingproto.Category{
						ID:           expectedID,
						DisplayText:  "testing categories",
						DisplayOrder: int32(1),
						CreatedAt:    createdAt,
						UpdatedAt:    updatedAt,
					},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OnboardingService{}
			got, err := s.Categories(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("OnboardingService.Categories() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OnboardingService.Categories() = %v, want %v", got, tt.want)
			}
		})
	}
}
