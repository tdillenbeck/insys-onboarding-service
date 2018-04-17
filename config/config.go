package config

import (
	"strings"
	"time"

	"weavelab.xyz/wlib/config"
	"weavelab.xyz/wlib/werror"
	"weavelab.xyz/wlib/wlog"
)

const (
	nsqLookupdAddrsConfig = "nsq-lookupd-addrs"
	nsqTopicConfig        = "nsq-listen-topic"
	nsqChannelConfig      = "nsq-listen-channel"

	maxInFlightConfig        = "max-in-flight"
	concurrentHandlersConfig = "concurrent-handlers"

	nsqdAddrConfig = "nsqd-addr"

	delayConfig = "delay" //just for example servers
)

var (
	NSQLookupdAddrs []string
	NSQTopic        string
	NSQChannel      string

	MaxInFlight        int
	ConcurrentHandlers int

	DelayConfig time.Duration
	NSQdAddr    string
)

func init() {
	config.Add(delayConfig, "1s", "amount of time to delay before responding")

	config.Add(nsqTopicConfig, "WebsocketManagerEvents", "The topic NSQ to consume")
	config.Add(nsqChannelConfig, "ClientWebsocketWatcher#ephemeral", "The channel on which to consume")
	config.Add(nsqLookupdAddrsConfig, "", "NSQ lookupd addresses")

	config.Add(maxInFlightConfig, "1000", "NSQ config number of times to attempt a message")
	config.Add(concurrentHandlersConfig, "100", "Number of concurrent handlers")

}

func Init() error {
	config.Parse()

	var err error

	NSQLookupdAddrs, err = config.GetAddressArray(nsqLookupdAddrsConfig, false)
	if err != nil {
		return werror.Wrap(err, "error getting lookupd-addrs")
	}

	NSQTopic = config.Get(nsqTopicConfig)

	NSQChannel = config.Get(nsqChannelConfig)
	if !strings.Contains(NSQChannel, "#ephemeral") {
		wlog.Info("Looks like you aren't using an ephemeral channel for listening... oh boy you're gutsy")
	}

	MaxInFlight, err = config.GetInt(maxInFlightConfig, false)
	if err != nil {
		return werror.Wrap(err, "error getting max attempts")
	}

	ConcurrentHandlers, err = config.GetInt(concurrentHandlersConfig, false)
	if err != nil {
		return werror.Wrap(err, "error getting concurrent handlers")
	}

	DelayConfig, err = config.GetDuration(delayConfig, false)
	if err != nil {
		return werror.Wrap(err, "error getting delay")
	}

	return nil
}
