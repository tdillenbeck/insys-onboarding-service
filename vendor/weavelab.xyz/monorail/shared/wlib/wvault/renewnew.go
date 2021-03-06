package wvault

import (
	"time"
)

func renewPeriod(s Secret) time.Duration {

	p := s.Parent()
	var parentPeriod time.Duration
	if p != nil {
		parentPeriod = renewPeriod(p)
	}

	exp := s.Expiration()

	period := Until(exp) / 5

	if parentPeriod > 0 && parentPeriod < period {
		return parentPeriod
	}

	return period

}
