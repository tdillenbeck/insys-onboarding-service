package wgrpcclient

import (
	"context"
	"crypto/tls"
	"net"
	"net/url"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"weavelab.xyz/monorail/shared/wlib/werror"
)

func loadTransportCredentials(ctx context.Context, target string) (grpc.DialOption, error) {
	isTLS, err := isTLSServer(ctx, target)
	if err != nil {
		return nil, werror.Wrap(err, "unable to detect if tls server")
	}

	if isTLS == false {
		return grpc.WithInsecure(), nil
	}

	cfg := &tls.Config{
		// TODO: verify certificates once certificate issuing is handled
		InsecureSkipVerify: true,
	}

	creds := credentials.NewTLS(cfg)

	opt := grpc.WithTransportCredentials(creds)
	return opt, nil
}

// isTLSServer returns whether or not the target machine
// is wanting TLS
func isTLSServer(ctx context.Context, target string) (bool, error) {
	// connect to target
	// see if it is trying to do TLS

	u, err := url.Parse(target)
	if err != nil {
		return false, werror.Wrap(err)
	}
	if u.Scheme == "dns" {
		// https://github.com/grpc/grpc/blob/master/doc/naming.md#name-syntax
		target = u.Path
	}
	target = strings.Trim(target, "/") // trim leading and/or trailed / from target if present

	_, _, err = net.SplitHostPort(target)
	if err != nil {
		return false, werror.Wrap(err, "unexpected host:port format").Add("target", target)
	}

	conn, err := net.Dial("tcp", target)
	if err != nil {
		return false, werror.Wrap(err, "unable to dial")
	}
	defer conn.Close()

	config := &tls.Config{InsecureSkipVerify: true}
	tConn := tls.Client(conn, config)
	defer tConn.Close()

	err = tConn.Handshake()
	if err != nil {
		// if handshake was unsuccessful, assume not TLS capable
		return false, nil
	}

	return true, nil

}
