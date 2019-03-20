package dev

import (
	"os"
	"strings"
)

var isDev bool

func init() {
	hostname, _ := os.Hostname()

	// for some tests, we want to test production settings
	switch {
	case os.Getenv("DEV") == "TRUE":
		isDev = true
	case os.Getenv("DEV") == "FALSE":
		isDev = false
	case strings.HasPrefix(hostname, "dev") || strings.HasPrefix(hostname, "local"):
		isDev = true
	}

}

func IsDev() bool {
	return isDev
}

func SetIsDev(b bool) {
	isDev = b
}
