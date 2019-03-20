package wvaultdb

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"

	"weavelab.xyz/monorail/shared/wlib/wapp"
	"weavelab.xyz/monorail/shared/wlib/werror"
	"weavelab.xyz/monorail/shared/wlib/whttp/whttpclient"
	"weavelab.xyz/monorail/shared/wlib/wlog"
	"weavelab.xyz/monorail/shared/wlib/wlog/tag"
	"weavelab.xyz/monorail/shared/wlib/wsql"
	"weavelab.xyz/monorail/shared/wlib/wvault"
	"weavelab.xyz/monorail/shared/wlib/wvault/wvaultauth"
	"weavelab.xyz/monorail/shared/wlib/wvault/wvaulttypes"
)

func New(ctx context.Context, role string, target string, s *wsql.Settings) (*wsql.PG, error) {

	// get database credentials lease
	c, err := createCredentials(ctx, role, target)
	if err != nil {
		return nil, werror.Wrap(err)
	}

	wapp.AddLivenessCheck(wapp.LivenessCheckFunc(c.alive))

	s.PrimaryConnectString.Username = c.Username()
	s.PrimaryConnectString.Password = c.Password()
	s.ReplicaConnectString.Username = c.Username()
	s.ReplicaConnectString.Password = c.Password()

	dbRole := strings.TrimPrefix(target, "db_")
	if s.PrimaryConnectString.Role == "" {
		s.PrimaryConnectString.Role = dbRole
	}
	if s.ReplicaConnectString.Role == "" {
		s.ReplicaConnectString.Role = dbRole
	}

	db, err := wsql.New(s)
	if err != nil {
		return nil, werror.Wrap(err)
	}

	c.setDB(db)

	go c.metricsLoop(ctx)

	return db, nil
}

// VaultDBCredentials looks up database credentials assigned to the given role
func createCredentials(ctx context.Context, role string, target string) (*Creator, error) {

	client, err := wvault.New(ctx)
	if err != nil {
		return nil, werror.Wrap(err)
	}

	ctx2, done := wvault.WithNewTracingParent(ctx)
	defer done()

	// get auth token
	authToken, err := wvaultauth.New(ctx2, client, role)
	if err != nil {
		return nil, werror.Wrap(err)
	}

	c := Creator{
		parent: authToken,
		client: client,
		target: target,
	}

	err = c.createCredentials(ctx2)
	if err != nil {
		return nil, werror.Wrap(err)
	}

	go client.AutoRenew(ctx, &c)

	return &c, nil
}

type Creator struct {
	sync.Mutex

	parent wvault.Secret
	client *wvault.Client
	db     refreshableCredentials

	leaseID          string
	expiration       time.Time
	requestIncrement time.Duration

	createdAt time.Time
	username  string
	password  string

	target string

	abandoned bool
}

func (c *Creator) Name() string {
	return "Database"
}

func (c *Creator) Parent() wvault.Secret {
	return c.parent
}

func (c *Creator) Expiration() time.Time {
	return c.expiration
}

func (c *Creator) LeaseID() string {
	return c.leaseID
}

func (c *Creator) Token() string {
	return c.Parent().Token()
}

func (c *Creator) Username() string {
	return c.username
}

func (c *Creator) Password() string {
	return c.password
}

func (c *Creator) ShouldRefresh() bool {
	lifetime := c.expiration.Sub(c.createdAt)

	halfway := c.createdAt.Add(lifetime / 2)

	now := wvault.Clock.Now()

	return now.After(halfway)

}

func (c *Creator) ShouldForceRecreate() bool {
	lifetime := c.expiration.Sub(c.createdAt)

	cutPoint := c.createdAt.Add(lifetime * 9 / 10)

	now := wvault.Clock.Now()

	return now.After(cutPoint)
}

func (c *Creator) createCredentials(ctx context.Context) error {

	wlog.InfoC(ctx, "[VaultDB] attempting to create database credentials")

	databaseSecretPath := `/v1/database/creds/` + c.target
	url := c.client.BaseAddress() + databaseSecretPath

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return werror.Wrap(err)
	}

	req.Header.Add("X-Vault-Token", c.Token())

	resp, err := whttpclient.Do(ctx, req)
	if err != nil {
		return werror.Wrap(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		body, _ := ioutil.ReadAll(resp.Body)
		return werror.New("unable to load database credentials").Add("code", resp.StatusCode).Add("Body", string(body)).Add("url", url).Add("target", c.target)
	}
	cr := wvaulttypes.CredentialsResponse{}
	err = json.NewDecoder(resp.Body).Decode(&cr)
	if err != nil {
		return werror.Wrap(err)
	}

	leaseDuration := time.Duration(cr.LeaseDuration) * time.Second
	expiration := wvault.Clock.Now().Add(leaseDuration)

	c.expiration = expiration
	c.leaseID = cr.LeaseID
	c.createdAt = wvault.Clock.Now()
	c.username = cr.Data.Username
	c.password = cr.Data.Password
	c.requestIncrement = leaseDuration

	wlog.InfoC(ctx, "[VaultDB] database credentials created", tag.String("username", cr.Data.Username), tag.Int("leaseDuration", cr.LeaseDuration), tag.String("leaseID", cr.LeaseID))
	return nil
}

func (c *Creator) renew(ctx context.Context) (time.Duration, bool, error) {

	c.Lock()
	defer c.Unlock()

	// if the increment is 0 recreate credentials
	// or if the increment is < 3/4 the request increment,
	// abandon these credentials and create a new set of credentials

	increment := c.requestIncrement
	resp, retryable, err := c.client.RenewLease(ctx, c, increment)
	if err != nil {
		return 0, retryable, werror.Wrap(err)
	}

	leaseDuration := time.Duration(resp.LeaseDuration) * time.Second

	c.expiration = wvault.Clock.Now().Add(leaseDuration)

	return leaseDuration, false, nil

}
