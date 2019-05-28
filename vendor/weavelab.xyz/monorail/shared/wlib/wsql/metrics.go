package wsql

import (
	"strconv"
	"strings"
	"time"

	"weavelab.xyz/monorail/shared/wlib/wmetrics"
)

const dbConnectionMetricPrefix string = "DBConnectionPool"

func init() {
	wmetrics.SetLabels(dbConnectionMetricPrefix, "PoolName", "DBHostname", "IsPrimary", "PoolID", "Stat")
}

func (p *DB) SendConnectionStatistics() {

	ticker := time.NewTicker(time.Second * 10)
	defer ticker.Stop()

	hostname := p.Hostname // database server hostname
	if hostname == "" {
		hostname = "unknown"
	}

	hostname = strings.Replace(hostname, ".", "-", -1)

	name := p.Name
	if name == "" {
		name = "unknown"
	}

	isPrimary := strconv.FormatBool(p.isPrimary)
loop:
	for {
		select {
		case <-ticker.C:
			s := p.Stats()

			wmetrics.Gauge(s.OpenConnections, dbConnectionMetricPrefix, name, hostname, isPrimary, p.poolID, "open")

			wmetrics.Gauge(s.InUse, dbConnectionMetricPrefix, name, hostname, isPrimary, p.poolID, "InUse") //int
			wmetrics.Gauge(s.Idle, dbConnectionMetricPrefix, name, hostname, isPrimary, p.poolID, "Idle")   //int

			wmetrics.Gauge(int(s.WaitCount), dbConnectionMetricPrefix, name, hostname, isPrimary, p.poolID, "WaitCount")                 //int64
			wmetrics.Gauge(int(s.WaitDuration), dbConnectionMetricPrefix, name, hostname, isPrimary, p.poolID, "WaitDuration")           // time.Duration
			wmetrics.Gauge(int(s.MaxIdleClosed), dbConnectionMetricPrefix, name, hostname, isPrimary, p.poolID, "MaxIdleClosed")         // int64
			wmetrics.Gauge(int(s.MaxLifetimeClosed), dbConnectionMetricPrefix, name, hostname, isPrimary, p.poolID, "MaxLifetimeClosed") // int64

		case <-p.stopMetrics:
			break loop
		}
	}

}
