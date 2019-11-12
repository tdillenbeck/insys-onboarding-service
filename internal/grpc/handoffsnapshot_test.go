package grpc

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"weavelab.xyz/insys-onboarding-service/internal/app"
	"weavelab.xyz/monorail/shared/go-utilities/null"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/insysproto"
	"weavelab.xyz/monorail/shared/wlib/uuid"
)

func Test_convertProtoToHandOffSnapshot(t *testing.T) {
	id := uuid.NewV4()
	userID := uuid.NewV4()
	onboardersLocationID := uuid.NewV4()
	sentAt := time.Now()
	createdAt := time.Now()
	updatedAt := time.Now()

	type args struct {
		proto *insysproto.HandOffSnapshotCreateOrUpdateRequest
	}
	tests := []struct {
		name    string
		args    args
		want    app.HandOffSnapshot
		wantErr bool
	}{
		{
			name: "ensure conversion works",
			args: args{
				proto: &insysproto.HandOffSnapshotCreateOrUpdateRequest{
					HandoffSnapshot: &insysproto.HandOffSnapshotRecord{
						Id:                   id.String(),
						OnboardersLocationId: onboardersLocationID.String(),
						CsatRecipientUserId:  userID.String(),
						CsatSentAt:           sentAt.UTC().Format(time.RFC3339),
						CreatedAt:            createdAt.UTC().Format(time.RFC3339),
						UpdatedAt:            updatedAt.UTC().Format(time.RFC3339),
					},
				},
			},
			want: app.HandOffSnapshot{
				ID:                   id,
				OnboardersLocationID: onboardersLocationID,
				CustomerSatisfactionSurveyRecipientUserID: null.NewUUIDUUID(userID),
				CustomerSatisfactionSurveySentAt:          null.NewTime(sentAt.UTC()),
				CreatedAt:                                 createdAt,
				UpdatedAt:                                 updatedAt,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := convertProtoToHandOffSnapshot(tt.args.proto)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertProtoToHandOffSnapshot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("convertProtoToHandOffSnapshot() = %v", cmp.Diff(got, tt.want))
			}
		})
	}
}
