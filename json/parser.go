package json

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"time"
)

// Errors for this package
var (
	ErrFileReadFailed      = errors.New("failed to read config file")
	ErrInvalidFormat       = errors.New("config file has invalid format")
	ErrConfigDoesNotExist  = errors.New("config does not exist")
	ErrInvalidIntConfig    = errors.New("config is not integer")
	ErrInvalidStringConfig = errors.New("config is not string")
	ErrInvalidDateConfig   = errors.New("config is not duration")
)

// Duration regular expressions
var (
	MillisecondRegExp, _ = regexp.Compile("[0-9]+.ms")
	SecondRegExp, _      = regexp.Compile("[0-9]+.s")
	MinuteRegExp, _      = regexp.Compile("[0-9]+.m")
	HourRegExp, _        = regexp.Compile("[0-9]+.h")
	NumRegExp, _         = regexp.Compile("[0-9]+")
)

// Parser represents parser for a json config file
type Parser struct {
	filename  string
	configMap map[string]interface{}
}

// SetFile sets the file in which the configs are to be read.
func (p *Parser) SetFile(filename string) error {
	// Open file
	//
	p.filename = filename
	return nil
}

// LoadConfig loads the config in the config file and stores it locally
// within the package.
func (p *Parser) LoadConfig() error {
	// Read whole contents of file
	//
	raw, err := ioutil.ReadFile(p.filename)
	if err != nil {
		return ErrFileReadFailed
	}

	err = json.Unmarshal(raw, &p.configMap)
	if err != nil {
		log.Printf("Json unmarshal returned with error: '%s'.\n",
			err.Error())
		return ErrInvalidFormat
	}

	return nil
}

// Int returns the int config for the key given.
func (p Parser) Int(key string) (int, error) {
	return p.int(key)
}

// IntDefault returns the int config for the key given. Returns the default
// value if the key was not found.
func (p Parser) IntDefault(key string, def int) int {
	val, err := p.int(key)
	if err != nil {
		return def
	}

	return val
}

func (p Parser) int(key string) (int, error) {
	// Get config from config map
	//
	val, exist := p.configMap[key]
	if !exist {
		return 0, ErrConfigDoesNotExist
	}

	// Get numeric config
	//
	floatVal, check := val.(float64)
	if !check {
		return 0, ErrInvalidIntConfig
	}

	// Convert numeric config to int
	//
	return int(floatVal), nil

}

// String returns the string config for the key given.
func (p Parser) String(key string) (string, error) {
	return p.string(key)
}

// StringDefault returns the string config for the key given. Returns default
// value if key was not found.
func (p Parser) StringDefault(key string, def string) string {
	val, err := p.string(key)
	if err != nil {
		return def
	}

	return val
}

func (p Parser) string(key string) (string, error) {
	// Get config from map
	//
	val, exist := p.configMap[key]
	if !exist {
		return "", ErrConfigDoesNotExist
	}

	// Check if config is string
	//
	stringVal, check := val.(string)
	if !check {
		return "", ErrInvalidStringConfig
	}

	return stringVal, nil
}

// Duration returns the duration config for the key given.
func (p Parser) Duration(key string) (time.Duration, error) {
	return p.duration(key)
}

// DurationDefault returns the duration config for the key given. Returns
// default value supplied if config was not found.
func (p Parser) DurationDefault(key string, def time.Duration) time.Duration {
	val, err := p.duration(key)
	if err != nil {
		return def
	}

	return val
}

func (p Parser) duration(key string) (time.Duration, error) {
	// Get config from config map
	//
	val, exist := p.configMap[key]
	if !exist {
		return 0, ErrConfigDoesNotExist
	}

	durationStringVal, check := val.(string)
	if !check {
		return 0, ErrInvalidDateConfig
	}

	// Get unit of duration
	//
	var unit time.Duration
	if MillisecondRegExp.MatchString(durationStringVal) {
		unit = time.Millisecond
	} else if SecondRegExp.MatchString(durationStringVal) {
		unit = time.Second
	} else if MinuteRegExp.MatchString(durationStringVal) {
		unit = time.Minute
	} else if HourRegExp.MatchString(durationStringVal) {
		unit = time.Hour
	} else {
		return 0, ErrInvalidDateConfig
	}

	// Get numeric coefficient
	//
	loc := NumRegExp.FindStringIndex(durationStringVal)
	coeffStr := durationStringVal[loc[0]:loc[1]] // TODO: get a proper name

	coeff, err := strconv.Atoi(coeffStr)
	if err != nil {
		return 0, ErrInvalidDateConfig
	}

	return time.Duration(coeff) * unit, nil
}
