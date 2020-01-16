package grpc

import (
"weavelab.xyz/monorail/shared/wlib/uuid"
"weavelab.xyz/monorail/shared/dist/go/messages/insysproto"
"weavelab.xyz/monorail/shared/dist/go/services/insys"
)

func (s *ChiliPiperScheduleEventServer) ReadRescheduleTracking(req *onboardingproto.RescheduleTrackingRequest) *onboardingproto.RescheduleTrackingResponse {

	locationID, err := uuid.Parse(req.LocationID)

}
