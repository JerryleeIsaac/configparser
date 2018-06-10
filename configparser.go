package configparser

import (
	"time"
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