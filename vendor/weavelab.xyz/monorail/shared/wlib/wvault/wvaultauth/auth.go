package wvaultauth

import (
	"context"
	"os"
	"sync"
	"time"

	"weavelab.xyz/monorail/shared/wlib/werror"
	"weavelab.xyz/monorail/shared/wlib/wlog"
	"weavelab.xyz/monorail/shared/wlib/wmetrics"
	"weavelab.xyz/monorail/shared/wlib/wvault"
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
	createdAt        time.Time
	expiration       time.Time
	requestIncrement time.Duration
}

// Auth returns a vault token associated with the given role
func New(ctx context.Context, c *wvault.Client, role string) (wvault.Secret, error) {

	// check to see if we have a vault token
	envToken := os.Getenv(ConfigVaultToken)
	if envToken != "" {
		wlog.InfoC(ctx, "[VaultAuth] Using auth token from environment variable")
		now := wvault.Clock.Now()

		t := token{
			token:      envToken,
			role:       role,
			client:     c,
			renewable:  true, // tokens are valid until they are revoked, a new token must be issued and configured
			createdAt:  now,
			expiration: now.AddDate(1000, 0, 0),
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
// Returns true if a refresh was possible, otherwise returns false
// to indicate that a new token was created
// If a new token was issued, all secrets depending on the previous
// token should be recreated
func (t *token) Refresh(ctx context.Context, recreate bool) (bool, error) {

	return t.refresh(ctx, recreate)
}

func (t *token) refresh(ctx context.Context, recreate bool) (bool, error) {
	// if the token is expiring soon, force recreate
	if wvault.Until(t.Expiration()) < time.Second*10 {
		recreate = true
	}

	if t.Renewable() == false || recreate {
		err := t.recreate(ctx)
		if err != nil {
			return false, werror.Wrap(err)
		}
		return false, nil
	}

	// TODO: increment should be generated in a more robust way
	requestIncrement := t.RequestIncrement()
	resp, retryable, err := t.client.RenewToken(ctx, t, requestIncrement)
	if err != nil {
		if retryable {
			return false, werror.Wrap(err)
		}
		// log and continue
		wlog.WErrorC(ctx, werror.Wrap(err, "[Vault] unable to renew token"))
	}

	if resp == nil || resp.LeaseDuration < int(requestIncrement/time.Second)/2 {
		return t.Refresh(ctx, true)
	}

	t.expiration = resp.NewExpiration()
	return true, nil

}

func (t *token) recreate(ctx context.Context) error {

	t.Lock()
	client := t.client
	role := t.role
	t.Unlock()

	newToken, err := New(ctx, client, role)
	if err != nil {
		return werror.Wrap(err)
	}

	*t = *(newToken.(*token))

	return nil // token was renewed
}

func (t *token) RequestIncrement() time.Duration {
	t.Lock()
	defer t.Unlock()

	return t.requestIncrement
}

func (t *token) Renewable() bool {
	t.Lock()
	defer t.Unlock()

	return t.renewable
}

func (t *token) CreatedAt() time.Time {
	t.Lock()
	defer t.Unlock()

	return t.createdAt
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

func (t *token) ShouldRefresh() bool {
	// if we're past the halfway mark, we should refresh
	expiration := t.Expiration()
	createdAt := t.CreatedAt()

	lifetime := expiration.Sub(createdAt)

	halfway := createdAt.Add(lifetime / 2)

	now := wvault.Clock.Now()

	shouldRefresh := now.After(halfway)

	return shouldRefresh
}
