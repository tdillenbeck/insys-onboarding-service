package client

import "time"

type nullClient struct{}

//Time can be used to directly send a Time metrics. Usually it's easier to use StartTimer
func (nullClient) Time(delta time.Duration, name string, s ...string) {}

//Incr increments a statistic by the given count
func (nullClient) Incr(count int, name string, s ...string) {}

//Decr decrements a statistic by the given count
func (nullClient) Decr(count int, name string, s ...string) {}

//Gauge sets a statistic to the given value
func (nullClient) Gauge(value int, name string, s ...string) {}

func (nullClient) StartTimer(name string, s ...string) (stop func(...string)) {
	return func(_ ...string) {}
}

func (nullClient) SetLabels(name string, labels ...string) {}
