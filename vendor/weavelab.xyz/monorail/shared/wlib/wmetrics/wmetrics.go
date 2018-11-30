package wmetrics

import (
	"time"

	"weavelab.xyz/monorail/shared/wlib/werror"
	"weavelab.xyz/monorail/shared/wlib/wmetrics/appstats"
	"weavelab.xyz/monorail/shared/wlib/wmetrics/client"
	"weavelab.xyz/monorail/shared/wlib/wmetrics/wmetricslog"
)

func init() {
	var err error
	DefaultClient, err = client.NewDefaultWMetricsClient(DefaultPrefix())
	if err != nil {
		wmetricslog.Logger.WError(werror.Wrap(err, "unable to setup default wmetrics client"))
	}

	wmetricslog.Logger.Debug("Starting FileHandle and Memory metrics tracking")
	appstats.SendStartup(DefaultClient)
	appstats.SendFileHandleStats(DefaultClient, 30*time.Second)
	appstats.SendMemoryStats(DefaultClient, 1*time.Minute)
}

// DefaultClient is where all global wmetrics.Incr etc. forwards to
var DefaultClient client.WMetrics

//StartTimer is a convenience method for timing metrics.
//It returns a stop function that will stop and send the timing metric using the default client.
func StartTimer(name string, s ...string) (stop func(s ...string)) {
	if DefaultClient == nil {
		return func(s ...string) {}
	}
	return DefaultClient.StartTimer(name, s...)
}

//Time can be used to directly send a Time metrics using the default client. Usually it's easier to use StartTimer
func Time(delta time.Duration, name string, s ...string) {
	if DefaultClient == nil {
		return
	}
	DefaultClient.Time(delta, name, s...)
}

//Incr increments a statistic by the given count using the default client
func Incr(count int, name string, s ...string) {
	if DefaultClient == nil {
		return
	}
	DefaultClient.Incr(count, name, s...)
}

//Decr decrements a statistic by the given count using the default client
func Decr(count int, name string, s ...string) {
	if DefaultClient == nil {
		return
	}
	DefaultClient.Decr(count, name, s...)
}

//Gauge sets a statistic to the given value using the default client
func Gauge(value int, name string, s ...string) {
	if DefaultClient == nil {
		return
	}
	DefaultClient.Gauge(value, name, s...)
}

func SetLabels(metricName string, labelNames ...string) {
	if DefaultClient == nil {
		return
	}
	DefaultClient.SetLabels(metricName, labelNames...)
}
