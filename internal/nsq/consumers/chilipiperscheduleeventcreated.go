package consumers

import (
	"context"
	"strings"

	"github.com/gogo/protobuf/proto"
	nsq "github.com/nsqio/go-nsq"
	"weavelab.xyz/insys-onboarding-service/internal/app"
	"weavelab.xyz/monorail/shared/wlib/werror"
)

type ChiliPiperScheduleEventCreatedSubscriber struct {
	onboarderLocationService app.OnboardersLocationService
	onboarderService         app.OnboarderService
}

func NewChiliPiperScheduleEventCreatedSubscriber(os app.OnboarderService, ols app.OnboardersLocationService) *ChiliPiperScheduleEventCreatedSubscriber {
	return &ChiliPiperScheduleEventCreatedSubscriber{
		onboarderLocationService: ols,
		onboarderService:         os,
	}
}

func (c ChiliPiperScheduleEventCreatedSubscriber) HandleMessage(ctx context.Context, m *nsq.Message) error {
	var chiliPiperScheduleEventResponse insysproto.ChiliPiperScheduleEventResponse
	err := proto.Unmarshal(m.Body, &chiliPiperScheduleEventResponse)
	if err != nil {
		return werror.Wrap(err, "could not unmarshal ChiliPiperScheduleEventCreated message body into proto for insysproto.ChiliPiperScheduleEventResponse struct")
	}

	// Only assign the onboarder on the next steps call event
	if strings.Contains(chiliPiperScheduleEventResponse.EventType, "next_steps") {

		onboarderID, err := c.onboarderService.ReadBySalesforceUserID(chiliPiperScheduleEventResponse.AssigneeID)
		if err != nil {
		}

		_, err := c.onboarderLocationService.CreateOrUpdate(ctx, &app.OnboardersLocation{
			OnboarderID: onboarderID,
			LocationID:  chiliPiperScheduleEventResponse.LocationID,
		})
	}

}
