//nsqwapp helps an app listen to a single NSQ topic on a single NSQ channel
package nsqwapp

import (
	"strings"

	nsq "github.com/nsqio/go-nsq"
	"weavelab.xyz/wlib/color"
	"weavelab.xyz/wlib/wapp"
	"weavelab.xyz/wlib/werror"
	"weavelab.xyz/wlib/wlog"
	"weavelab.xyz/wlib/wmetrics"
	"weavelab.xyz/wlib/wnsq"
)

type Config struct {
	NSQConfig          *nsq.Config
	ConcurrentHandlers int
}

// NewConfig returns an NSQ configuration with Weave appropriate default values
func NewConfig() *Config {
	c := Config{
		NSQConfig:          nsq.NewConfig(),
		ConcurrentHandlers: 1,
	}

	c.NSQConfig.MaxInFlight = 15

	return &c
}

func bootstrapConsumer(topic string, channel string, cfg *Config, h wnsq.Handler) (*wnsq.Consumer, error) {
	if topic == "" {
		return nil, werror.New("nsq-topic config is required")
	}

	if channel == "" {
		return nil, werror.New("nsq-channel config is required")
	}

	if cfg == nil {
		wlog.Info("Using default NSQ configuration")
		cfg = NewConfig()
	}

	if cfg.ConcurrentHandlers <= 0 {
		wlog.Info("Using default concurrent handler value of 1")
		cfg.ConcurrentHandlers = 1
	}

	if cfg.NSQConfig == nil {
		c := NewConfig()
		cfg.NSQConfig = c.NSQConfig
		wlog.Info("No nsq.Config provided, using a default config instead.")
	}

	if cfg.NSQConfig.MaxInFlight < 10 {
		wlog.Info(maxInFlightErrorMessage)
	}

	q, err := wnsq.NewConsumer(topic, channel, cfg.NSQConfig)
	if err != nil {
		return nil, werror.Wrap(err, "error initializing new consumer")
	}
	q.AddConcurrentHandlers(h, cfg.ConcurrentHandlers)

	return q, nil
}

// Bootstrap takes an nsq-config, lookupd-addresses, and a wnsq.Handler and returns a Starter that can be passed to wapp.Up
func Bootstrap(topic string, channel string, lookupdAddrs []string, cfg *Config, h wnsq.Handler) wapp.StartFunc {

	return func() ([]wapp.StopFunc, error) {
		wmetrics.Incr(1, "wapp", "nsqwapp")

		q, err := bootstrapConsumer(topic, channel, cfg, h)
		if err != nil {
			return nil, err
		}

		err = q.ConnectToNSQLookupds(lookupdAddrs)
		if err != nil {
			return nil, werror.Wrap(err, "error connecting to NSQ Lookupds").Add("lookupds", strings.Join(lookupdAddrs, ", "))
		}

		return []wapp.StopFunc{wapp.WrapSimpleStopFunc(q.Stop)}, nil
	}
}

func BootstrapNSQD(topic string, channel string, nsqdAddress []string, cfg *Config, h wnsq.Handler) wapp.StartFunc {

	return func() ([]wapp.StopFunc, error) {
		wmetrics.Incr(1, "wapp", "nsqwapp")

		q, err := bootstrapConsumer(topic, channel, cfg, h)
		if err != nil {
			return nil, err
		}

		err = q.ConnectToNSQDs(nsqdAddress)
		if err != nil {
			return nil, werror.Wrap(err, "error connecting to nsqd addresses").Add("nsqdAddress", strings.Join(nsqdAddress, ", "))
		}

		return []wapp.StopFunc{wapp.WrapSimpleStopFunc(q.Stop)}, nil
	}
}

var maxInFlightErrorMessage = color.SprintFunc(color.FgHiRed)(`
***************************************************************************
	MAX IN FLIGHT VALUE IS TOO LOW - YOU MAY EXPERIENCE RANDOM SLOWNESS
***************************************************************************`)
