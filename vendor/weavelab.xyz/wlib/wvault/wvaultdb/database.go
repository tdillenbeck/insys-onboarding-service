package wvaultdb

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"

	"weavelab.xyz/wlib/wapp"
	"weavelab.xyz/wlib/werror"
	"weavelab.xyz/wlib/whttp/whttpclient"
	"weavelab.xyz/wlib/wlog"
	"weavelab.xyz/wlib/wlog/tag"
	"weavelab.xyz/wlib/wsql"
	"weavelab.xyz/wlib/wvault"
	"weavelab.xyz/wlib/wvault/wvaultauth"
	"weavelab.xyz/wlib/wvault/wvaulttypes"
)

var (
	maxRenewInterval time.Duration
)

func init() {
	i := os.Getenv("WVAULT_DB_MAX_RENEW_INTERVAL")
	maxRenewInterval, _ = time.ParseDuration(i)
}

func New(ctx context.Context, role string, target string, s *wsql.Settings) (*wsql.PG, error) {

	// get database credentials lease
	c, err := CreateCredentials(ctx, role, target)
	if err != nil {
		return nil, werror.Wrap(err)
	}

	wapp.AddLivenessCheck(wapp.LivenessCheckFunc(c.alive))

	s.PrimaryConnectString.Username = c.Username()
	s.PrimaryConnectString.Password = c.Password()
	s.ReplicaConnectString.Username = c.Username()
	s.ReplicaConnectString.Password = c.Password()

	db, err := wsql.New(s)
	if err != nil {
		return nil, werror.Wrap(err)
	}

	c.setDB(db)

	go c.metricsLoop(ctx)

	return db, nil
}

// VaultDBCredentials looks up database credentials assigned to the given role
func CreateCredentials(ctx context.Context, role string, target string) (*Creator, error) {

	client, err := wvault.New(ctx)
	if err != nil {
		return nil, werror.Wrap(err)
	}

	// get auth token
	authToken, err := wvaultauth.New(ctx, client, role)
	if err != nil {
		return nil, werror.Wrap(err)
	}

	c := Creator{
		parent: authToken,
		client: client,
		target: target,
	}

	err = c.createCredentials(ctx)
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

func (c *Creator) createCredentials(ctx context.Context) error {

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
		return werror.New("unable to load database credentials").Add("code", resp.StatusCode).Add("Body", string(body)).Add("url", url)
	}
	cr := wvaulttypes.CredentialsResponse{}
	err = json.NewDecoder(resp.Body).Decode(&cr)
	if err != nil {
		return werror.Wrap(err)
	}

	leaseDuration := time.Duration(cr.LeaseDuration) * time.Second
	expiration := time.Now().Add(leaseDuration)

	c.expiration = expiration
	c.leaseID = cr.LeaseID
	c.createdAt = time.Now()
	c.username = cr.Data.Username
	c.password = cr.Data.Password
	c.requestIncrement = leaseDuration

	wlog.Info("[VaultDB] database credentials created", tag.String("username", cr.Data.Username), tag.Int("leaseDuration", cr.LeaseDuration), tag.String("leaseID", cr.LeaseID))
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

	c.expiration = time.Now().Add(leaseDuration)

	return leaseDuration, false, nil

}
