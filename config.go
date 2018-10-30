package nettrigger

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// Config holds configuration values.
type Config struct {
	Params            ArgMap
	Rules             []RuleSpec
	Timeout           time.Duration
	DigitalOceanToken string
}

// DefaultConfig holds the default configuration values.
var DefaultConfig = Config{}

// ParseEnv will parse environment variables and apply them to the
// configuration.
func (c *Config) ParseEnv() error {
	var (
		token, hasToken     = os.LookupEnv("DIGITAL_OCEAN_TOKEN")
		timeout, hasTimeout = os.LookupEnv("TIMEOUT")
	)

	if hasToken {
		c.DigitalOceanToken = token
	}

	if hasTimeout {
		t, err := time.ParseDuration(timeout)
		if err != nil {
			return fmt.Errorf("invalid timeout \"%s\": %v", timeout, err)
		}
		c.Timeout = t
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
