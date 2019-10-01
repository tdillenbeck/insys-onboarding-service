package mock

import (
	"context"

	"weavelab.xyz/monorail/shared/grpc-clients/client-grpc-clients/authclient"
	"weavelab.xyz/monorail/shared/wlib/uuid"

	"weavelab.xyz/monorail/shared/grpc-clients/client-grpc-clients/authclient/authproto"
)

type Auth struct {
	UserLocationsFn func(ctx context.Context, userID uuid.UUID) (*authclient.UserAccess, error)
}

func (a *Auth) UserLocations(ctx context.Context, userID uuid.UUID) (*authclient.UserAccess, error) {
	return a.UserLocationsFn(ctx, userID)
}

// Users is deprecated and is used by hydrators to get basic user data
func (a *Auth) Users(ctx context.Context, userIDs []uuid.UUID) ([]authclient.User, error) {
	return nil, nil
}

// User is deprecated and is used by hydrators to get basic user data
func (a *Auth) User(ctx context.Context, userID uuid.UUID) (authclient.User, error) {
	return authclient.User{}, nil
}

// UpdateLegacyProfile - uses legacy proto to update user profile
func (a *Auth) UpdateLegacyProfile(ctx context.Context, user authclient.UserProfile, password string) error {
	return nil
}

// UserProfile uses the version 2 client to get the profile for a userID. It returns false if a user with that ID doesn't exist.
func (a *Auth) UserProfiles(ctx context.Context, userIDs []uuid.UUID) ([]authclient.UserProfile, error) {
	return nil, nil
}

// UserProfile uses the version 2 client to get the profile for a userID. It returns false if a user with that ID doesn't exist.
func (a *Auth) UserProfile(ctx context.Context, userID uuid.UUID) (authclient.UserProfile, bool, error) {
	return authclient.UserProfile{}, true, nil
}

// SearchUserProfiles takes locationID, and search term to search all users across the system
func (a *Auth) SearchUserProfiles(ctx context.Context, locations []string, searchTerm string, limit, skip int) ([]authclient.UserProfile, error) {
	return nil, nil
}

// UserProfileByUsername uses the version 2 client to get a profile for a username. It returns false if a user with that username doesn't exist.
func (a *Auth) UserProfileByUsername(ctx context.Context, username string) (authclient.UserProfile, bool, error) {
	return authclient.UserProfile{}, true, nil
}

//Login returns the token as a string -- can pass nil for clientMetadata
func (a *Auth) Login(ctx context.Context, username, password string, expiration, refreshwindow int32, clientMetadata *authclient.ClientMetadata) (string, error) {
	return "", nil
}

//PublicKey returns the public Key as a byte array
func (a *Auth) PublicKey(ctx context.Context) ([]byte, error) {
	return nil, nil
}

//LocationUsers returns all users associated with the location and their ACL's for that location
func (a *Auth) LocationUsers(ctx context.Context, locationID uuid.UUID) ([]authclient.UserAccess, error) {
	return nil, nil
}

// DeleteUserAccess removes a user from a location using the v2 client
func (a *Auth) DeleteUserAccess(ctx context.Context, locationID, userID uuid.UUID) error {
	return nil
}

// AddOrReplaceUserAccess adds access for a user to a loaction using the v2 client
func (a *Auth) AddOrReplaceUserAccess(ctx context.Context, userID uuid.UUID, locationID uuid.UUID, roles []int32) error {
	return nil
}

// CreateUserProfile uses the v2 client to create a new user profile
func (a *Auth) CreateUserProfile(ctx context.Context, user authclient.CreateUpdateUserProfile) (uuid.UUID, error) {
	return uuid.NewV4(), nil
}

// UpdateUserProfile uses the v2 client to update a user profile. The Password field is optional and will not be updated if it is not present
func (a *Auth) UpdateUserProfile(ctx context.Context, locationID, userID uuid.UUID, user authclient.CreateUpdateUserProfile) error {
	return nil
}

//List Roles fetches all the possible roles
func (a *Auth) ListRoles(ctx context.Context) ([]authclient.Role, error) {
	return nil, nil
}

func (a *Auth) MapOfRoles(ctx context.Context) (map[int32]authclient.Role, error) {
	return nil, nil
}

func mapFromProtoLocation(in []*authproto.Location, possibleRoles map[int32]authclient.Role) ([]authclient.Location, error) {
	return nil, nil
}

//LocationUser fetches a user and it's roles for that location
func (a *Auth) LocationUser(ctx context.Context, locationID, userID uuid.UUID) (authclient.LocationUser, error) {
	return authclient.LocationUser{}, nil
}

// RefreshToken
func (a *Auth) RefreshToken(ctx context.Context, token, exp string) (string, error) {
	return "", nil
}

// AcceptInvite -
func (a *Auth) AcceptInvite(ctx context.Context, locationID uuid.UUID, user authclient.UserProfile) error {
	return nil
}

//Login returns the token as a string -- can pass nil for clientMetadata
func (a *Auth) InternalOauthLogin(ctx context.Context, username, firstName, lastName, googleAccessToken string, clientMetadata *authclient.ClientMetadata) (string, error) {
	return "", nil
}
