package wgrpcclient

import (
	"context"
	"crypto/tls"
	"net"

	"weavelab.xyz/monorail/shared/wlib/werror"
)

// isTLSServer returns whether or not the target machine
// is wanting TLS
func isTLSServer(ctx context.Context, target string) (bool, error) {
	// connect to target
	// see if it is trying to do TLS

	_, _, err := net.SplitHostPort(target)
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
