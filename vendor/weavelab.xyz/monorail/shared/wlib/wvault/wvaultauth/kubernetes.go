package wvaultauth

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
	"time"

	"weavelab.xyz/monorail/shared/wlib/werror"
	"weavelab.xyz/monorail/shared/wlib/whttp/whttpclient"
	"weavelab.xyz/monorail/shared/wlib/wlog"
	"weavelab.xyz/monorail/shared/wlib/wlog/tag"
	"weavelab.xyz/monorail/shared/wlib/wmetrics"
	"weavelab.xyz/monorail/shared/wlib/wvault"
)

const (
	kubernetesServiceTokenConfig = "KUBERNETES_SERVICE_TOKEN"
)

func kubernetesLogin(ctx context.Context, c *wvault.Client, role string) (*token, error) {

	wmetrics.Incr(1, vaultAuthMetric, "kubernetesAuth", "attempt")

	serviceToken, err := kubernetesServiceToken()
	if err != nil {
		return nil, werror.Wrap(err)
	}

	// ca.crt, namespace, and token are in that directory

	// use the kubernetes vault backend to use the kubernetes
	// service account credentials to get a vault token

	loginRequest := KubernetesLoginRequest{
		JWT:  serviceToken,
		Role: role,
	}

	buff := bytes.NewBuffer(nil)
	err = json.NewEncoder(buff).Encode(loginRequest)
	if err != nil {
		return nil, werror.Wrap(err)
	}

	url := c.BaseAddress() + "/v1/auth/kubernetes/login"

	resp, err := whttpclient.Post(ctx, url, "application/json", buff)
	if err != nil {
		return nil, werror.Wrap(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, werror.New("auth vault login - unexpected status code").Add("code", resp.StatusCode).Add("body", string(body)).Add("url", url).Add("role", role)
	}

	k := wvault.AuthLoginResponse{}

	err = json.NewDecoder(resp.Body).Decode(&k)
	if err != nil {
		return nil, werror.Wrap(err)
	}

	leaseDuration := time.Duration(k.Auth.LeaseDuration) * time.Second
	expiration := time.Now().Add(leaseDuration)

	token := token{
		client: c,

		leaseID:  k.LeaseID, // lease id is always empty
		accessor: k.Auth.Accessor,
		token:    k.Auth.ClientToken,

		renewable:        k.Auth.Renewable,
		expiration:       expiration,
		requestIncrement: leaseDuration,
	}

	wlog.Info("[VaultAuth] kubernetes auth", tag.Int("duration", k.Auth.LeaseDuration), tag.Bool("renewable", k.Auth.Renewable))

	wmetrics.Incr(1, vaultAuthMetric, "kubernetesAuth", "success")

	return &token, nil

}

func kubernetesServiceToken() (string, error) {

	// TODO: move service token logic to separate function
	serviceToken := []byte(os.Getenv(kubernetesServiceTokenConfig))
	if len(serviceToken) != 0 {
		return string(serviceToken), nil
	}

	serviceTokenFilename := `/var/run/secrets/kubernetes.io/serviceaccount/token`

	var err error
	serviceToken, err = ioutil.ReadFile(serviceTokenFilename)
	if err != nil {
		return "", werror.Wrap(err)
	}

	return string(serviceToken), nil
}

type KubernetesLoginRequest struct {
	JWT  string `json:"jwt"`
	Role string `json:"role"`
}
