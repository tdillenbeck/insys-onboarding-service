package config

import (
	"fmt"
	"time"
)

func GetDuration(name string, die bool) (time.Duration, error) {
	return DefaultConfig.GetDuration(name, die)
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

// GetDurationArray calls DefaultConfig equivalent
func GetDurationArray(name string, die bool) ([]time.Duration, error) {
	return DefaultConfig.GetDurationArray(name, die)
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
