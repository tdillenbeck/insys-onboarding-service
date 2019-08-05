package grpc

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"weavelab.xyz/insys-onboarding-service/internal/app"
	"weavelab.xyz/insys-onboarding-service/internal/mock"
	"weavelab.xyz/monorail/shared/go-utilities/null"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/insysproto"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/sharedproto"
	"weavelab.xyz/monorail/shared/wlib/uuid"
)

func TestOnboarderServer_ReadByUserID(t *testing.T) {
	existingOnboarderID := uuid.NewV4()

	successfulOnboarderService := &mock.OnboarderService{
		ReadByUserIDFn: func(ctx context.Context, userID uuid.UUID) (*app.Onboarder, error) {
			return &app.Onboarder{
				ID:               userID,
				UserID:           userID,
				SalesforceUserID: null.NewString("testing salesforce user id"),
			}, nil
		},
	}

	type fields struct {
		onboarderService app.OnboarderService
	}
	type args struct {
		ctx context.Context
		req *insysproto.Onboarder
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *insysproto.Onboarder
		wantErr bool
	}{
		{
			name:   "successfully read by user id",
			fields: fields{onboarderService: successfulOnboarderService},
			args:   args{context.Background(), &insysproto.Onboarder{UserID: sharedproto.UUIDToProto(existingOnboarderID)}},
			want: &insysproto.Onboarder{
				ID:               sharedproto.UUIDToProto(existingOnboarderID),
				UserID:           sharedproto.UUIDToProto(existingOnboarderID),
				SalesforceUserID: "testing salesforce user id",
				CreatedAt:        nil,
				UpdatedAt:        nil,
			},
			wantErr: false,
		},
		//		{
		//			name:    "invalid user id in request",
		//			wantErr: true,
		//		},
		//		{
		//			name:    "no user found",
		//			wantErr: true,
		//		},
	}

	opts := []cmp.Option{
		cmpopts.IgnoreFields(insysproto.Onboarder{}, "CreatedAt", "UpdatedAt", "DeletedAt"),
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OnboarderServer{
				onboarderService: tt.fields.onboarderService,
			}
			got, err := s.ReadByUserID(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("OnboarderServer.ReadByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("OnboarderServer.ReadByUserID(). Diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}
