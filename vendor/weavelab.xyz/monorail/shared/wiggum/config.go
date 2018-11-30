package wiggum

/*
	These values will need to either get passed
	when adding wiggum to middle ware
	or pulled from a config file
*/

const (
	pubKeyPath = "/Users/scott/.ssh/milhouse.rsa.pub"
)

var (
	verifyKeyVal []byte
	cookieName   string
)

func CookieName() string {
	return cookieName
}

func VerifyKey() []byte {
	return verifyKeyVal
}
