package wiggum

import (
	"weavelab.xyz/monorail/shared/wlib/werror"
)

/*
	These values will need to either get passed
	when adding wiggum to middle ware
	or pulled from a config file
*/

func CookieName() string {
	return Default.CookieName()
}

func (k *KeySet) CookieName() string {
	return k.cookieName
}

func VerifyKey(keyID string, alg string) (interface{}, error) {
	return Default.VerifyKey(keyID, alg)
}

func (k *KeySet) VerifyKey(keyID string, alg string) (interface{}, error) {
	key, err := k.key(keyID)
	if err != nil {
		return nil, werror.Wrap(err, "unable to find key")
	}

	if key.Algorithm != alg {
		return nil, werror.New("key mismatch")
	}

	if key.IsPublic() == false {
		return nil, werror.New("not a valid verify key")
	}

	return key.Key, nil
}
