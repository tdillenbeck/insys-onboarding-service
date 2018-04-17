package wiggum

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	//Vendor
	jwt "github.com/dgrijalva/jwt-go"
	"weavelab.xyz/wlib/uuid"
)

const ExpiredErrorMessage = "token is expired"

type TokenInterface interface {
	SetString(v string)
	String() string
	SetJWT(val *jwt.Token)
	JWT() *jwt.Token
	Error() error
	Valid() bool
	ValidSkipExpiration() bool
	Claims() map[string]interface{}
	UserID() uuid.UUID
	Username() string
	AclType() ACLType
	Expiration() float64
	HTTPStatus() int
	JSON() string
	Locations() []uuid.UUID
	LocationACLs(uuid.UUID) []Permission
	Can(locationId uuid.UUID, permission Permission) bool
}

type Token struct {
	jwtToken    *jwt.Token
	tokenString string
	err         error
	valid       bool
	expired     bool
	httpStatus  int
}

func (t *Token) SetString(val string) {
	t.tokenString = val
	return
}

func (t *Token) String() string {
	return t.tokenString
}

func (t *Token) SetJWT(val *jwt.Token) {
	t.jwtToken = val
	return
}

func (t *Token) JWT() *jwt.Token {
	return t.jwtToken
}

func (t *Token) Error() error {
	return t.err
}

func (t *Token) Valid() bool {
	if t.err != nil && strings.ToLower(t.err.Error()) != ExpiredErrorMessage {
		return false
	}
	return t.valid && !t.expired
}

func (t *Token) ValidSkipExpiration() bool {
	return t.valid
}

func (t *Token) Claims() map[string]interface{} {
	if !t.ValidSkipExpiration() {
		return make(map[string]interface{})
	}
	return t.jwtToken.Claims
}

func (t *Token) UserID() uuid.UUID {
	if !t.ValidSkipExpiration() {
		return uuid.UUID{}
	}

	if v, ok := t.jwtToken.Claims["user_id"].(string); ok {
		if u, err := uuid.Parse(v); err == nil {
			return u
		}
	}
	return uuid.UUID{}
}

func (t *Token) Username() string {
	if !t.ValidSkipExpiration() {
		return ""
	}
	if v, ok := t.jwtToken.Claims["username"].(string); ok {
		return v
	}
	return ""
}

func (t *Token) ACLType() ACLType {
	if !t.ValidSkipExpiration() {
		return ""
	}
	if v, ok := t.jwtToken.Claims["type"].(string); ok {
		return ACLType(v)
	}
	return ""
}

func (t *Token) Expiration() float64 {
	if !t.ValidSkipExpiration() {
		return float64(0)
	}
	if v, ok := t.jwtToken.Claims["exp"].(float64); ok {
		return v
	}
	return float64(0)
}

func (t *Token) ExpirationBuffer() float64 {
	if !t.ValidSkipExpiration() {
		return float64(0)
	}
	if v, ok := t.jwtToken.Claims["expBuffer"].(float64); ok {
		return v
	}
	return float64(0)
}

func (t *Token) HTTPStatus() int {
	return t.httpStatus
}

func (t *Token) JSON() string {
	if t.Valid() {
		b, err := json.Marshal(t.Claims())
		if err != nil {
			log.Printf("ERROR: unable to Marshal Claims:%v\n", t.Claims())
			return `{"error":"invalid_token"}`
		}
		return string(b)
	}
	return `{"error":"invalid_token"}`
}

func (t *Token) Locations() []uuid.UUID {
	return t.acls().Locations()
}

/*
	LocationACLs returns a slice of Permissions for a given location ID
*/
func (t *Token) LocationACLs(locationID uuid.UUID) []Permission {
	permissions, ok := t.acls()[locationID.String()]
	if !ok {
		return []Permission{}
	}

	return permissions
}

/*
	Can returns whether or not a user has a given permission on a location.
*/
func (t *Token) Can(locationID uuid.UUID, permission Permission) bool {
	// check all loccation acls if type is WeaveACLType
	if t.ACLType() == WeaveACLType {
		for _, aclSet := range t.ACLS() {
			for _, allowed := range aclSet {
				if allowed == permission {
					return true
				}
			}
		}
	}

	for _, allowed := range t.LocationACLs(locationID) {
		if allowed == permission {
			return true
		}
	}

	return false
}

func (t *Token) ACLS() ACL {
	return t.acls()
}

func (t *Token) acls() ACL {
	returnACLS := ACL{}
	interfaceACLS, ok := t.Claims()["ACLS"]
	if !ok {
		return returnACLS
	}
	claimACLS, ok := interfaceACLS.(map[string]interface{})
	if !ok {
		return returnACLS
	}

	// loop over every location
	for key, value := range claimACLS {
		// loop over each permission in location
		permissions, ok := value.([]interface{})
		aclArray := []Permission{}
		if ok {
			for _, permission := range permissions {
				v, ok := permission.(float64)
				if ok == false {
					continue
				}

				p := Permission(int(v))
				aclArray = append(aclArray, p)

			}
		}

		returnACLS[key] = aclArray
	}

	return returnACLS
}

// Note: returning a nil Token pointer will blow things up
func newToken(s string) (*Token, error) {
	t := &Token{tokenString: strings.TrimSpace(s)}
	err := t.parseJWT()
	// Allow expired tokens
	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError:
			if t.err.(*jwt.ValidationError).Errors == jwt.ValidationErrorExpired {
				t.valid = true
				t.expired = true
			}
		}
		return t, err
	}

	t.setHttpStatus()

	return t, nil
}

func (t *Token) parseJWT() error {
	t.jwtToken, t.err = jwt.Parse(t.tokenString, func(jt *jwt.Token) (interface{}, error) {
		return VerifyKey(), nil
	})
	if t.err != nil {
		return t.err
	}
	return nil
}

func (t *Token) setHttpStatus() {
	switch t.err.(type) {
	case nil: // no error
		if !t.jwtToken.Valid {
			t.valid = false
			t.httpStatus = http.StatusUnauthorized
		} else {
			t.valid = true
			t.httpStatus = http.StatusOK
		}
	case *jwt.ValidationError: // something was wrong during the validation
		vErr := t.err.(*jwt.ValidationError)
		switch vErr.Errors {
		case jwt.ValidationErrorExpired:
			// Valid, but expired token
			t.valid = true
			t.expired = true
			t.httpStatus = http.StatusUnauthorized
		default:
			t.valid = false
			t.httpStatus = http.StatusInternalServerError
		}
	default: // something else went wrong
		t.valid = false
		t.httpStatus = http.StatusInternalServerError
	}
	return
}

// ValidateLocation checks locationID is valid for token
func (t *Token) ValidateLocation(locationID uuid.UUID) bool {
	for _, l := range t.Locations() {
		if uuid.Equal(l, locationID) {
			return true
		}
	}
	return false
}
