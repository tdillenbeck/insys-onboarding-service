package wsql

import (
	"net/url"

	"weavelab.xyz/wlib/version"
)

type ConnectString struct {
	connectString string

	Driver   string // defaults to "postgres"
	Host     string
	Database string
	Username string
	Password string
}

func (c *ConnectString) String() string {

	if c.Host == "" || c.Database == "" || c.Username == "" {
		return c.connectString
	}

	applicationName := version.Info().Name
	if applicationName == "" {
		applicationName = "unknown"
	}

	v := url.Values{}
	v.Add("sslmode", "disable")
	v.Add("application_name", applicationName)

	q := v.Encode()

	up := url.UserPassword(c.Username, c.Password)

	u := url.URL{
		Scheme:   "postgres",
		User:     up,
		Host:     c.Host,
		Path:     c.Database,
		RawQuery: q,
	}

	cs := u.String()

	return cs
}

// SetConnectString allows you to directly set the connection string that is used,
// and bypasses templating
func (c *ConnectString) SetConnectString(s string) {
	c.connectString = s
}
