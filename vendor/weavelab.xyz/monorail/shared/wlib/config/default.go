package config

import (
	"fmt"
	"time"
)

// DefaultConfig is the main config object used by default when including this package
var DefaultConfig *Config

func init() {
	DefaultConfig = New()
}

// All calls DefaultConfig.All()
func All() map[string]string {
	return DefaultConfig.All()
}

// Add creates a new config setting for the app on the default config
func Add(name string, defaultValue string, message string, flags ...string) {
	DefaultConfig.Add(name, defaultValue, message, flags...)
}

// Set modifies the value of an already-created config setting on the default config
func Set(name string, value string) {
	DefaultConfig.Set(name, value)
}

// Get returns the value for the given setting name on the default config
func Get(name string) string {
	return DefaultConfig.Get(name)
}

// GetArray calls DefaultConfig.GetArray
func GetArray(name string) []string {
	return DefaultConfig.GetArray(name)
}

// GetAddress calls DefaultConfig equivalent
func GetAddress(name string, die bool) (string, error) {
	return DefaultConfig.GetAddress(name, die)
}

// GetAddressArray calls DefaultConfig equivalent
func GetAddressArray(name string, die bool) ([]string, error) {
	return DefaultConfig.GetAddressArray(name, die)
}

func GetBool(name string, die bool) (bool, error) {
	return DefaultConfig.GetBool(name, die)
}

// GetBoolArray calls DefaultConfig equivalent
func GetBoolArray(name string, die bool) ([]bool, error) {
	return DefaultConfig.GetBoolArray(name, die)
}

// GetDuration returns the config setting for `name` with type `time.Duration`
func (c *Config) GetDuration(name string, die bool) (time.Duration, error) {
	value := c.Get(name)
	result, err := time.ParseDuration(value)
	if err != nil {
		if die {
			exit(name, value, err)
		}

		// lookup the default value
		c.mu.Lock()
		setting, ok := c.Settings[name]
		c.mu.Unlock()

		if ok != true {
			return 0, fmt.Errorf("setting %s was never created so has no duration value", name)
		}

		defaultValue, err := time.ParseDuration(setting.DefaultValue)
		if err != nil {
			return 0, fmt.Errorf("unable to parse %s as duration, default value also invalid %s", value, setting.DefaultValue)
		}

		return defaultValue, fmt.Errorf("unable to parse %s as duration", value)
	}

	return result, nil
}

// GetDurationArray returns a slice of durationss for a given setting `name`
func (c *Config) GetDurationArray(name string, die bool) ([]time.Duration, error) {

	stringArr := c.GetArray(name)
	durationArr := make([]time.Duration, len(stringArr))
	for i, v := range stringArr {
		x, err := time.ParseDuration(v)
		if err != nil {
			if die {
				exit(name, fmt.Sprintf("[%v]", stringArr), err)
			}

			return nil, fmt.Errorf("unable to parse %s as duration error=[%s]", v, err)
		}
		durationArr[i] = x
	}
	return durationArr, nil
}

// GetInt calls DefaultConfig.GetInt
func GetInt(name string, die bool) (int, error) {
	return DefaultConfig.GetInt(name, die)
}

// GetIntArray calls DefaultConfig equivalent
func GetIntArray(name string, die bool) ([]int, error) {
	return DefaultConfig.GetIntArray(name, die)
}

// Parse calls DefaultConfig.parse()
func Parse() {
	DefaultConfig.Parse()
	DefaultConfig.PrettyPrint()
}
