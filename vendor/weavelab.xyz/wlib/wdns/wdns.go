package wdns

import (
	"net"
	"strconv"
	"strings"

	"weavelab.xyz/wlib/werror"
	"weavelab.xyz/wlib/wlog"
)

var defaultLookupSRV = net.LookupSRV

//ResolveAddress resolves an address
func ResolveAddress(serviceAddr string) (string, error) {
	//Split the port and host so we can see if we need to do an SRV lookup
	host, port, err := net.SplitHostPort(serviceAddr)
	if err != nil {
		return "", werror.Wrap(err, "could not split serviceAddr into host and port").Add("serviceAddr", serviceAddr)
	}

	if strings.HasSuffix(host, ".local") {
		wlog.WError(werror.New("warning: forcing hostname to be FQDN by adding '.' on the end").Add("original hostname", host))
		host = host + "."
	}

	//If string can be successfully converted to int then we don't need to do SRV lookup
	_, err = strconv.Atoi(port)
	if err == nil {
		return serviceAddr, nil
	}

	_, addrs, err := defaultLookupSRV(port, "tcp", host)
	if err != nil {
		return "", werror.Wrap(err, "could not look up SRV record")
	}

	//Make sure there is at least one addr
	if len(addrs) < 1 {
		return "", werror.New("Lookup did not return any SRV records").Add("addr", serviceAddr)
	}

	p := int(addrs[0].Port)

	if p == 0 {
		return "", werror.New("Invalid SRV port").Add("port", p)
	}

	port = strconv.Itoa(p)

	return net.JoinHostPort(host, port), nil
}
