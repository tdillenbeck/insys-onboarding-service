// Package wvault provides methods for connecting to vault
// and retrieving secrets.
// This package is UNSTABLE and likely to change.
package wvault

import (
	"context"
	"os"
	"strings"
)

// https://www.vaultproject.io/docs/commands/environment.html
// kubectl port-forward --context=dev-ut -n ops vault-0 8200:8200

const (
	VaultAddressConfig = "VAULT_ADDR"
)

type Client struct {
	address string
}

// New returns an initialized wvault.Client
func New(ctx context.Context) (*Client, error) {

	address := os.Getenv(VaultAddressConfig)

	// check if the default address is still set,
	// if it is, update to use our default
	if address == "" {
		address = "http://vault.ops.svc.cluster.local.:8200"
	}

	address = strings.TrimSuffix(address, "/")

	v := Client{
		address: address,
	}

	return &v, nil
}

func (c *Client) BaseAddress() string {
	return c.address
}
