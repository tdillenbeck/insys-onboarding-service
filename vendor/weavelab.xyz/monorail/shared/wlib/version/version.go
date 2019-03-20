package version

import (
	"fmt"
	"sync"
	"time"
)

type AppInfo struct {
	Version string

	Name      string
	GitHash   string
	GitBranch string

	Path      string
	GoVersion string

	StartTime time.Time

	Hostname             string
	FileModificationTime time.Time

	minor int
	major int
}

var (
	_info AppInfo
	l     sync.Mutex
)

func Info() AppInfo {

	l.Lock()
	info := _info
	l.Unlock()

	return info
}

// Register sets the major and minor version of the app
func Register(major, minor int) {
	l.Lock()
	_info.major = major
	_info.minor = minor
	l.Unlock()

	setVersion()

	return
}

// Version returns the major and minor version of the app
func Version() (int, int) {
	info := Info()
	return info.major, info.minor
}

func setVersion() {
	l.Lock()
	_info.Version = fmt.Sprintf("%d.%d.0", _info.major, _info.minor)
	l.Unlock()
}

// String returns a stringified version of the app
func String() string {
	info := Info()
	return info.Version
}
