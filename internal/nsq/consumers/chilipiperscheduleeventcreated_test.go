package consumers

import (
	"context"
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/nsqio/go-nsq"
	gnsq "github.com/nsqio/go-nsq"
	"weavelab.xyz/insys-onboarding-service/internal/app"
	"weavelab.xyz/insys-onboarding-service/internal/mock"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/insysproto"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/services/insys"
	"weavelab.xyz/monorail/shared/wlib/uuid"
)

func TestChiliPiperScheduleEventCreatedSubscriber_HandleMessage(t *testing.T) {
	locationID := uuid.NewV4()

	successfulChiliPiperScheduleEventService := &mock.ChiliPiperScheduleEventService{
		CanceledCountByLocationIDAndEventTypeFn: func(ctx context.Context, locationID uuid.UUID, eventType string) (int, error) {
			return 6, nil
		},
	}

	successfulRescheduleTrackingService := &mock.RescheduleTrackingEventService{
		CreateOrUpdateFn: func(ctx context.Context, locationID uuid.UUID, count int, eventType string) (*app.RescheduleTracking, error) {
			return nil, nil
		},
		ReadRescheduleTrackingFn: func(ctx context.Context, in *insysproto.RescheduleTrackingRequest) (*app.RescheduleTracking, error) {
			return nil, nil
		},
	}

	message, _ := proto.Marshal(&insysproto.CreateChiliPiperScheduleEventResponse{
		Event: &insysproto.ChiliPiperScheduleEventRecord{
			LocationId: locationID.String(),
			EventType:  "software_install_call",
		},
	})

	type fields struct {
		chiliPiperScheduleEventService app.ChiliPiperScheduleEventService
		onboarderService               app.OnboarderService
		rescheduleTrackingEventService app.RescheduleTrackingEventService
		featureFlagsClient             FeatureFlagsClient
		onboardersLocationServer       insys.OnboardersLocationServer
		onboardingServer               insys.OnboardingServer
	}
	type args struct {
		ctx context.Context
		m   *gnsq.Message
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "successfully create reschedule event",
			fields: fields{
				chiliPiperScheduleEventService: successfulChiliPiperScheduleEventService,
				rescheduleTrackingEventService: successfulRescheduleTrackingService,
			},
			args: args{
				ctx: context.Background(),
				m: &nsq.Message{
					Body: message,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ChiliPiperScheduleEventCreatedSubscriber{
				chiliPiperScheduleEventService: tt.fields.chiliPiperScheduleEventService,
				onboarderService:               tt.fields.onboarderService,
				rescheduleTrackingEventService: tt.fields.rescheduleTrackingEventService,
				featureFlagsClient:             tt.fields.featureFlagsClient,
				onboardersLocationServer:       tt.fields.onboardersLocationServer,
				onboardingServer:               tt.fields.onboardingServer,
			}
			if err := c.HandleMessage(tt.args.ctx, tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("ChiliPiperScheduleEventCreatedSubscriber.HandleMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
