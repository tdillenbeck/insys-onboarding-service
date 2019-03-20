package client

import (
	"fmt"
	"io"
	"net"
	"strings"
	"sync/atomic"
	"time"
	"unsafe"

	"weavelab.xyz/monorail/shared/wlib/config"
	"weavelab.xyz/monorail/shared/wlib/werror"
	"weavelab.xyz/monorail/shared/wlib/wlog"
	"weavelab.xyz/monorail/shared/wlib/wmetrics/aggregator"
	"weavelab.xyz/monorail/shared/wlib/wmetrics/wmetricsconfig"
	"weavelab.xyz/monorail/shared/wlib/wmetrics/wmetricslog"
)

type WMetrics interface {
	StartTimer(name string, s ...string) (stop func(...string))
	Time(delta time.Duration, name string, s ...string)
	Incr(count int, name string, s ...string)
	Decr(count int, name string, s ...string)
	Gauge(value int, name string, s ...string)
	SetLabels(metricName string, labelNames ...string)
}

const (
	//TODO are these suffixes right?
	countSfx = "count"
	timerSfx = "timer"
	gaugeSfx = "gauge"

	wmetricsdHostCfg = "wmetricshost"

	intervalUpdateDNS = time.Minute * 5
)

func init() {
	config.Add(wmetricsdHostCfg, "127.0.0.1", "The host at which wmetricsd is running. This should normally be running on localhost, but can be overridden to point at a remote wmetricsd or statsd instance.", "WMETRICSD_HOST")
}

//WMetricsClient handles adding the prefix to every stat sent through it
type WMetricsClient struct {
	prefix          string
	aggregateWriter io.WriteCloser

	*Labels
}

//NewDefaultWMetricsClient returns a client with a prefix for sending metrics.
func NewDefaultWMetricsClient(prefix ...string) (*WMetricsClient, error) {

	wmetricsdAddr := net.JoinHostPort(config.Get(wmetricsdHostCfg), config.Get(wmetricsconfig.WMetricsDPortCfg))

	packetSize, err := config.GetInt(wmetricsconfig.PacketSizeCfg, false)
	if err != nil {
		return nil, werror.Wrap(err, "packetSize config must be an integer")
	}

	c, err := NewAggregateWriterWMetricsClient(wmetricsdAddr, packetSize, time.Second, prefix...)
	if err != nil {
		return nil, werror.Wrap(err)
	}

	return c, nil
}

//NewAggregateWriterWMetricsClient can be used to create a WMetrics client with configurable values for the underlying aggregated writer
func NewAggregateWriterWMetricsClient(wmetricsdAddr string, packetSize int, sendInterval time.Duration, prefix ...string) (*WMetricsClient, error) {

	aggregateWriter, err := aggregateWriter(wmetricsdAddr, sendInterval, packetSize)
	if err != nil {
		return nil, werror.Wrap(err)
	}

	return New(aggregateWriter, prefix...), nil
}

func New(dst io.WriteCloser, prefix ...string) *WMetricsClient {
	prefixStr := ""
	if len(prefix) > 0 {
		prefix = clean(prefix...)
		//Join the prefix by . and add one on the end since it will be joined to the specific stat name as well
		prefixStr = strings.Join(prefix, ".") + "."
	}

	return &WMetricsClient{
		prefix:          prefixStr,
		aggregateWriter: dst,

		Labels: newLabels(),
	}
}

//aggregateWriter instantiates an aggregate.AggregateWriter
func aggregateWriter(addr string, sendInterval time.Duration, packetSize int) (io.WriteCloser, error) {

	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return nil, werror.Wrap(err)
	}

	go func() {
		t := time.Tick(intervalUpdateDNS)
		for {
			a, err := net.ResolveUDPAddr("udp", addr)
			if err != nil {
				wmetricslog.Logger.WError(werror.Wrap(err))
				return
			}

			dst := unsafe.Pointer(udpAddr)
			atomic.SwapPointer(&dst, unsafe.Pointer(a))
			<-t
		}
	}()

	connFunc := func() (io.WriteCloser, error) {

		conn, err := net.DialUDP("udp", nil, udpAddr)
		if err != nil {
			return nil, werror.Wrap(err, "unable to dial wmetricsd").Add("addr", addr)
		}

		return conn, nil
	}

	w := aggregator.NewAggregateWriter([]byte("\n"), sendInterval, packetSize, connFunc)

	return w, nil
}

//StartTimer is a convenience method for timing metrics. It returns a stop function that will stop and send the timing metric. Safe to call from multiple go-routines--stop will only send the metric for the first time it is called, otherwise it is a no-op.
func (w *WMetricsClient) StartTimer(name string, s ...string) (stop func(...string)) {
	start := time.Now()

	var stopped int32

	// stop func can label values as well, which will be appended to any that were passed to start
	// if stop inadvertently gets called multiple times we will only send metrics for it the first time it is called
	stop = func(ss ...string) {
		// check if we've already executed the stop function
		if !atomic.CompareAndSwapInt32(&stopped, 0, 1) {
			return
		}

		duration := time.Since(start)

		labels := append(s, ss...)
		w.Time(duration, name, labels...)
	}

	return stop
}

//Time can be used to directly send a Time metrics. Usually it's easier to use StartTimer
func (w *WMetricsClient) Time(delta time.Duration, name string, s ...string) {
	w.send("%d|ms", int(delta/time.Millisecond), name, s, timerSfx)
}

//Incr increments a statistic by the given count
func (w *WMetricsClient) Incr(count int, name string, s ...string) {
	w.send("%d|c", count, name, s, countSfx)
}

//Decr decrements a statistic by the given count
func (w *WMetricsClient) Decr(count int, name string, s ...string) {
	w.send("%d|c", -count, name, s, countSfx)
}

//Gauge sets a statistic to the given value
func (w *WMetricsClient) Gauge(value int, name string, s ...string) {
	w.send("%d|g", value, name, s, gaugeSfx)
}

func (w *WMetricsClient) send(format string, value int, name string, s []string, suffix string) {
	//If w is nil or no name has been given then simply return
	if w == nil || name == "" {
		return
	}

	s = verifySuffix(s, suffix)
	name = cleanStr(name)
	s = clean(s...)

	// the first item is the name and the last item is the suffix, so add labels to everything in between
	labelValues := s[:len(s)-1]

	labelNames, useLabelNames := w.Labels.labels(name)

	// only useLabelNames if len of labelValues matches len of labelNames, otherwise log an error
	if len(labelValues) != len(labelNames) {
		if useLabelNames {
			wlog.WError(werror.New("unexpected number of labels").Add("metric", name).Add("values", labelValues).Add("labels", labelNames))
		}
		useLabelNames = false
	}

	// add label names to label values
	for i, labelValue := range labelValues {
		labelName := string(i + 'a')
		if useLabelNames {
			labelName = labelNames[i]
		}

		if labelValue == "" {
			labelValue = "empty"
		}

		// replace corresponding string in stats name array
		s[i] = labelName + "=" + labelValue
	}

	//Format the stat
	stat := strings.Join(s, ".")

	format = fmt.Sprintf("%s%s.%s:%s|@1.00", w.prefix, name, stat, format)

	w.sendToWMetricsD(format, value)
}

func (w *WMetricsClient) sendToWMetricsD(format string, value int) {
	go func() {
		s := fmt.Sprintf(format, value)
		if _, err := w.aggregateWriter.Write([]byte(s)); err != nil {
			wmetricslog.Logger.WError(werror.Wrap(err, "Error sending to wmetricsd"))
			return
		}
	}()
}

func clean(strs ...string) []string {
	var cleaned []string

	for _, s := range strs {
		// split string on "." to treat each label separately
		split := strings.Split(s, ".")
		for _, ss := range split {
			if ss == "" {
				continue
			}
			cleaned = append(cleaned, cleanStr(ss))
		}
	}

	return cleaned
}

func cleanStr(s string) string {
	s = strings.Replace(s, "/", "_", -1)
	s = strings.Replace(s, " ", "_", -1)
	s = strings.Replace(s, ":", "_", -1)
	s = strings.Trim(s, ".")

	return s
}

//Check if the stat ends with the correct suffix and add it if not
func verifySuffix(s []string, suffix string) []string {
	if len(s) == 0 || s[len(s)-1] != suffix {
		return append(s, suffix)
	}
	return s
}
