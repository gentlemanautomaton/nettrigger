package nettrigger

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Config holds configuration values.
type Config struct {
	Params            ArgMap
	Rules             []RuleSpec
	Timeout           time.Duration
	GoogleProject     string
	DigitalOceanToken string
	Concurrent        bool
}

// DefaultConfig holds the default configuration values.
var DefaultConfig = Config{}

// ParseEnv will parse environment variables and apply them to the
// configuration.
func (c *Config) ParseEnv() error {
	var (
		googleProject, hasGoogleProject         = os.LookupEnv("GOOGLE_PROJECT")
		digitalOceanToken, hasDigitalOceanToken = os.LookupEnv("DIGITAL_OCEAN_TOKEN")
		timeout, hasTimeout                     = os.LookupEnv("TIMEOUT")
		concurrent, hasConcurrent               = os.LookupEnv("CONCURRENT")
	)

	if hasGoogleProject {
		c.GoogleProject = googleProject
	}

	if hasDigitalOceanToken {
		c.DigitalOceanToken = digitalOceanToken
	}

	if hasTimeout {
		t, err := time.ParseDuration(timeout)
		if err != nil {
			return fmt.Errorf("invalid timeout \"%s\": %v", timeout, err)
		}
		c.Timeout = t
	}

	if hasConcurrent {
		b, err := strconv.ParseBool(concurrent)
		if err != nil {
			return fmt.Errorf("invalid concurrent value \"%s\": %v", concurrent, err)
		}
		c.Concurrent = b
	}

	for i := 1; i < 1000; i++ {
		def, exists := os.LookupEnv(fmt.Sprintf("ARG%d", i))
		if !exists {
			break
		}
		if c.Params == nil {
			c.Params = make(ArgMap)
		}
		c.Params[strings.ToLower(def)] = i - 1
	}

	for i := 1; i < 1000; i++ {
		def, exists := os.LookupEnv(fmt.Sprintf("RULE%d", i))
		if !exists {
			break
		}
		rule, err := ParseRule(def)
		if err != nil {
			return fmt.Errorf("failed to parse rule %d: %v", i, err)
		}
		c.Rules = append(c.Rules, rule)
	}
	return nil
}
