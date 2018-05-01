package wvaultauth

import (
	"context"
	"os"
	"sync"
	"time"

	"weavelab.xyz/wlib/werror"
	"weavelab.xyz/wlib/wlog"
	"weavelab.xyz/wlib/wmetrics"
	"weavelab.xyz/wlib/wvault"
)

const (
	ConfigVaultToken = "VAULT_TOKEN"

	vaultAuthMetric = "vaultAuth"
)

func init() {
	wmetrics.SetLabels(vaultAuthMetric, "authType", "action")
}

type token struct {
	sync.Mutex
	client *wvault.Client

	leaseID  string
	accessor string
	token    string
	role     string

	renewable        bool
	expiration       time.Time
	requestIncrement time.Duration
}

// Auth returns a vault token associated with the given role
func New(ctx context.Context, c *wvault.Client, role string) (wvault.Secret, error) {

	// check to see if we have a vault token
	envToken := os.Getenv(ConfigVaultToken)
	if envToken != "" {
		wlog.Info("[VaultAuth] Using auth token from environment variable")
		t := token{
			token:      envToken,
			role:       role,
			client:     c,
			renewable:  true, // tokens are valid until they are revoked, a new token must be issued and configured
			expiration: time.Now().AddDate(1000, 0, 0),
		}
		return &t, nil
	}

	// if there is no token, fallback to kubernetes auth
	t, err := kubernetesLogin(ctx, c, role)
	if err != nil {
		return nil, werror.Wrap(err)
	}

	return t, nil
}

func (t *token) Name() string {
	return "Auth"
}

func (t *token) Token() string {

	t.Lock()
	defer t.Unlock()

	return t.token
}

// Renews the auth token. If recreate is true or the token
// cannot be refreshed, a new token is obtained.
// Returns true if a fresh was possible, otherwise returns false
// to indicate that a new token was created
func (t *token) Refresh(ctx context.Context, recreate bool) (bool, error) {
	t.Lock()
	defer t.Unlock()

	if t.renewable == false || recreate {
		err := t.recreate(ctx)
		if err != nil {
			return false, werror.Wrap(err)
		}
		return false, nil
	}

	// TODO: increment should be generated in a more robust way
	increment := t.requestIncrement
	resp, retryable, err := t.client.RenewToken(ctx, t, increment)
	if err != nil {
		if retryable {
			return false, werror.Wrap(err)
		}

		err := t.recreate(ctx)
		if err != nil {
			return false, werror.Wrap(err)
		}
		return false, nil
	}

	t.expiration = resp.NewExpiration()

	return true, nil
}

func (t *token) recreate(ctx context.Context) error {

	newToken, err := New(ctx, t.client, t.role)
	if err != nil {
		return werror.Wrap(err)
	}

	t = newToken.(*token)
	return nil // token was renewed
}

func (t *token) Expiration() time.Time {
	t.Lock()
	defer t.Unlock()

	return t.expiration
}

func (t *token) LeaseID() string {
	t.Lock()
	defer t.Unlock()

	return t.leaseID
}

func (t *token) Parent() wvault.Secret {
	return nil
}
