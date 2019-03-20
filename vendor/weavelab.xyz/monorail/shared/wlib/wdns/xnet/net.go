package net

import (
	"errors"
)

// Various errors contained in DNSError.
var (
	errNoSuchHost = errors.New("no such host")
)
