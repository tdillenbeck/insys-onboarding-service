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
	// "strings"

	// Weave
	jwt "weavelab.xyz/monorail/shared/wiggum/jwt-go"
	"weavelab.xyz/monorail/shared/wlib/uuid"
	"weavelab.xyz/monorail/shared/wlib/werror"
)

// Set required values
func Register(key []byte, cookie string) {
	if len(key) == 0 {
		key = []byte(`-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAuIrag84JwUnWwrG72T5s
wokwqynX3pzDfcM2tTDOg9Gp4h5YEHdmpKAMklcfIiPnDbXCOzybrb6nnABXeoXG
zlHstBJxzJWfcopI+wGJKnRyN+Z70F8ScZedNoaWsGdjSq6N8rcv0Vk8Fl2WG/2A
mLTdTBGwGAvu7QDlLRTSC3DmMW+6Xbw+q+tGZuNEBfZj3+p3rXUdTWNoD64XcUoq
bp7UHSBZcP9F4azRjMn+wYCO/mVzsUhp+b/PK/j2qt3xIePQLQdHTc6rCUEP6MGF
eX3LbWqkxhgpnYcw0049bEOdc8eDu8x+hSOV9peBJdrKhNWAuiyFH18u2Zig3VYY
RwIDAQAB
-----END PUBLIC KEY-----`)
	}

	verifyKeyVal = key

	if cookie == "" {
		cookieName = "wiggum"
	} else {
		cookieName = cookie
	}
	return
}

// Makes a new signed JWT token string
func Make(acls ACL, userID uuid.UUID, username string, aclT ACLType, exp int64, buffer int64, signKey []byte) (string, error) {
	t := jwt.New(jwt.GetSigningMethod("RS256"))

	t.Claims["ACLS"] = acls
	t.Claims["user_id"] = userID
	t.Claims["username"] = username
	t.Claims["type"] = aclT
	t.Claims["exp"] = exp
	t.Claims["expBuffer"] = buffer
	return t.SignedString(signKey)
}

// Gets the passed jwt string from either the request cookie or header
func TokenString(r *http.Request) string {
	headerValue, ok := getHeader(r)
	if ok {
		return headerValue
	}
	cookieValue, ok := getCookie(r)
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
	return newToken(tokenString)
}

func ParseSkipValidation(tokenString string) (*Token, error) {
	token, err := newToken(tokenString)
	if !token.ValidSkipExpiration() {
		return token, werror.Wrap(err).Add("context", "Invalid Token")
	}
	return token, nil
}
