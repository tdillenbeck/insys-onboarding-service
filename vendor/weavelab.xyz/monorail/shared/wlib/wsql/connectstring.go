package wsql

import (
	"net/url"

	"weavelab.xyz/monorail/shared/wlib/version"
)

type ConnectString struct {
	connectString string

	Driver   string // defaults to "postgres"
	Host     string
	Database string
	Username string
	Password string

	DefaultSchema string // common options
	SSLMode       string
	Role          string

	Params url.Values // allow arbitrary parameters
}

func (c *ConnectString) String() string {

	if c.Host == "" || c.Database == "" || c.Username == "" {
		return c.connectString
	}

	if c.Params == nil {
		c.Params = url.Values{}
	}

	if c.DefaultSchema != "" {
		c.Params.Set("search_path", c.DefaultSchema)
	}

	if c.Params.Get("application_name") == "" {
		applicationName := version.Info().Name
		if applicationName == "" {
			applicationName = "unknown"
		}
		c.Params.Set("application_name", applicationName)
	}

	if c.Params.Get("sslmode") == "" {
		if c.SSLMode == "" {
			c.SSLMode = "disable"
		}
		c.Params.Set("sslmode", c.SSLMode)
	}

	if c.Role != "" {
		c.Params.Set("role", c.Role)
	}

	q := c.Params.Encode()

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
