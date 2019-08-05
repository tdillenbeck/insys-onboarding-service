package grpc

import (
	"context"
	"reflect"
	"testing"

	"weavelab.xyz/insys-onboarding-service/internal/app"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/insysproto"
)

func TestOnboarderServer_ReadByUserID(t *testing.T) {
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
			name:    "successfully read by user id",
			wantErr: false,
		},
		{
			name:    "invalid user id in request",
			wantErr: true,
		},
		{
			name:    "no user found",
			wantErr: true,
		},
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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OnboarderServer.ReadByUserID() = %v, want %v", got, tt.want)
			}
		})
	}
}
