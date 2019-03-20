package wvault

import (
	"time"

	"weavelab.xyz/monorail/shared/wlib/wvault/clockwork"
)

var Clock = clockwork.NewRealClock()

func Until(t time.Time) time.Duration {

	n := Clock.Now()

	return t.Sub(n)
}
