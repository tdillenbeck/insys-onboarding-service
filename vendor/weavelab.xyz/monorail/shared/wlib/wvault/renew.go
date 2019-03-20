package wvault

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"weavelab.xyz/monorail/shared/wlib/werror"
	"weavelab.xyz/monorail/shared/wlib/whttp/whttpclient"
	"weavelab.xyz/monorail/shared/wlib/wlog"
	"weavelab.xyz/monorail/shared/wlib/wlog/tag"
	"weavelab.xyz/monorail/shared/wlib/wmetrics"
)

const (
	vaultClientMetric = "vaultClient"
)

func init() {
	wmetrics.SetLabels(vaultClientMetric, "action", "result", "vaultType")
}

// AutoRenew will automatically attempt to renew the given Secret every renewInterval.
func (c *Client) AutoRenew(ctx context.Context, secret Secret) {

	wlog.InfoC(ctx, "[Vault] Starting auto renew loop")
	defer wlog.InfoC(ctx, "[Vault] Exiting auto renew loop")

	ctx, closeSpan := WithNewTracingParent(ctx)

	for {

		period := renewPeriod(secret) // determine the period to renew secret chain at
		wlog.InfoC(ctx, "[Vault] Next renew scheduled", tag.Duration("wait", period))
		closeSpan() // close existing span before sleeping

		select {
		case <-Clock.After(period):
		case <-ctx.Done():
			return
		}

		// open a new tracing span
		ctx, closeSpan = WithNewTracingParent(ctx)

		// loop through secret chain, if parent has to recreate,
		// force children to recreate also.
		_, err := autoRenew(ctx, secret)
		if err != nil {
			wlog.WErrorC(ctx, werror.Wrap(err, "[Vault] unable to renew secret"))
			continue
		}

	}
}

// autoRenew loops through all of the secrets in the chain,
// and renews as appropriate
func autoRenew(ctx context.Context, s Secret) (bool, error) {

	p := s.Parent()
	recreated := false
	if p != nil {
		var err error
		// test the parent secret to see if it needs
		// to be renewed
		recreated, err = autoRenew(ctx, p)
		if err != nil {
			return false, werror.Wrap(err)
		}
	}

	// TODO: define refreshIfSoonerThan better
	if recreated == false && s.ShouldRefresh() == false {
		wlog.InfoC(ctx, "[Vault] Does not need to be refreshed", tag.String("expiration", s.Expiration().Format(time.RFC3339)), tag.String("secret", s.LeaseID()))
		return false, nil
	}

	wlog.InfoC(ctx, "[Vault] Refreshing", tag.String("expiration", s.Expiration().Format(time.RFC3339)), tag.String("now", Clock.Now().Format(time.RFC3339)), tag.String("secret", s.LeaseID()), tag.Bool("forceRecreate", recreated))
	refreshed, err := s.Refresh(ctx, recreated)
	if err != nil {
		return false, werror.Wrap(err)
	}

	wlog.InfoC(ctx, "[Vault] Secret refreshed", tag.String("expiration", s.Expiration().Format(time.RFC3339)), tag.String("secret", s.LeaseID()), tag.Bool("refreshed", refreshed))

	recreate := !refreshed

	return recreate, nil

}

type RenewRequest struct {
	Increment int    `json:"increment"`          // number of seconds to extend
	LeaseID   string `json:"lease_id,omitempty"` // id string of secret to renew
}

type RenewResult struct {
	expiration    time.Time
	RequestID     string `json:"request_id"`
	LeaseID       string `json:"lease_id"`
	Renewable     bool   `json:"renewable"`
	LeaseDuration int    `json:"lease_duration"`
}

func (r *RenewResult) NewExpiration() time.Time {
	return r.expiration
}

func (c *Client) RenewToken(ctx context.Context, s Secret, increment time.Duration) (*RenewResult, bool, error) {
	path := "/v1/auth/token/renew-self"
	return c.renew(ctx, path, s, increment)
}

func (c *Client) RenewLease(ctx context.Context, s Secret, increment time.Duration) (*RenewResult, bool, error) {
	path := "/v1/sys/leases/renew"
	return c.renew(ctx, path, s, increment)
}

func (c *Client) renew(ctx context.Context, path string, s Secret, increment time.Duration) (*RenewResult, bool, error) {

	leaseID := s.LeaseID()

	wlog.InfoC(ctx, "[Vault] Renewing vault lease or token", tag.String("leaseID", leaseID), tag.Duration("increment", increment), tag.String("path", path))
	wmetrics.Incr(1, vaultClientMetric, "renew", "attempt", s.Name())

	r := RenewRequest{
		Increment: int(increment / time.Second),
		LeaseID:   leaseID,
	}

	buff := bytes.NewBuffer(nil)
	err := json.NewEncoder(buff).Encode(r)
	if err != nil {
		return nil, true, werror.Wrap(err)
	}

	url := c.BaseAddress() + path

	req, err := http.NewRequest(http.MethodPut, url, buff)
	if err != nil {
		return nil, true, werror.Wrap(err)
	}

	// add timeout to context
	ctx, done := context.WithTimeout(ctx, time.Second*30)
	defer done()

	req.Header.Add("X-Vault-Token", s.Token())

	resp, err := whttpclient.Do(ctx, req)
	if err != nil {
		return nil, true, werror.Wrap(err)
	}
	defer resp.Body.Close()

	// check for expired lease
	// 400 {"errors":["lease expired"]}
	// 400 {"errors":["lease not found or lease is not renewable"]}

	if resp.StatusCode >= 300 {
		body, _ := ioutil.ReadAll(resp.Body)
		code := werror.HttpToWeaveCode(resp.StatusCode)

		err := werror.New("lease renew error - unexpected status code").Add("code", resp.StatusCode).Add("body", string(body)).Add("url", url).SetCode(code)
		if resp.StatusCode >= 400 && resp.StatusCode < 500 {
			return nil, false, err
		} else {
			return nil, true, err
		}
	}

	tresult := AuthLoginResponse{}

	err = json.NewDecoder(resp.Body).Decode(&tresult)
	if err != nil {
		return nil, true, werror.Wrap(err)
	}

	renewable := tresult.Renewable || tresult.Auth.Renewable

	leaseDuration := tresult.LeaseDuration
	if leaseDuration < tresult.Auth.LeaseDuration {
		leaseDuration = tresult.Auth.LeaseDuration
	}

	result := RenewResult{
		expiration:    Clock.Now().Add(time.Second * time.Duration(leaseDuration)),
		RequestID:     tresult.RequestID,
		Renewable:     renewable,
		LeaseDuration: leaseDuration,
		LeaseID:       tresult.LeaseID,
	}

	wmetrics.Incr(1, vaultClientMetric, "renew", "success", s.Name())

	wlog.InfoC(ctx, "[Vault] Vault lease|token renewed", tag.String("leaseID", leaseID), tag.Int("duration", result.LeaseDuration))

	return &result, true, nil
}

type AuthLoginResponse struct {
	RequestID     string      `json:"request_id"`
	LeaseID       string      `json:"lease_id"`
	LeaseDuration int         `json:"lease_duration"`
	Renewable     bool        `json:"renewable"`
	Data          interface{} `json:"data"`
	Warnings      interface{} `json:"warnings"`
	Auth          struct {
		ClientToken string   `json:"client_token"`
		Accessor    string   `json:"accessor"`
		Policies    []string `json:"policies"`
		Metadata    struct {
			Role                     string `json:"role"`
			ServiceAccountName       string `json:"service_account_name"`
			ServiceAccountNamespace  string `json:"service_account_namespace"`
			ServiceAccountSecretName string `json:"service_account_secret_name"`
			ServiceAccountUID        string `json:"service_account_uid"`
		} `json:"metadata"`
		LeaseDuration int  `json:"lease_duration"`
		Renewable     bool `json:"renewable"`
	} `json:"auth"`
}
