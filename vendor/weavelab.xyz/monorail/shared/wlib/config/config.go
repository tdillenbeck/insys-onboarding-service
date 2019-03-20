package config

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"weavelab.xyz/monorail/shared/wlib/wdns"
	"weavelab.xyz/monorail/shared/wlib/werror"
)

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

func (s *Setting) init(flagSet *flag.FlagSet) {
	if s.Env != "" {
		s.EnvValue = os.Getenv(s.Env)
	}

	if s.Flag != "" {
		flagSet.Var(&s.FlagValue, s.Flag, s.HelpMessage)
	}

	if s.ShortFlag != "" {
		flagSet.Var(&s.ShortFlagValue, s.ShortFlag, s.HelpMessage)
	}
}

// Config is the struct that holds all the settings for the app
type Config struct {
	Settings map[string]*Setting
	mu       sync.Mutex
	flagSet  *flag.FlagSet
}

// New returns a new config
func New() *Config {
	return &Config{
		Settings: make(map[string]*Setting),
		flagSet:  flag.NewFlagSet(os.Args[0], flag.ExitOnError),
	}
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
	setting.init(c.flagSet)

	c.mu.Lock()
	c.Settings[setting.Name] = &setting
	c.mu.Unlock()

	return
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

// GetArray returns all the values for a given setting
func (c *Config) GetArray(name string) []string {
	s := c.Get(name)
	if s == "" {
		return []string{}
	}

	s = strings.Trim(s, ";")
	return strings.Split(s, ";")
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

// GetBool returns the config setting for `name` with type `bool`
func (c *Config) GetBool(name string, die bool) (bool, error) {

	value := c.Get(name)
	result, err := strconv.ParseBool(value)
	if err != nil {
		if die {
			exit(name, value, err)
		}

		// lookup the default value
		c.mu.Lock()
		setting, ok := c.Settings[name]
		c.mu.Unlock()

		if ok != true {
			return false, fmt.Errorf("setting %s was never created so has no bool value", name)
		}

		defaultValue, err := strconv.ParseBool(setting.DefaultValue)
		if err != nil {
			return false, fmt.Errorf("unable to parse %s as bool, default value also invalid %s", value, setting.DefaultValue)
		}

		return defaultValue, fmt.Errorf("unable to parse %s as bool", value)
	}

	return result, nil
}

// GetBoolArray returns a slice of bools for a given setting `name`
func (c *Config) GetBoolArray(name string, die bool) ([]bool, error) {
	var err error
	var x bool
	stringArr := c.GetArray(name)
	boolArr := make([]bool, len(stringArr))
	for i, v := range stringArr {
		x, err = strconv.ParseBool(v)
		if err != nil {
			if die {
				exit(name, fmt.Sprintf("[%v]", stringArr), err)
			}

			return nil, fmt.Errorf("unable to parse %s as bool error=[%s]", v, err)
		}
		boolArr[i] = x
	}
	return boolArr, nil
}

func GetDuration(name string, die bool) (time.Duration, error) {
	return DefaultConfig.GetDuration(name, die)
}

// GetDurationArray calls DefaultConfig equivalent
func GetDurationArray(name string, die bool) ([]time.Duration, error) {
	return DefaultConfig.GetDurationArray(name, die)
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

func resolveAddress(addr string) (string, error) {

	a, err := wdns.ResolveAddress(addr)
	if err != nil {
		return "", err
	}

	return a, nil
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

// Parse will call flag.Parse() if it hasn't been called yet
func (c *Config) Parse() {
	if c.flagSet.Parsed() == false {
		c.flagSet.Parse(os.Args[1:])
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
