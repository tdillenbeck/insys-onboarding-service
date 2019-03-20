# Config Helper

This is to help services use the default way of configuration, environment variables, but also use flags if wanted. For now it only supports strings...

## Usage

* Define all config settings in init() function of your module
* Use the config.Add function to add a new config setting (See example below)
* Use config.Get(name) function to get settings with the following precedents:
  1. Command-line flags
  1. Environment variables
  1. Default values

The `Add()` function accepts up to 6 arguments. The first 3 are required:

1. name - the name of the config setting
1. defaultValue - the default value of the setting
1. helpMessage - the help message for the setting when using --help flag
1. environmentVariable - maps the setting to an environment variable (defaults to `name` in all uppercase if not provided)
1. flag - the command line flag to map to this setting (defaults to `name` if not provided)
1. shortFlag - if you want a flag parameter that is a short version... jason thinks this is overkill ;-)

```Go
package main

import (
    "flag"

    "lab.getweave.com/weave/go-weave-utilities/config"
)

// All config settings should be defined in init() function of your package
func init() {
    // config.Add(name, defaultValue, helpMessage, [environmentVariable, [flag, [shortFlag]]])
    config.Add("foo","bar","Set the foo Setting", "ENV_FOO") // Not using flag config
    config.Add("michael", "scott", "Which Michael?", "", "michael", "m") // Using flags
}

func main() {
    // Configs based on environment variables are available immediately
    config.Get("foo") // "bar"

    // Flags aren't available until you parse flags
    // Assuming this app started with flag -michael=vick
    config.Get("michael") // "scott"
    flag.Parse()
    config.Get("michael") // "vick"

}
```
