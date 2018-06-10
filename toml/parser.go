package toml

import "time"

// Parser represents parser for a toml config file
type Parser struct {
}

// SetFile sets the file in which the configs are to be read.
func (p *Parser) SetFile(filename string) error {
	return nil
}

// LoadConfig loads the config in the config file and stores it locally
// within the package.
func (p *Parser) LoadConfig() error {
	return nil
}

// Int returns the int config for the key given.
func (p *Parser) Int(key string) (int, error) {
	return 0, nil
}

// IntDefault returns the int config for the key given. Returns the default
// value if the key was not found.
func (p *Parser) IntDefault(key string, def int) int {
	return 0
}

// String returns the string config for the key given.
func (p *Parser) String(key string) (string, error) {
	return "", nil
}

// StringDefault returns the string config for the key given. Returns default
// value if key was not found.
func (p *Parser) StringDefault(key string, def string) string {
	return ""
}

// Duration returns the duration config for the key given.
func (p *Parser) Duration(key string) (time.Duration, error) {
	return 0, nil
}

// Duration default returns the duration config for the key given. Returns
// default value supplied if config was not found.
func (p *Parser) DurationDefault(key string, def time.Duration) time.Duration {
	return 0
}
