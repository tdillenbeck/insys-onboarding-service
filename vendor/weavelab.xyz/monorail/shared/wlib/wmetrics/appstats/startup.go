package appstats

import "weavelab.xyz/monorail/shared/wlib/wmetrics/client"

//SendStartup increments a counter with the app version on every startup.
func SendStartup(client client.WMetrics) {
	client.Incr(1, "appstats.start")
}
