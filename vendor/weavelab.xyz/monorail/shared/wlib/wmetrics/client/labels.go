package client

import "sync"

type Labels struct {
	labelNames map[string][]string

	labelsMtx sync.RWMutex
}

func newLabels() *Labels {
	return &Labels{
		labelNames: make(map[string][]string),
		labelsMtx:  sync.RWMutex{},
	}
}

func (l *Labels) SetLabels(metricName string, labelNames ...string) {
	l.labelsMtx.Lock()
	l.labelNames[metricName] = labelNames
	l.labelsMtx.Unlock()
}

func (l *Labels) labels(metricName string) ([]string, bool) {
	l.labelsMtx.RLock()
	labelNames, ok := l.labelNames[metricName]
	l.labelsMtx.RUnlock()

	return labelNames, ok
}
