package grpc

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"weavelab.xyz/insys-onboarding-service/internal/app"
	"weavelab.xyz/insys-onboarding-service/internal/mock"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/insysproto"
	"weavelab.xyz/monorail/shared/wlib/uuid"
)

func TestHandoffSnapshotServer_CreateOrUpdate(t *testing.T) {

	userID := uuid.NewV4()
	sentAt := time.Now()
	onboardersLocationID := uuid.NewV4()

	type fields struct {
		handOffSnapshotService app.HandoffSnapshotService
	}
	type args struct {
		ctx context.Context
		req *insysproto.HandoffSnapshotCreateOrUpdateRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *insysproto.HandoffSnapshotCreateOrUpdateResponse
		wantErr bool
	}{
		{
			name: "creates snapshot successfully",
			fields: fields{
				handOffSnapshotService: &mock.HandoffSnapshotService{
					CreateOrUpdateFn: func(ctx context.Context, snapshot app.HandoffSnapshot) (app.HandoffSnapshot, error) {
						snapshot.CreatedAt = time.Now()
						snapshot.UpdatedAt = time.Now()
						snapshot.ID = uuid.NewV4()
						return snapshot, nil
					},
				},
			},
			args: args{
				ctx: context.Background(),
				req: &insysproto.HandoffSnapshotCreateOrUpdateRequest{
					HandoffSnapshot: &insysproto.HandoffSnapshotRecord{
						CsatRecipientUserId:  userID.String(),
						CsatSentAt:           sentAt.Format(time.RFC3339),
						OnboardersLocationId: onboardersLocationID.String(),
					},
				},
			},
			want: &insysproto.HandoffSnapshotCreateOrUpdateResponse{
				HandoffSnapshot: &insysproto.HandoffSnapshotRecord{
					OnboardersLocationId: onboardersLocationID.String(),
					CsatRecipientUserId:  userID.String(),
					CsatSentAt:           sentAt.Format(time.RFC3339),
				},
			},
			wantErr: false,
		},
	}

	// custom functions to ignore fields in cmp.Equal comparison
	opts := []cmp.Option{
		cmpopts.IgnoreFields(insysproto.HandoffSnapshotRecord{}, "Id", "CreatedAt", "UpdatedAt", "CsatSentAt"),
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &HandoffSnapshotServer{
				handOffSnapshotService: tt.fields.handOffSnapshotService,
			}
			got, err := s.CreateOrUpdate(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("HandoffSnapshotServer.CreateOrUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("HandoffSnapshotServer.CreateOrUpdate() = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}
