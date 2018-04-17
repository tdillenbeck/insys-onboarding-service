package config

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	"weavelab.xyz/wlib/wdns"
	"weavelab.xyz/wlib/werror"
)

// DefaultConfig is the main config object used by default when including this package
var DefaultConfig *Config

func init() {
	DefaultConfig = &Config{
		Settings: make(map[string]*Setting),
	}
}

// Values is a custom type which allows you to set a flag multiple times
type Values []string

// Set allows you to set any value in the Values type (currently strings only)
func (a *Values) Set(s string) error {
	*a = append(*a, s)
	return nil
}

// String returns a concatenated list of all the values in a config
func (a *Values) String() string {
	return strings.Join(*a, ";")
}

// Setting is the struct that contains all information for one single config setting
type Setting struct {
	Name         string
	DefaultValue string
	HelpMessage  string

	Env      string
	EnvValue string

	Flag      string
	FlagValue Values

	ShortFlag      string
	ShortFlagValue Values
}

func (s *Setting) init() {
	if s.Env != "" {
		s.EnvValue = os.Getenv(s.Env)
	}

	if s.Flag != "" {
		flag.Var(&s.FlagValue, s.Flag, s.HelpMessage)
	}

	if s.ShortFlag != "" {
		flag.Var(&s.ShortFlagValue, s.ShortFlag, s.HelpMessage)
	}
}

// Config is the struct that holds all the settings for the app
type Config struct {
	Settings map[string]*Setting
	mu       sync.Mutex
}

// New returns a new config
func New() *Config {
	return &Config{
		Settings: make(map[string]*Setting),
	}
}

// Add creates a new config setting for the app on the default config
func Add(name string, defaultValue string, message string, flags ...string) {
	DefaultConfig.Add(name, defaultValue, message, flags...)
}

// Add is for adding a config setting to a config struct
func (c *Config) Add(name string, defaultValue string, message string, flags ...string) {
	if len(flags) > 3 {
		panic("Too many arguments to Config.Add(name, default, [env, [flag, [shortFlag]]])")
	}

	if defaultValue != "" {
		message = fmt.Sprintf("%s | Default: %s", message, defaultValue)
	}

	setting := Setting{
		Name:         name,
		HelpMessage:  message,
		DefaultValue: defaultValue,
		Env:          strings.Replace(strings.ToUpper(name), "-", "_", -1),
		Flag:         name,
	}

	if len(flags) > 0 {
		setting.Env = flags[0]
	}
	if len(flags) > 1 {
		setting.Flag = flags[1]
	}
	if len(flags) > 2 {
		setting.ShortFlag = flags[2]
	}
	setting.init()

	c.mu.Lock()
	c.Settings[setting.Name] = &setting
	c.mu.Unlock()

	return
}

// Set modifies the value of an already-created config setting on the default config
func Set(name string, value string) {
	DefaultConfig.Set(name, value)
}

// Set modifies the value of an already-created config setting on instantiated config struct
func (c *Config) Set(name string, value string) {
	c.mu.Lock()
	if setting, ok := c.Settings[name]; ok {
		setting.EnvValue = value
		c.Settings[name] = setting
	}
	c.mu.Unlock()
}

// Get returns the value for the given setting name on the default config
func Get(name string) string {
	return DefaultConfig.Get(name)
}

// Get returns the value for a given setting name on the instantiated config struct
func (c *Config) Get(name string) string {

	c.mu.Lock()
	defer c.mu.Unlock()

	setting, ok := c.Settings[name]
	if ok != true || setting == nil {
		return ""
	}

	return setting.value()
}

// GetArray calls DefaultConfig.GetArray
func GetArray(name string) []string {
	return DefaultConfig.GetArray(name)
}

// GetArray returns all the values for a given setting
func (c *Config) GetArray(name string) []string {
	s := c.Get(name)
	if s == "" {
		return []string{}
	}

	s = strings.Trim(s, ";")
	return strings.Split(s, ";")
}

// GetInt calls DefaultConfig.GetInt
func GetInt(name string, die bool) (int, error) {
	return DefaultConfig.GetInt(name, die)
}

// GetInt returns the config setting for `name` with type `int`
func (c *Config) GetInt(name string, die bool) (int, error) {

	value := c.Get(name)
	result, err := strconv.Atoi(value)
	if err != nil {

		if die {
			exit(name, value, err)
		}

		// lookup the default value
		c.mu.Lock()
		setting, ok := c.Settings[name]
		c.mu.Unlock()

		if ok != true {
			return 0, fmt.Errorf("setting %s was never created so has no int value", name)
		}

		defaultValue, err := strconv.Atoi(setting.DefaultValue)
		if err != nil {
			return 0, fmt.Errorf("unable to parse %s as number, default value also invalid %s", value, setting.DefaultValue)
		}

		return defaultValue, fmt.Errorf("unable to parse %s as number", value)
	}

	return result, nil
}

// GetIntArray calls DefaultConfig equivalent
func GetIntArray(name string, die bool) ([]int, error) {
	return DefaultConfig.GetIntArray(name, die)
}

// GetIntArray returns a slice of ints for a given setting `name`
func (c *Config) GetIntArray(name string, die bool) ([]int, error) {
	var err error
	var x int
	stringArr := c.GetArray(name)
	intArr := make([]int, len(stringArr))
	for i, v := range stringArr {
		x, err = strconv.Atoi(v)
		if err != nil {
			if die {
				exit(name, fmt.Sprintf("[%v]", stringArr), err)
			}
			return nil, fmt.Errorf("unable to parse %s as number error=[%s]", v, err)
		}
		intArr[i] = x
	}
	return intArr, nil
}

// GetAddress calls DefaultConfig equivalent
func GetAddress(name string, die bool) (string, error) {
	return DefaultConfig.GetAddress(name, die)
}

// GetAddress looks up the value of the setting and converts it to a resolvable address.
// If the lookup fails GetAddress will either panic or attempt to use the default value, depending on the die bool.
// GetAddress uses wdns to do an SRV record lookup if necessary to resolve the port on which the service runs.
// For example, GetAddress can handle either "serviceB.serviceBNamespace:http" or "serviceB.serviceBNamespace:8080"
func (c *Config) GetAddress(name string, die bool) (string, error) {
	var err error
	value := c.Get(name)

	addr, err := resolveAddress(value)
	if err != nil {
		if die {
			panic(err)
		}
		return "", werror.Wrap(err, "Unable to resolve address for config value, will try using default instead.").Add("configValue", addr)
	}

	return addr, nil
}

// GetAddressArray calls DefaultConfig equivalent
func GetAddressArray(name string, die bool) ([]string, error) {
	return DefaultConfig.GetAddressArray(name, die)
}

func (c *Config) GetAddressArray(name string, die bool) ([]string, error) {
	values := c.GetArray(name)
	addrs := make([]string, 0, len(values))
	for _, v := range values {
		a, err := resolveAddress(v)
		if err != nil {
			if die {
				panic(err)
			}
			return nil, err
		}

		addrs = append(addrs, a)
	}

	return addrs, nil
}

func resolveAddress(addr string) (string, error) {

	a, err := wdns.ResolveAddress(addr)
	if err != nil {
		return "", err
	}

	return a, nil
}

// All calls DefaultConfig.All()
func All() map[string]string {
	return DefaultConfig.All()
}

// All returns a map of all the config settings for the given config object
func (c *Config) All() map[string]string {
	c.mu.Lock()
	defer c.mu.Unlock()

	all := make(map[string]string)
	for _, v := range c.Settings {
		all[v.Name] = v.value()
	}

	return all
}

// Parse calls DefaultConfig.parse()
func Parse() {
	DefaultConfig.Parse()
	DefaultConfig.PrettyPrint()
}

// Parse will call flag.Parse() if it hasn't been called yet
func (c *Config) Parse() {
	if flag.Parsed() == false {
		flag.Parse()
	}
}

func (s *Setting) value() string {
	switch {
	case s.EnvValue != "":
		return s.EnvValue
	case s.FlagValue.String() != "":
		return s.FlagValue.String()
	case s.ShortFlagValue.String() != "":
		return s.ShortFlagValue.String()
	default:
		return s.DefaultValue
	}
}

func exit(name string, value string, err error) {
	log.Println()
	os.Exit(1)
}
