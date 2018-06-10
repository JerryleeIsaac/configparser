package basic

import (
	"bufio"
	"errors"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

// Errors for this package
var (
	ErrFileOpenFailed     = errors.New("failed to open config file")
	ErrInvalidFormat      = errors.New("config file has invalid format")
	ErrConfigDoesNotExist = errors.New("config does not exist")
	ErrInvalidIntConfig   = errors.New("config is not integer")
	ErrInvalidDateConfig  = errors.New("config is not duration")
)

// Duration regular expressions
var (
	MillisecondRegExp, _ = regexp.Compile("[0-9]+.ms")
	SecondRegExp, _      = regexp.Compile("[0-9]+.s")
	MinuteRegExp, _      = regexp.Compile("[0-9]+.m")
	HourRegExp, _        = regexp.Compile("[0-9]+.h")
	NumRegExp, _         = regexp.Compile("[0-9]+")
)

// Parser represents parser for a basic config file
type Parser struct {
	filename  string
	file      *os.File
	configMap map[string]string
}

// SetFile sets the file in which the configs are to be read.
func (p *Parser) SetFile(filename string) error {
	// Open file
	//
	var err error
	p.filename = filename
	p.file, err = os.Open(p.filename)
	if err != nil {
		log.Printf("Failed to open file with error: '%s'", err.Error())
		return err
	}
	return nil
}

// LoadConfig loads the config in the config file and stores it locally
// within the package.
func (p *Parser) LoadConfig() error {
	p.configMap = make(map[string]string)
	// Get bufio for reading file
	//
	scanner := bufio.NewScanner(p.file)

	for {
		// Check first character of line
		//
		if !scanner.Scan() {
			break
		}
		tok1 := scanner.Text()

		// Skip this line if it is a comment
		//
		if tok1 == "#" || tok1[0] == '#' {
			continue
		}

		// Get 2nd and 3rd token
		//
		if !scanner.Scan() {
			return ErrInvalidFormat
		}
		tok2 := scanner.Text()
		if !scanner.Scan() {
			return ErrInvalidFormat
		}
		tok3 := scanner.Text()

		// Check if token 2 follows the right format
		//
		if tok2 != "=" {
			return ErrInvalidFormat
		}

		p.configMap[tok1] = tok3
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

	// Convert config to integer
	//
	intVal, err := strconv.Atoi(val)
	if err != nil {
		return 0, ErrInvalidIntConfig
	}

	return intVal, nil
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
	// Get config from config map
	//
	val, exist := p.configMap[key]
	if !exist {
		return "", ErrConfigDoesNotExist
	}

	return val, nil
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

	// Get unit of duration
	//
	var unit time.Duration
	if MillisecondRegExp.MatchString(val) {
		unit = time.Millisecond
	} else if SecondRegExp.MatchString(val) {
		unit = time.Second
	} else if MinuteRegExp.MatchString(val) {
		unit = time.Minute
	} else if HourRegExp.MatchString(val) {
		unit = time.Hour
	} else {
		return 0, ErrInvalidDateConfig
	}

	// Get numeric coefficient
	//
	loc := NumRegExp.FindStringIndex(val)
	coeffStr := val[loc[0]:loc[1]] // TODO: get a proper name

	coeff, err := strconv.Atoi(coeffStr)
	if err != nil {
		return 0, ErrInvalidDateConfig
	}

	return time.Duration(coeff) * unit, nil
}
