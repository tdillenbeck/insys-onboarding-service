package wiggum

/*
  Manage a pool of valid keys, update keyset every N hours
  to pull in new keys.

  auth-api will make a new key available for N hours before
  starting to use the new key.
*/

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/url"
	"strings"
	"sync"
	"time"

	"gopkg.in/square/go-jose.v2"
	"weavelab.xyz/monorail/shared/wlib/werror"
	"weavelab.xyz/monorail/shared/wlib/whttp/whttpclient"
	"weavelab.xyz/monorail/shared/wlib/wlog"
)

type KeySet struct {
	sync.RWMutex // could be replaced with atomic pointer swap on keys
	keys         jose.JSONWebKeySet

	cookieName string
}

// NewPollingKeySetAsDefault is the replacement for wiggum.Register
func NewPollingKeySetAsDefault(ctx context.Context, authAPIAddr string, cookie string) (*KeySet, error) {
	ks, err := NewPollingKeySet(ctx, authAPIAddr, cookie)
	if err != nil {
		return nil, werror.Wrap(err)
	}

	Default = ks

	return ks, nil
}

// NewPollingKeySet loads public verifying keys periodically from the auth-api
// It will attempt to load the keys once and will return an error if
// unsuccessful. It will periodically load new keys every 80-90 minutes.
func NewPollingKeySet(ctx context.Context, authAPIAddr string, cookie string) (*KeySet, error) {

	k := KeySet{
		cookieName: cookie,
	}

	if k.cookieName == "" {
		k.cookieName = "wiggum"
	}

	err := k.pollKeySet(ctx, authAPIAddr)
	if err != nil {
		return nil, werror.Wrap(err, "unable to load keys")
	}

	return &k, nil
}

func (k *KeySet) AddKey(key jose.JSONWebKey) error {
	k.Lock()
	defer k.Unlock()

	k.keys.Keys = append(k.keys.Keys, key)

	return nil
}

func (k *KeySet) key(keyID string) (jose.JSONWebKey, error) {
	k.RLock()
	defer k.RUnlock()

	keys := k.keys.Key(keyID)
	if len(keys) == 0 {
		return jose.JSONWebKey{}, werror.New("key not found")
	}

	return keys[0], nil

}

/*
	PublicKeys returns all of the public keys in the KeySet
*/
func (k *KeySet) PublicKeys() jose.JSONWebKeySet {
	k.RLock()
	defer k.RUnlock()

	keys := make([]jose.JSONWebKey, 0, len(k.keys.Keys))

	for _, v := range k.keys.Keys {
		if v.IsPublic() == false {
			continue
		}
		keys = append(keys, v)
	}

	return jose.JSONWebKeySet{Keys: keys}
}

func (k *KeySet) pollKeySet(ctx context.Context, authAPIAddr string) error {

	var ksURL *url.URL
	if strings.Contains(authAPIAddr, "/") {
		var err error
		ksURL, err = url.Parse(authAPIAddr)
		if err != nil {
			return werror.Wrap(err, "invalid auth api addr")
		}
	} else {
		ksURL = &url.URL{
			Host: authAPIAddr,
		}
	}

	// set default scheme if not set
	if ksURL.Scheme == "" {
		ksURL.Scheme = "http"
	}

	// set default path if not set
	if ksURL.Path == "" {
		ksURL.Path = ".well-known/jwks.json"
	}

	loadAndSet := func() error {
		// load keys
		keys, err := loadKeySet(ctx, ksURL.String())
		if err != nil {
			return werror.Wrap(err)
		}
		// set the keys
		k.setKeys(keys)
		return nil
	}

	// must load keys a first time successfully, otherwise return a error
	err := loadAndSet()
	if err != nil {
		return werror.Wrap(err)
	}

	go func() {
		for {
			// sleep for some random amount of time between 80 and 90 minutes
			time.Sleep(time.Minute*80 + time.Duration(rand.Intn(int(time.Minute*10))))
			err := loadAndSet()
			if err != nil {
				wlog.WErrorC(ctx, werror.Wrap(err))
				continue
			}
		}
	}()

	return nil
}

func (k *KeySet) setKeys(keys *jose.JSONWebKeySet) {
	k.Lock()
	defer k.Unlock()

	k.keys = *keys
}

func loadKeySet(ctx context.Context, ksURL string) (*jose.JSONWebKeySet, error) {
	resp, err := whttpclient.Get(ctx, ksURL)
	if err != nil {
		return nil, werror.Wrap(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return nil, werror.New("unexpected response status code").Add("code", resp.StatusCode)
	}

	var ks jose.JSONWebKeySet
	err = json.NewDecoder(resp.Body).Decode(&ks)
	if err != nil {
		return nil, werror.Wrap(err)
	}

	return &ks, nil
}
