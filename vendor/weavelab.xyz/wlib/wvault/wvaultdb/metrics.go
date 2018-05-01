package wvaultdb

import (
	"context"
	"time"

	"weavelab.xyz/wlib/wmetrics"
)

const (
	// some arbitrary number comparable to metrics polling interval
	metricsInterval = time.Second * 10

	metricTTL = "wvaultDBCredentialsTTL"
)

func (c *Creator) metricsLoop(ctx context.Context) {

	t := time.NewTicker(metricsInterval)
	defer t.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-t.C:
			c.pushMetrics()
		}
	}
}

func (c *Creator) pushMetrics() {

	exp := c.Expiration()
	ttl := int(time.Until(exp) / time.Second)

	wmetrics.Gauge(ttl, metricTTL)
}
