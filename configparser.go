package configparser

import (
	"errors"
	"time"

	"gitlab.com/Jerrylee/configparser/basic"
	"gitlab.com/Jerrylee/configparser/json"
	"gitlab.com/Jerrylee/configparser/toml"
)

// ConfigType represents the type of file config
type ConfigType string

// Config types
var (
	JSONConfig  ConfigType = "json"
	BasicConfig ConfigType = "basic"
	TOMLConfig  ConfigType = "toml"
)

// ConfigParser is an interface for parsing config files
type ConfigParser interface {
	// SetFile sets the file in which the configs are to be read.
	SetFile(filename string) error

	// LoadConfig loads the config in the config file and stores it locally
	// within the package.
	LoadConfig() error

	// Int returns the int config for the key given.
	Int(key string) (int, error)

	// IntDefault returns the int config for the key given. Returns the default
	// value if the key was not found.
	IntDefault(key string, def int) int

	// String returns the string config for the key given.
	String(key string) (string, error)

	// StringDefault returns the string config for the key given. Returns default
	// value if key was not found.
	StringDefault(key string, def string) string

	// Duration returns the duration config for the key given.
	Duration(key string) (time.Duration, error)

	// Duration default returns the duration config for the key given. Returns
	// default value supplied if config was not found.
	DurationDefault(key string, def time.Duration) time.Duration
}

// NewConfigParser creates a config parser for a given config type
func NewConfigParser(configType ConfigType) (ConfigParser, error) {
	var configParser ConfigParser

	switch configType {
	case JSONConfig:
		configParser = new(json.Parser)
	case TOMLConfig:
		configParser = new(toml.Parser)
	case BasicConfig:
		configParser = new(basic.Parser)
	default:
		return nil, errors.New("config type is currently unsupported")
	}

	return configParser, nil
}
