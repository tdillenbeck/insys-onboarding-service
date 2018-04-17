// +build go1.4

package appstats

import (
	"math"
	"runtime"
	"sort"
	"time"

	"weavelab.xyz/wlib/wmetrics/client"
)

type Uint64Slice []uint64

const (
	memoryPrefix = "appstats.mem"
)

func (s Uint64Slice) Len() int {
	return len(s)
}

func (s Uint64Slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Uint64Slice) Less(i, j int) bool {
	return s[i] < s[j]
}

func SendMemoryStats(client client.WMetrics, interval time.Duration) {
	go func() {
		var lastMemStats runtime.MemStats

		ticker := time.Tick(interval)

		for range ticker {
			var memStats runtime.MemStats
			runtime.ReadMemStats(&memStats)

			// sort the GC pause array
			length := len(memStats.PauseNs)
			if int(memStats.NumGC) < length {
				length = int(memStats.NumGC)
			}
			gcPauses := make(Uint64Slice, length)
			copy(gcPauses, memStats.PauseNs[:length])
			sort.Sort(gcPauses)

			client.Gauge(int(memStats.HeapObjects), memoryPrefix, "heap_objects")
			client.Gauge(int(memStats.HeapIdle), memoryPrefix, "heap_idle_bytes")
			client.Gauge(int(memStats.HeapInuse), memoryPrefix, "heap_in_use_bytes")
			client.Gauge(int(memStats.HeapReleased), memoryPrefix, "heap_released_bytes")
			client.Gauge(int(percentile(100.0, gcPauses, len(gcPauses))/1000), memoryPrefix, "gc_pause_usec_100")
			client.Gauge(int(percentile(99.0, gcPauses, len(gcPauses))/1000), memoryPrefix, "gc_pause_usec_99")
			client.Gauge(int(percentile(95.0, gcPauses, len(gcPauses))/1000), memoryPrefix, "gc_pause_usec_95")
			client.Gauge(int(memStats.NextGC), memoryPrefix, "next_gc_bytes")
			client.Incr(int(memStats.NumGC-lastMemStats.NumGC), memoryPrefix, "gc_runs")

			lastMemStats = memStats
		}
	}()
}

func percentile(perc float64, arr []uint64, length int) uint64 {
	if length == 0 {
		return 0
	}
	indexOfPerc := int(math.Floor(((perc / 100.0) * float64(length)) + 0.5))
	if indexOfPerc >= length {
		indexOfPerc = length - 1
	}
	return arr[indexOfPerc]
}
