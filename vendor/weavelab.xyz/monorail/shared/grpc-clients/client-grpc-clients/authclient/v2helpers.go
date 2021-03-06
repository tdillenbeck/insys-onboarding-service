package authclient

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	"weavelab.xyz/monorail/shared/go-utilities/null"
	"weavelab.xyz/monorail/shared/grpc-clients/client-grpc-clients/authclient/authv2proto"
	"weavelab.xyz/monorail/shared/wlib/werror"
	"weavelab.xyz/monorail/shared/wlib/wgrpc/wgrpcproto"
)

func fromGRPCUserProfileV2(r *authv2proto.UserProfileResponse) (UserProfile, error) {

	uid, err := wgrpcproto.UUID(r.UserID)
	if err != nil {
		return UserProfile{}, werror.Wrap(err, "error parsing userID as uuid")
	}

	var verifiedAt time.Time

	if r.MobileNumberVerifiedAt != nil {
		verifiedAt, err = ptypes.Timestamp(r.MobileNumberVerifiedAt)
		if err != nil {
			return UserProfile{}, werror.Wrap(err, "error parsing mobileNumberVerifiedAt")
		}
	}

	u := UserProfile{
		UserID:                 uid,
		Username:               r.Username,
		FirstName:              r.FirstName,
		LastName:               r.LastName,
		Type:                   fromGRPCUserType(r.Type),
		Status:                 r.Status.String(),
		MobileNumber:           r.MobileNumber,
		MobileNumberVerifiedAt: null.NewTimeDefaultAsNull(verifiedAt),
	}

	return u, nil
}

type UserType string

const (
	UserTypeWeave    UserType = "weave"
	UserTypePractice UserType = "practice"
)

func fromGRPCUserType(t authv2proto.UserType) UserType {

	var userType UserType
	switch t {
	case authv2proto.UserType_Weave:
		userType = UserTypeWeave
	case authv2proto.UserType_Practice:
		userType = UserTypePractice
	}

	return userType
}

func toGRPCUserType(s UserType) authv2proto.UserType {

	var userType authv2proto.UserType
	switch s {
	case UserTypeWeave:
		userType = authv2proto.UserType_Weave
	case UserTypePractice:
		userType = authv2proto.UserType_Practice
	}

	return userType
}
