package wiggum

/*
  Middleware Hook
	This is the function that gets called in the middleware stack
	It either passes the request through to the next request or
	rejects it returnig an error state or
	redirects to the redirect url
*/

import (
	"net/http"
	"time"

	// Weave
	jwt "weavelab.xyz/monorail/shared/wiggum/jwt-go"
	"weavelab.xyz/monorail/shared/wlib/uuid"
	"weavelab.xyz/monorail/shared/wlib/werror"
)

var Default = &KeySet{}

// Set required values
func Register(verifyKey []byte, cookie string) error {
	var err error
	Default, err = NewLegacyKeySet(verifyKey, cookie)
	if err != nil {
		return werror.Wrap(err)
	}

	return nil
}

// Makes a new signed JWT token string
func (k *KeySet) Make(acls ACL, userID uuid.UUID, username string, aclT ACLType, exp int64, buffer int64, keyID string, audience ...string) (string, error) {

	key, err := k.key(keyID)
	if err != nil {
		return "", werror.Wrap(err, "unable to lookup key")
	}

	if key.IsPublic() {
		return "", werror.Wrap(err, "unable to sign with public key")
	}

	signingMethod := jwt.GetSigningMethod(key.Algorithm)

	t := jwt.New(signingMethod)

	t.Claims["ACLS"] = acls
	t.Claims["user_id"] = userID
	t.Claims["username"] = username
	t.Claims["type"] = aclT
	if len(audience) > 0 {
		t.Claims["aud"] = audience
	}
	t.Claims["exp"] = exp
	t.Claims["expBuffer"] = buffer
	t.Claims["iat"] = time.Now().Unix()
	t.Claims["jti"] = uuid.NewV4()

	if keyID != "" {
		// add hint as to which signing key was used
		// https://tools.ietf.org/html/rfc7515#section-4.1.4
		t.Header["kid"] = keyID
	}

	ss, err := t.SignedString(key.Key)
	if err != nil {
		return "", werror.Wrap(err, "unable to sign")
	}

	return ss, nil
}

// Gets the passed jwt string from either the request cookie or header
func TokenString(r *http.Request) string {
	return Default.TokenString(r)
}

// Gets the passed jwt string from either the request cookie or header
func (k *KeySet) TokenString(r *http.Request) string {
	headerValue, ok := getHeader(r)
	if ok {
		return headerValue
	}
	cookieValue, ok := getCookie(r, k.CookieName())
	if ok {
		return cookieValue
	}
	r.ParseForm()
	paramValue, ok := getQuery(r.Form)
	if ok {
		return paramValue
	}
	return ""
}

// Convert a jwt string into a wiggum token
func ParseAndValidate(tokenString string) (*Token, error) {
	return Default.ParseAndValidate(tokenString)
}

// Convert a jwt string into a wiggum token
func (k *KeySet) ParseAndValidate(tokenString string) (*Token, error) {
	return k.newToken(tokenString)
}

func ParseSkipValidation(tokenString string) (*Token, error) {
	return Default.ParseSkipValidation(tokenString)
}

func (k *KeySet) ParseSkipValidation(tokenString string) (*Token, error) {
	token, err := k.newToken(tokenString)
	if !token.ValidSkipExpiration() {
		return token, werror.Wrap(err).Add("context", "Invalid Token")
	}
	return token, nil
}
