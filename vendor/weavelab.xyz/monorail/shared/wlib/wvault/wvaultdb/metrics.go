package wvaultdb

import (
	"context"
	"time"

	"weavelab.xyz/monorail/shared/wlib/wmetrics"
	"weavelab.xyz/monorail/shared/wlib/wvault"
)

const (
	// some arbitrary number comparable to metrics polling interval
	metricsInterval = time.Second * 10

	metricTTL = "wvaultDBCredentialsTTL"
)

func (c *Creator) metricsLoop(ctx context.Context) {

	t := wvault.Clock.NewTicker(metricsInterval)
	defer t.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-t.Chan():
			c.pushMetrics()
		}
	}
}

func (c *Creator) pushMetrics() {

	exp := c.Expiration()
	ttl := int(wvault.Until(exp) / time.Second)

	wmetrics.Gauge(ttl, metricTTL)
}
