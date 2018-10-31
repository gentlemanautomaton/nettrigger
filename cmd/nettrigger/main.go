package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"syscall"

	"github.com/gentlemanautomaton/nettrigger"
	"github.com/gentlemanautomaton/signaler"
)

func main() {
	// Prepare config from environment
	c := nettrigger.DefaultConfig
	if err := c.ParseEnv(); err != nil {
		fmt.Printf("Configuration Error: %v\n", err)
		return
	}

	// Parse flags
	var (
		debug bool
	)
	flag.BoolVar(&debug, "debug", false, "print debug messages")
	flag.Parse()

	// Parse args
	args := flag.Args()

	// Possible filters
	filterBuilders := []nettrigger.FilterBuilder{
		nettrigger.PatternBuilder,
		nettrigger.RegexpBuilder,
	}

	// Possible actions
	actionBuilders := []nettrigger.ActionBuilder{
		nettrigger.DomainRecordActionBuilder,
	}

	// Environment
	env := nettrigger.Environment(nettrigger.MapperSet{
		nettrigger.Concat,
		nettrigger.Hasher,
		nettrigger.Upper,
		nettrigger.Lower,
		c.Params.Map(args...),
		nettrigger.Literal,
		nettrigger.SimpleMapper(os.LookupEnv).Mapper,
	})

	// Providers
	prov := nettrigger.Providers{}
	switch {
	case c.GoogleProject != "":
		var err error
		prov.DNS, err = nettrigger.NewGoogleDNS(c.GoogleProject)
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
			return
		}
	case c.DigitalOceanToken != "":
		prov.DNS = nettrigger.NewDigitalOceanDNS(c.DigitalOceanToken)
	}

	// Debug messages
	if debug {
		fmt.Printf("ARGS:\n")
		var lines []string
		for name := range c.Params {
			if value, ok := c.Params.Value(name, args...); ok {
				lines = append(lines, fmt.Sprintf("  %s: %s", name, value))
			} else {
				lines = append(lines, fmt.Sprintf("  %s: undefined", name))
			}
		}
		sort.Strings(lines)
		for _, line := range lines {
			fmt.Println(line)
		}
		fmt.Printf("RULES:\n")
		for r := range c.Rules {
			fmt.Printf("  [%d]: %v\n", r, c.Rules[r])
		}
	}

	// Build rules
	rules, err := nettrigger.BuildRules(c.Rules, filterBuilders, actionBuilders)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}

	// Capture shutdown signals
	shutdown := signaler.New().Capture(os.Interrupt, syscall.SIGTERM)

	// Handle timeouts
	if c.Timeout != 0 {
		shutdown.Timeout(c.Timeout)
	}

	// Prepare a context
	ctx := shutdown.Context()

	// Apply matching rules
	for r, rule := range rules {
		if rule.Match(env) {
			for a, action := range rule.Actions {
				err := action(ctx, env, prov)
				if err != nil {
					fmt.Printf("RULE %d ACTION %d: %v\n", r+1, a+1, err)
				}
			}
		}
	}
}
