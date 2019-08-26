package consumers

import (
	"context"
	"strings"

	"github.com/gogo/protobuf/proto"
	nsq "github.com/nsqio/go-nsq"
	"weavelab.xyz/insys-onboarding-service/internal/app"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/insysproto"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/messages/sharedproto"
	"weavelab.xyz/monorail/shared/protorepo/dist/go/services/insys"
	"weavelab.xyz/monorail/shared/wlib/uuid"
	"weavelab.xyz/monorail/shared/wlib/werror"
)

type ChiliPiperScheduleEventCreatedSubscriber struct {
	onboarderService         app.OnboarderService
	onboardersLocationServer insys.OnboardersLocationServer
}

func NewChiliPiperScheduleEventCreatedSubscriber(os app.OnboarderService, ols insys.OnboardersLocationServer) *ChiliPiperScheduleEventCreatedSubscriber {
	return &ChiliPiperScheduleEventCreatedSubscriber{
		onboardersLocationServer: ols,
		onboarderService:         os,
	}
}

func (c ChiliPiperScheduleEventCreatedSubscriber) HandleMessage(ctx context.Context, m *nsq.Message) error {
	var chiliPiperScheduleEventResponse insysproto.CreateChiliPiperScheduleEventResponse

	err := proto.Unmarshal(m.Body, &chiliPiperScheduleEventResponse)
	if err != nil {
		return werror.Wrap(err, "could not unmarshal ChiliPiperScheduleEventCreated message body into proto for insysproto.ChiliPiperScheduleEventResponse struct")
	}

	// Only assign the onboarder on the next steps call event. Assigning the onbaorder will turn on the tracker
	if strings.Contains(chiliPiperScheduleEventResponse.Event.EventType, "next_steps") {
		onboarder, err := c.onboarderService.ReadBySalesforceUserID(ctx, chiliPiperScheduleEventResponse.Event.AssigneeId)
		if err != nil {
			return werror.Wrap(err, "could not read onboarder by salesforce user id").Add("salesforce user id", chiliPiperScheduleEventResponse.Event.AssigneeId)
		}
		locationID, err := uuid.Parse(chiliPiperScheduleEventResponse.Event.LocationId)
		if err != nil {
			return werror.Wrap(err, "could not parse location id from chili piper schedule event create response").Add("LocationId", chiliPiperScheduleEventResponse.Event.LocationId)
		}
		_, err = c.onboardersLocationServer.CreateOrUpdate(ctx, &insysproto.OnboardersLocation{
			OnboarderID: sharedproto.UUIDToProto(onboarder.ID),
			LocationID:  sharedproto.UUIDToProto(locationID),
		})
		if err != nil {
			return werror.Wrap(err, "could not assign onboarder to location").Add("onboarder.ID", onboarder.ID).Add("locationID", locationID)
		}
	}

	return nil
}
