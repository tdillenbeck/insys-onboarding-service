package wiggum

import (
	"context"

	"weavelab.xyz/wlib/uuid"
)

type key int

const (
	tokenKey      = key(iota)
	locationIDKey = "locationid"
)

/*
	ContextAddToken parses and validates the token string t and adds the
	Token to the ctx if it is valid. Use ContextToken(ctx) to retrieve the Token.
*/
func ContextAddTokenString(ctx context.Context, t string) (context.Context, *Token, error) {
	// validate that the token is validate before adding it to the context
	tok, err := ParseAndValidate(t)
	if err != nil {
		return nil, nil, err
	}

	ctx = ContextAddToken(ctx, tok)

	return ctx, tok, err

}

/*
	ContextAddValidatedToken adds Token to the context.
	It is assumed that the token has already been validated.
	Use ContextToken(ctx) to retrieve the Token.
*/
func ContextAddToken(ctx context.Context, tok *Token) context.Context {

	ctx = context.WithValue(ctx, tokenKey, tok)

	return ctx
}

/*
	ContextAddLocationID adds locationid to the context.
*/
func ContextAddLocationID(ctx context.Context, l uuid.UUID) context.Context {

	ctx = context.WithValue(ctx, locationIDKey, l)

	return ctx
}

/*
	ContextToken retrieves the Token from the context.
	If no Token is found, nil, false is returned
*/
func ContextToken(ctx context.Context) (*Token, bool) {

	t, ok := ctx.Value(tokenKey).(*Token)
	if !ok {
		return t, false
	}

	return t, true
}

/*
	ContextCan looks up the Permission in the context and returns
	true or false if the user has permission for the given location.
*/
func ContextCan(ctx context.Context, locationID uuid.UUID, acl Permission) bool {

	t, ok := ContextToken(ctx)
	if !ok {
		return false
	}

	return t.Can(locationID, acl)
}

/*
	ContextUserID returns the user id or "", false if no token.
*/
func ContextUserID(ctx context.Context) (uuid.UUID, bool) {

	t, ok := ContextToken(ctx)
	if !ok {
		return uuid.UUID{}, false
	}

	return t.UserID(), true
}

/*
	ContextUserType returns the user type or "", false if no token.
*/
func ContextUserType(ctx context.Context) (string, bool) {

	t, ok := ContextToken(ctx)
	if !ok {
		return "", false
	}

	return string(t.ACLType()), true
}

/*
	ContextLocations returns a slice or locations or nil, false if no token.
*/
func ContextLocations(ctx context.Context) ([]uuid.UUID, bool) {

	t, ok := ContextToken(ctx)
	if !ok {
		return nil, false
	}

	return t.Locations(), true
}

/*
	ContextLocationID returns locationID, false if no value.
*/
func ContextLocationID(ctx context.Context) (uuid.UUID, bool) {
	l, ok := ctx.Value(locationIDKey).(uuid.UUID)
	if !ok {
		return uuid.UUID{}, false
	}

	return l, true
}
