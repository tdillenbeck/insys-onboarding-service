#Sample Metrics Usage

```go
package main

import (
	"math/rand"
	"time"

	"github.com/weave-lab/wlib/wlog"
	"github.com/weave-lab/wlib/wmetrics"
	"github.com/weave-lab/wlib/wmetrics/client"
	"github.com/weave-lab/wlib/wmetrics/wmetricslog"
)

//A custom metrics client with a different prefix
var specialTimingClient = client.NewDefaultWMetricsClient(wmetrics.DefaultPrefix(), "special", "timing")

func main() {
	//Replace wmetrics' logger with a custom logger
	wmetricslog.Logger = wlog.NewWLogger(wlog.WlogdLogger)
	wmetricslog.Logger.SetDebugLogging(true)

	var count int64
	for {
		wmetrics.Incr(1, "index", "loop_started")

		work()

		wmetrics.Decr(1, "index", "loop_completed")

		wmetrics.Gauge(count, "index")
		count++
	}
}

func work() {
	//Starts timer now and ends it when work is completed
	defer specialTimingClient.StartTimer("work")()

	wait := rand.Intn(10)
	<-time.After(time.Duration(wait) * time.Second)
}

```