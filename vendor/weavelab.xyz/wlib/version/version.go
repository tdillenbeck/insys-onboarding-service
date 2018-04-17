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

	InstanceID           string
	Hostname             string
	FileModificationTime time.Time

	maintenance int
	minor       int
	major       int
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

// Version returns the major, minor, and maintenance version of the app
func Version() (int, int, int) {
	info := Info()
	return info.major, info.minor, info.maintenance
}

func setVersion() {
	l.Lock()
	_info.Version = fmt.Sprintf("%d.%d.%d", _info.major, _info.minor, _info.maintenance)
	l.Unlock()
}

// String returns a stringified version of the app
func String() string {
	info := Info()
	return info.Version
}

// InstanceID returns the instance id of the app
func InstanceID() string {
	info := Info()
	return info.InstanceID
}
