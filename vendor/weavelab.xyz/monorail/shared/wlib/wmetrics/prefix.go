package wmetrics

import (
	"os"
	"path/filepath"
	"strings"
)

var defaultPrefix string

//DefaultPrefix describes the service and location
func DefaultPrefix() string {
	if defaultPrefix == "" {
		hostname, _ := os.Hostname()
		hostname = strings.Replace(hostname, ".", "_", -1)

		exeName := filepath.Base(os.Args[0])
		exeName = strings.Replace(exeName, ".", "_", -1)

		defaultPrefix = exeName + "." + hostname
	}

	return defaultPrefix
}
