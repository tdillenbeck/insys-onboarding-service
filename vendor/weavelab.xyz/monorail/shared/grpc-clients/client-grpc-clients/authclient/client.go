package authclient

import (
	"context"
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"weavelab.xyz/monorail/shared/go-utilities/null"
	"weavelab.xyz/monorail/shared/grpc-clients/client-grpc-clients/authclient/authproto"
	"weavelab.xyz/monorail/shared/grpc-clients/client-grpc-clients/authclient/authv2proto"
	"weavelab.xyz/monorail/shared/wiggum"
	"weavelab.xyz/monorail/shared/wlib/uuid"
	"weavelab.xyz/monorail/shared/wlib/wdns"
	"weavelab.xyz/monorail/shared/wlib/werror"
	"weavelab.xyz/monorail/shared/wlib/wgrpc"
	"weavelab.xyz/monorail/shared/wlib/wgrpc/wgrpcclient"
	"weavelab.xyz/monorail/shared/wlib/wgrpc/wgrpcproto"
	"weavelab.xyz/monorail/shared/wlib/wgrpc/wgrpcproto/wgrpcprotouuid"
)

type Auth struct {
	s  authproto.AuthClient
	v2 authv2proto.AuthV2Client
}

type UserProfile struct {
	UserID                 uuid.UUID `json:"UserID,omitempty"`
	Username               string
	FirstName              string
	LastName               string
	Type                   UserType
	Status                 string
	MobileNumber           string
	MobileNumberVerifiedAt null.Time `json:"MobileNumberVerifiedAt,omitempty"`
}

type CreateUpdateUserProfile struct {
	UserID                 null.UUID `json:"UserID,omitempty"`
	Username               string
	FirstName              string
	LastName               string
	Password               string `json:"Password,omitempty"`
	Type                   UserType
	Status                 string
	MobileNumber           string
	MobileNumberVerifiedAt null.Time `json:"MobileNumberVerifiedAt,omitempty"`
}

type UsersSearch struct {
	LocationIDs []string
	SearchTerms []string
	Limit       int32
	Skip        int32
}
type Role struct {
	ID   int32
	Name string
	Type string
}

type UserAccess struct {
	UserID                 null.UUID
	Username               string
	FirstName              string
	LastName               string
	Type                   UserType
	Locations              []Location
	Status                 string
	MobileNumber           string
	MobileNumberVerifiedAt null.Time `json:"MobileNumberVerifiedAt,omitempty"`
}

type CreateUpdateUserAccess struct {
	UserID                 null.UUID
	Username               string
	FirstName              string
	LastName               string
	Type                   UserType
	Password               string
	Locations              []Location
	MobileNumber           string
	MobileNumberVerifiedAt null.Time `json:"MobileNumberVerifiedAt,omitempty"`
}
type Location struct {
	LocationID uuid.UUID
	Roles      []Role
	Accepted   bool
}

const (
	defaultAuthAddress = "auth-api.auth.svc.cluster.local:grpc"
)

func New(ctx context.Context, addr string) (*Auth, error) {

	if addr == "" {
		var err error
		addr, err = wdns.ResolveAddress(defaultAuthAddress)
		if err != nil {
			return nil, werror.Wrap(err, "unable to use default address")
		}
	}

	g, err := wgrpcclient.NewDefault(ctx, addr)
	if err != nil {
		return nil, werror.Wrap(err, "unable to setup grpc client")
	}

	c := authproto.NewAuthClient(g)
	v2Client := authv2proto.NewAuthV2Client(g)

	a := Auth{
		s:  c,
		v2: v2Client,
	}

	return &a, nil

}

func fromGRPCUser(r *authproto.User) (UserProfile, error) {
	uid, err := uuid.Parse(r.UserID)
	if err != nil {
		return UserProfile{}, werror.Wrap(err, "error parsing userID as uuid")
	}
	var mobileNumberVerfiedAt time.Time
	if r.MobileNumberVerifiedAt != nil {
		mobileNumberVerfiedAt, err = ptypes.Timestamp(r.MobileNumberVerifiedAt)
		if err != nil {
			return UserProfile{}, werror.Wrap(err, "error parsing MobileNumberVerifiedAt")
		}
	}

	u := UserProfile{
		UserID:                 uid,
		Username:               r.Username,
		FirstName:              r.FirstName,
		LastName:               r.LastName,
		Type:                   UserType(r.Type),
		MobileNumber:           r.MobileNumber,
		MobileNumberVerifiedAt: null.NewTimeDefaultAsNull(mobileNumberVerfiedAt),
	}

	return u, nil
}

type User = UserProfile

type LocationUser struct {
	UserProfile
	Roles []Role
}

// Users is deprecated and is used by hydrators to get basic user data
func (a *Auth) Users(ctx context.Context, userIDs []uuid.UUID) ([]User, error) {
	users, err := a.UserProfiles(ctx, userIDs)
	if err != nil {
		return nil, werror.Wrap(err)
	}

	return users, nil
}

// User is deprecated and is used by hydrators to get basic user data
func (a *Auth) User(ctx context.Context, userID uuid.UUID) (User, error) {
	users, err := a.Users(ctx, []uuid.UUID{userID})
	if err != nil {
		return User{}, werror.Wrap(err)
	}

	if len(users) == 0 {
		return User{}, werror.New("user with userID not found").Add("userID", userID.String()).SetCode(werror.CodeNotFound)
	}

	return users[0], nil
}

// UpdateLegacyProfile - uses legacy proto to update user profile
func (a *Auth) UpdateLegacyProfile(ctx context.Context, user UserProfile, password string) error {
	var verifiedAt *timestamp.Timestamp
	var err error

	if user.MobileNumberVerifiedAt.Valid {
		verifiedAt, err = ptypes.TimestampProto(user.MobileNumberVerifiedAt.Time)
		if err != nil {
			return werror.Wrap(err, "error parsing MobileNumberVerifiedAt")
		}
	}

	in := &authproto.Profile{
		Password:               password,
		Username:               user.Username,
		FirstName:              user.FirstName,
		LastName:               user.LastName,
		Type:                   string(user.Type),
		UserID:                 user.UserID.String(),
		Status:                 authproto.StatusType(authproto.StatusType_value[user.Status]),
		MobileNumber:           user.MobileNumber,
		MobileNumberVerifiedAt: verifiedAt,
	}

	_, err = a.s.UpdateProfile(ctx, in)
	if err != nil {
		return werror.Wrap(err)
	}

	return nil
}

// UserProfile uses the version 2 client to get the profile for a userID. It returns false if a user with that ID doesn't exist.
func (a *Auth) UserProfiles(ctx context.Context, userIDs []uuid.UUID) ([]UserProfile, error) {

	ids := make([]*wgrpcprotouuid.UUID, len(userIDs))
	for i, v := range userIDs {
		id := wgrpcproto.UUIDProto(v)

		ids[i] = id
	}

	req := authv2proto.UserProfilesRequest{
		UserIDs: ids,
	}

	results, err := a.v2.UserProfiles(ctx, &req)
	if err != nil {
		return nil, werror.Wrap(err, "unable to find users")
	}

	profiles := make([]User, len(results.Profiles))
	for i, v := range results.Profiles {
		u, err := fromGRPCUserProfileV2(v)
		if err != nil {
			return nil, werror.Wrap(err)
		}

		profiles[i] = u
	}

	return profiles, nil
}

// UserProfile uses the version 2 client to get the profile for a userID. It returns false if a user with that ID doesn't exist.
func (a *Auth) UserProfile(ctx context.Context, userID uuid.UUID) (UserProfile, bool, error) {
	r, err := a.v2.UserProfile(ctx, wgrpcproto.UUIDProto(userID))
	if status, ok := status.FromError(err); ok {
		if status.Code() == codes.Code(wgrpc.CodeNotFound) {
			return UserProfile{}, false, nil
		}
	}
	if err != nil {
		return UserProfile{}, false, werror.Wrap(err, "unable to find user").Add("userID", userID)
	}

	u, err := fromGRPCUserProfileV2(r)
	if err != nil {
		return UserProfile{}, true, werror.Wrap(err)
	}

	return u, true, nil
}

// SearchUserProfiles takes locationID, and search term to search all users across the system
func (a *Auth) SearchUserProfiles(ctx context.Context, locations []string, searchTerm string, limit, skip int) ([]UserProfile, error) {

	searcher := &authproto.UserSearch{
		LocationIDs: locations,
		SearchTerms: strings.Split(searchTerm, " "),
		Limit:       int32(limit),
		Skip:        int32(skip),
	}

	up, err := a.s.FindUsers(ctx, searcher)
	if err != nil {
		return nil, werror.Wrap(err)
	}

	users := make([]UserProfile, 0, len(up.Users))
	for _, v := range up.Users {

		u, err := fromGRPCUser(v)
		if err != nil {
			return nil, werror.Wrap(err, "error parsing user")
		}

		users = append(users, u)
	}
	return users, nil
}

// UserProfileByUsername uses the version 2 client to get a profile for a username. It returns false if a user with that username doesn't exist.
func (a *Auth) UserProfileByUsername(ctx context.Context, username string) (UserProfile, bool, error) {
	req := &authv2proto.FindByUsernameRequest{
		Username: username,
	}

	r, err := a.v2.FindProfileByUsername(ctx, req)
	if status, ok := status.FromError(err); ok {
		if status.Code() == codes.Code(wgrpc.CodeNotFound) {
			return UserProfile{}, false, nil
		}
	}
	if err != nil {
		return UserProfile{}, false, werror.Wrap(err, "unable to get user").Add("username", username)
	}

	u, err := fromGRPCUserProfileV2(r)
	if err != nil {
		return UserProfile{}, false, werror.Wrap(err)
	}

	return u, true, nil
}

type ClientMetadata struct {
	ClientType        string
	RequestingService []string
	RequestHeaders    map[string][]string
}

//Login returns the token as a string -- can pass nil for clientMetadata
func (a *Auth) Login(ctx context.Context, username, password string, expiration, refreshwindow int32, clientMetadata *ClientMetadata) (string, error) {

	in := &authproto.LoginRequest{
		Username:      username,
		Password:      password,
		Expiration:    expiration,
		RefreshWindow: refreshwindow,
	}

	//Add ClientMetadata if its populated
	if clientMetadata != nil {
		//Shove all the request headers into a map
		headers := make(map[string]string, len(clientMetadata.RequestHeaders))
		for k, v := range clientMetadata.RequestHeaders {
			headers[k] = strings.Join(v, "|")
		}
		in.ClientMetadata = &authproto.ClientMetadata{
			ClientType:        clientMetadata.ClientType,
			RequestingService: clientMetadata.RequestingService,
			RequestHeaders:    headers,
		}
	}

	r, err := a.s.Login(ctx, in)
	if err != nil {
		return "", werror.Wrap(err, "unable to login").Add("username", username)
	}

	return r.Token, nil
}

//PublicKey returns the public Key as a byte array
func (a *Auth) PublicKey(ctx context.Context) ([]byte, error) {

	r, err := a.s.PublicKey(ctx, &authproto.EmptyRequest{})
	if err != nil {
		return nil, werror.Wrap(err, "unable to fetch public key")
	}

	return r.PublicKey, nil
}

//LocationUsers returns all users associated with the location and their ACL's for that location
func (a *Auth) LocationUsers(ctx context.Context, locationID uuid.UUID) ([]UserAccess, error) {

	r, err := a.s.LocationUsers(ctx, &authproto.UsersRequest{LocationID: locationID.String()})
	if err != nil {
		return nil, werror.Wrap(err, "unable to fetch location's users")
	}

	possibleRoles, err := a.MapOfRoles(ctx)
	if err != nil {
		return nil, werror.Wrap(err, "error fetching map of possible roles")
	}

	//map users
	users := make([]UserAccess, 0, len(r.Users))

	for _, v := range r.Users {
		uid, err := uuid.Parse(v.UserID)
		if err != nil {
			return nil, werror.Wrap(err, "error parsing userID as uuid")
		}

		locations, err := mapFromProtoLocation(v.Locations, possibleRoles)
		if err != nil {
			return nil, werror.Wrap(err, "error parsing location")
		}

		var verifiedAt time.Time
		if v.MobileNumberVerifiedAt != nil {
			verifiedAt, err = ptypes.Timestamp(v.MobileNumberVerifiedAt)
			if err != nil {
				return nil, werror.Wrap(err, "error parsing MobileNumberVerifiedAt")
			}
		}

		users = append(users, UserAccess{
			FirstName:              v.FirstName,
			LastName:               v.LastName,
			Type:                   UserType(v.Type),
			UserID:                 null.NewUUIDUUID(uid),
			Username:               v.Username,
			Locations:              locations,
			Status:                 v.Status.String(),
			MobileNumber:           v.MobileNumber,
			MobileNumberVerifiedAt: null.NewTimeDefaultAsNull(verifiedAt),
		})
	}

	return users, nil
}

//UserLocations returns all locations associated with that user and their ACL's for each location
func (a *Auth) UserLocations(ctx context.Context, id uuid.UUID) (*UserAccess, error) {

	r, err := a.v2.User(ctx, wgrpcproto.UUIDProto(id))
	if err != nil {
		return nil, werror.Wrap(err, "unable to fetch user locations").Add("uuid", id)
	}

	possibleRoles, err := a.MapOfRoles(ctx)
	if err != nil {
		return nil, werror.Wrap(err, "error fetching map of possible roles")
	}

	var verifiedAt time.Time
	if r.Profile.MobileNumberVerifiedAt != nil {
		verifiedAt, err = ptypes.Timestamp(r.Profile.MobileNumberVerifiedAt)
		if err != nil {
			return nil, werror.Wrap(err, "error parsing MobileNumberVerifiedAt")
		}
	}

	userAccess := &UserAccess{
		UserID:                 null.NewUUIDUUID(id),
		Username:               r.Profile.Username,
		FirstName:              r.Profile.FirstName,
		LastName:               r.Profile.LastName,
		Type:                   fromGRPCUserType(r.Profile.Type),
		MobileNumber:           r.Profile.MobileNumber,
		MobileNumberVerifiedAt: null.NewTimeDefaultAsNull(verifiedAt),
	}

	locations := make([]Location, 0, len(r.Access))
	for _, a := range r.Access {

		if a.LocationID == nil {
			return nil, nil
		}

		if len(a.Roles) == 0 {
			return nil, nil
		}

		locationId, err := a.LocationID.UUID()
		if err != nil {
			return nil, werror.Wrap(err, "error parsing locationID as uuid").Add("locationID", a.LocationID)
		}

		roles := make([]Role, 0, len(a.Roles))
		//range over all roles the user has
		for _, r := range a.Roles {

			//check if role is in the list of possibleRoles
			role, ok := possibleRoles[r]
			if !ok {
				continue
			}

			roles = append(roles, role)
		}

		locations = append(locations, Location{
			LocationID: locationId,
			Roles:      roles,
		})
	}

	userAccess.Locations = locations

	return userAccess, nil
}

// DeleteUserAccess removes a user from a location using the v2 client
func (a *Auth) DeleteUserAccess(ctx context.Context, locationID, userID uuid.UUID) error {

	_, err := a.v2.DeleteUserAccess(ctx, &authv2proto.DeleteUserAccessRequest{
		LocationID: wgrpcproto.UUIDProto(locationID),
		UserID:     wgrpcproto.UUIDProto(userID),
	})
	if err != nil {
		status, ok := status.FromError(err)
		if ok && status.Code() == codes.PermissionDenied {
			return wiggum.NotAuthorizedError.Here("not authorized to update user profile")
		}
		return werror.Wrap(err, "unable to delete user from location").
			Add("userID", userID.String()).
			Add("locationID", locationID.String())
	}

	return nil
}

// AddOrReplaceUserAccess adds access for a user to a loaction using the v2 client
func (a *Auth) AddOrReplaceUserAccess(ctx context.Context, userID uuid.UUID, locationID uuid.UUID, roles []int32) error {
	_, err := a.v2.AddOrReplaceUserAccess(ctx, &authv2proto.AddUserAccessRequest{
		UserID: wgrpcproto.UUIDProto(userID),
		Access: &authv2proto.UserAccess{
			LocationID: wgrpcproto.UUIDProto(locationID),
			Roles:      roles,
		},
	})
	if err != nil {
		status, ok := status.FromError(err)
		if ok && status.Code() == codes.PermissionDenied {
			return wiggum.NotAuthorizedError.Here("not authorized to update user profile")
		}
		return werror.Wrap(err, "unable to add user to location").
			Add("userID", userID.String()).
			Add("locationID", locationID.String())
	}

	return nil
}

// CreateUserProfile uses the v2 client to create a new user profile
func (a *Auth) CreateUserProfile(ctx context.Context, user CreateUpdateUserProfile) (uuid.UUID, error) {

	var verifiedAt *timestamp.Timestamp
	var err error

	if user.MobileNumberVerifiedAt.Valid {
		verifiedAt, err = ptypes.TimestampProto(user.MobileNumberVerifiedAt.Time)
		if err != nil {
			return uuid.UUID{}, werror.Wrap(err, "error parsing MobileNumberVerifiedAt")
		}
	}

	u, err := a.v2.CreateUserProfile(ctx, &authv2proto.UserProfileRequest{
		Password:               user.Password,
		Username:               user.Username,
		FirstName:              user.FirstName,
		LastName:               user.LastName,
		Type:                   toGRPCUserType(user.Type),
		Status:                 authv2proto.StatusType(authv2proto.StatusType_value[user.Status]),
		MobileNumber:           user.MobileNumber,
		MobileNumberVerifiedAt: verifiedAt,
	})
	if err != nil {
		return uuid.UUID{}, werror.Wrap(err, "error Creating User Profile")
	}

	uID, err := wgrpcproto.UUID(u)
	if err != nil {
		return uuid.UUID{}, werror.Wrap(err, "error parsing user ID")
	}

	return uID, nil
}

// UpdateUserProfile uses the v2 client to update a user profile. The Password field is optional and will not be updated if it is not present
func (a *Auth) UpdateUserProfile(ctx context.Context, locationID, userID uuid.UUID, user CreateUpdateUserProfile) error {
	var verifiedAt *timestamp.Timestamp
	var err error

	if user.MobileNumberVerifiedAt.Valid {
		verifiedAt, err = ptypes.TimestampProto(user.MobileNumberVerifiedAt.Time)
		if err != nil {
			return werror.Wrap(err, "error parsing MobileNumberVerifiedAt")
		}
	}

	req := &authv2proto.UpdateUserProfileRequest{
		UserID:     wgrpcproto.UUIDProto(userID),
		LocationID: wgrpcproto.UUIDProto(locationID),
		UserProfile: &authv2proto.UserProfileRequest{
			Password:               user.Password,
			Username:               user.Username,
			FirstName:              user.FirstName,
			LastName:               user.LastName,
			Type:                   toGRPCUserType(user.Type),
			Status:                 authv2proto.StatusType(authv2proto.StatusType_value[user.Status]),
			MobileNumber:           user.MobileNumber,
			MobileNumberVerifiedAt: verifiedAt,
		},
	}

	_, err = a.v2.UpdateUserProfile(ctx, req)
	if err != nil {
		status, ok := status.FromError(err)
		if ok && status.Code() == codes.PermissionDenied {
			return wiggum.NotAuthorizedError.Here("not authorized to update user profile")
		}
		return werror.Wrap(err, "error Updating User Profile").Add("userID", user.UserID.String())
	}

	return nil
}

//List Roles fetches all the possible roles
func (a *Auth) ListRoles(ctx context.Context) ([]Role, error) {

	r, err := a.v2.ListRoles(ctx, &empty.Empty{})
	if err != nil {
		return nil, werror.Wrap(err)
	}

	roles := make([]Role, 0, len(r.Roles))
	for _, v := range r.Roles {
		roles = append(roles, Role{
			ID:   v.RoleID,
			Name: v.Name,
			Type: v.Type,
		})
	}

	return roles, nil
}

func (a *Auth) MapOfRoles(ctx context.Context) (map[int32]Role, error) {
	dbRoles, err := a.ListRoles(ctx)
	if err != nil {
		return nil, werror.Wrap(err, "error listing possible Roles")
	}

	//turn dbRoles slice into map
	rolesMap := make(map[int32]Role)
	for _, v := range dbRoles {
		rolesMap[v.ID] = v
	}
	return rolesMap, nil
}
func mapFromProtoLocation(in []*authproto.Location, possibleRoles map[int32]Role) ([]Location, error) {

	l := make([]Location, 0, len(in))
	for _, v := range in {
		if v.LocationID == "" {
			return nil, nil
		}

		if len(v.Roles) == 0 {
			return nil, nil
		}

		lid, err := uuid.Parse(v.LocationID)
		if err != nil {
			return nil, werror.Wrap(err, "error parsing locationID as uuid")
		}

		roles := make([]Role, 0, len(v.Roles))
		//range over all roles the user has
		for _, v := range v.Roles {

			//check if role is in the list of possibleRoles
			role, ok := possibleRoles[v]
			if !ok {
				continue
			}

			roles = append(roles, role)
		}

		l = append(l, Location{
			LocationID: lid,
			Roles:      roles,
			Accepted:   v.Accepted,
		})
	}

	return l, nil
}

//LocationUser fetches a user and it's roles for that location
func (a *Auth) LocationUser(ctx context.Context, locationID, userID uuid.UUID) (LocationUser, error) {

	user, err := a.v2.User(ctx, wgrpcproto.UUIDProto(userID))
	if err != nil {

		return LocationUser{}, werror.Wrap(err, "error fetching User")
	}

	possibleRoles, err := a.MapOfRoles(ctx)
	if err != nil {
		return LocationUser{}, werror.Wrap(err, "error fetching map of possible roles")
	}

	//populate locationuser
	l, err := populateLocationUser(locationID, possibleRoles, user)
	if err != nil {
		return LocationUser{}, werror.Wrap(err, "error populating LocationUser")
	}

	return l, nil
}

func populateLocationUser(locationID uuid.UUID, possibleRoles map[int32]Role, user *authv2proto.UserResponse) (LocationUser, error) {
	var l LocationUser

	//Loop over all locations the user has access to
	for _, v := range user.Access {
		lid, err := wgrpcproto.UUID(v.LocationID)
		if err != nil {
			return LocationUser{}, werror.Wrap(err, "error parsing locationID")
		}

		//check if user has access on locationID passed in
		if lid == locationID {

			uid, err := wgrpcproto.UUID(user.Profile.UserID)
			if err != nil {
				return LocationUser{}, werror.Wrap(err, "error parsing userID")
			}

			var verifiedAt time.Time
			if user.Profile.MobileNumberVerifiedAt != nil {
				verifiedAt, err = ptypes.Timestamp(user.Profile.MobileNumberVerifiedAt)
				if err != nil {
					return LocationUser{}, werror.Wrap(err, "error parsing MobileNumberVerifiedAt")
				}
			}
			l.UserProfile.FirstName = user.Profile.FirstName
			l.UserProfile.LastName = user.Profile.LastName
			l.UserProfile.Username = user.Profile.Username
			l.UserProfile.UserID = uid
			l.UserProfile.Type = fromGRPCUserType(user.Profile.Type)
			l.UserProfile.Status = authv2proto.StatusType_name[int32(user.Profile.Status)]
			l.UserProfile.MobileNumber = user.Profile.MobileNumber
			l.UserProfile.MobileNumberVerifiedAt = null.NewTimeDefaultAsNull(verifiedAt)

			roles := make([]Role, 0, len(v.Roles))
			//range over all roles the user has
			for _, v := range v.Roles {

				//check if role is in the list of possibleRoles
				role, ok := possibleRoles[v]
				if !ok {
					continue
				}

				roles = append(roles, role)
			}

			l.Roles = roles

		}
	}

	//return error if not Roles found (user has no access)
	if len(l.Roles) == 0 {
		return LocationUser{}, werror.New("user does not have access to location").Add("locationID", locationID)
	}

	return l, nil
}

// RefreshToken
func (a *Auth) RefreshToken(ctx context.Context, token, exp string) (string, error) {

	r, err := a.s.RefreshToken(ctx, &authproto.RefreshTokenRequest{
		Token:      token,
		Expiration: exp,
	})
	if err != nil {
		return "", werror.Wrap(err, "unable to refresh token")
	}

	return r.Token, nil
}

// AcceptInvite -
func (a *Auth) AcceptInvite(ctx context.Context, locationID uuid.UUID, user UserProfile) error {

	_, err := a.v2.AcceptInvite(ctx, &authv2proto.AcceptInviteRequest{
		LocationID: wgrpcproto.UUIDProto(locationID),
		UserID:     wgrpcproto.UUIDProto(user.UserID),
		UserProfile: &authv2proto.UserProfileRequest{
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Username:  user.Username,
			Type:      toGRPCUserType(user.Type),
			Status:    authv2proto.StatusType(authv2proto.StatusType_value[user.Status]),
		},
	})
	if err != nil {
		return werror.Wrap(err, "unable to accept invite").Add("userID", user.UserID.String())
	}

	return nil
}

//Login returns the token as a string -- can pass nil for clientMetadata
func (a *Auth) InternalOauthLogin(ctx context.Context, username, firstName, lastName, googleAccessToken string, clientMetadata *ClientMetadata) (string, error) {

	in := &authproto.OauthLoginRequset{
		Username:          username,
		GoogleAccessToken: googleAccessToken,
		FirstName:         firstName,
		LastName:          lastName,
	}

	//Add ClientMetadata if its populated
	if clientMetadata != nil {
		//Shove all the request headers into a map
		headers := make(map[string]string, len(clientMetadata.RequestHeaders))
		for k, v := range clientMetadata.RequestHeaders {
			headers[k] = strings.Join(v, "|")
		}
		in.ClientMetadata = &authproto.ClientMetadata{
			ClientType:        clientMetadata.ClientType,
			RequestingService: clientMetadata.RequestingService,
			RequestHeaders:    headers,
		}
	}

	r, err := a.s.InternalOauthLogin(ctx, in)
	if err != nil {
		return "", werror.Wrap(err, "unable to do internal oauth login").Add("username", username)
	}

	return r.WeaveToken, nil
}
