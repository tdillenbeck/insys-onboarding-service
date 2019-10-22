package wsql

import (
	"net/url"
	"strconv"
	"strings"

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

	host string // internal use only
}

func (c *ConnectString) String() string {
	var cs string

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

	host, port := splitHostAndPort(c.Host)
	c.Params.Set("host", host)
	if port != "" {
		c.Params.Set("port", port)
	}

	q := c.Params.Encode()

	up := url.UserPassword(c.Username, c.Password)

	u := url.URL{
		Scheme:   "postgres",
		Host:     "",
		User:     up,
		Path:     "/" + c.Database,
		RawQuery: q,
	}

	cs = u.String()

	return cs
}

func splitHostAndPort(hostAndPort string) (string, string) {
	portColonIdx := strings.LastIndex(hostAndPort, ":")

	port := hostAndPort[portColonIdx+1:]
	_, err := strconv.Atoi(port)
	if err != nil {
		port = ""
		portColonIdx = len(hostAndPort)
	}
	host := hostAndPort[:portColonIdx]

	return host, port
}

// SetConnectString allows you to directly set the connection string that is used,
// and bypasses templating
func (c *ConnectString) SetConnectString(s string) {
	c.connectString = s
}

func (c *ConnectString) Hostname() string {
	if c.host != "" {
		return c.host
	}
	c.host = hostnameFromConnectionString(c.String())

	return c.host
}

func hostnameFromConnectionString(cs string) string {
	u, err := url.Parse(cs)
	if err != nil {
		return ""
	}

	if u.Hostname() != "" {
		return u.Hostname()
	}

	return u.Query().Get("host")
}
