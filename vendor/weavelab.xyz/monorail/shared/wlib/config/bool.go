package config

import (
	"fmt"
	"strconv"
)

func GetBool(name string, die bool) (bool, error) {
	return DefaultConfig.GetBool(name, die)
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

// GetBoolArray calls DefaultConfig equivalent
func GetBoolArray(name string, die bool) ([]bool, error) {
	return DefaultConfig.GetBoolArray(name, die)
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
