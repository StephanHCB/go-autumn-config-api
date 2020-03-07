package auconfigapi

// Type for a config item validation function.
//
// Important: note that you are given the KEY of the config item. This avoids having to use interface{} here,
// so you won't have to do type casting. Instead, just look up the value with
// viper.GetString(key), or viper.GetUint(key), or whatever you need.
type ConfigValidationFunc  func(key string) error

// Empty validation function. Use this if you don't need validation for a key.
func ConfigNeedsNoValidation(_ string) error {
	return nil
}

// Represents a configuration item for go-autumn-config.
//
// When you call auconfig.Setup(...) with a list of these, it will configure a command line flag and
// an environment variable.
//
// When you request configuration to be loaded, which you must do
// yourself with a call to auconfig.Load(), every key is assigned its value by going through
// the following precedence list:
//
//  - command line flag
//  - environment variable
//  - configuration read from secrets.(yaml|json|properties)
//  - configuration read from the config-[profile].(yaml|json|properties|...) in reverse order
//  - configuration read from config.(yaml|json|properties)
//  - default value specified in ConfigItems
//
type ConfigItem struct {
	// Where in the configuration structure the config item resides.
	//
	// Hierarchical levels are separated by '.', key components should match [a-z][a-z0-9-]*
	// that is, stick to lowercase letters and possibly - sign.
	//
	// Examples: "server.host",
	//           "server.port",
	//           "profiles"
	Key         string

	// Default value that also specifies the type of the value.
	//
	// You must always specify a value that has a type, or else detection of types will not work.
	// That is, "" is ok, nil is not. Numeric types need a type specifier such as int32(-10).
	//
	// Examples: "localhost",
	//           uint(8080),
	//           []string{}  (for a list of strings)
	Default     interface{}

	// A human readable description.
	//
	// Examples: "The hostname or ip address that determines the interface to listen on, defaults to localhost",
	//           "The port your main web controller should listen on. Defaults to 8080",
	//           "The list of profiles to load",
	Description string

	// Override name for environment variable (optional)
	//
	// If left blank, defaults to CONFIG_ + uppercase key with all non [a-z0-9] replaced with _.
	EnvName     string

	// Override name for the command line flag (optional)
	//
	// If left blank, defaults to the key.
	FlagName    string

	// Validation function that should return an error if the validation failed
	//
	// You MUST provide one, but you can just use ConfigNeedsNoValidation
	Validate    ConfigValidationFunc
}

// Type for a fatal error handler during initial configuration load. Expected to halt program execution.
//
// Example: func fail(err error) { panic(err) }
//
// ... but you probably want to use a function from the logging package of your choice.
type ConfigFailFunc func(err error)

// Type for a warning message logging handler during configuration load and validation. Should not halt execution.
//
// Example: func warn(message string) { log.Printf(message) }
//
// ... but you probably want to use a function from the logging package of your choice.
type ConfigWarnFunc func(message string)
