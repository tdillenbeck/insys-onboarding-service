package appstats

import (
	"os"
	"runtime"
	"time"

	"weavelab.xyz/monorail/shared/wlib/werror"
	"weavelab.xyz/monorail/shared/wlib/wmetrics/client"
	"weavelab.xyz/monorail/shared/wlib/wmetrics/wmetricslog"
)

const (
	appStatPrefix = "appstats"
)

func SendFileHandleStats(client client.WMetrics, interval time.Duration) {

	client.SetLabels(appStatPrefix, "metric")

	var path string
	switch runtime.GOOS {
	case "windows":
		wmetricslog.Logger.Debug("Not starting file handle stats because operating system does not support it.")
		return
	case "linux":
		path = "/proc/self/fd"
	case "darwin":
		path = "/dev/fd/"
	}

	go func() {
		ticker := time.Tick(interval)
		for range ticker {
			sendFileHandleStats(path, client)
		}
	}()
}

func sendFileHandleStats(path string, client client.WMetrics) {

	n := runtime.NumGoroutine()
	client.Gauge(n, appStatPrefix, "goroutines")

	d, err := os.Open(path)
	if err != nil {
		wmetricslog.Logger.WError(werror.Wrap(err))
		return
	}
	defer d.Close()

	files, err := d.Readdirnames(0)
	if err != nil {
		wmetricslog.Logger.WError(werror.Wrap(err))

		return
	}

	client.Gauge(len(files), appStatPrefix, "open_handles")

}
