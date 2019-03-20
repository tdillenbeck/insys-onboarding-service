package net

import (
	"context"
	"fmt"
	"net"
)

func (r *Resolver) dial(ctx context.Context, network, server string) (net.Conn, error) {
	return nil, fmt.Errorf("not implemented")
}

func (r *Resolver) lookupSRV(ctx context.Context, service, proto, name string) (string, []*net.SRV, error) {
	return "", nil, fmt.Errorf("not implemented")
}
