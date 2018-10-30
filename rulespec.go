package nettrigger

import (
	"fmt"
	"strings"
)

// RuleSpec is a parsed rule specification.
type RuleSpec struct {
	Filters []FilterSpec
	Actions []ActionSpec
}

// ParseRule parses v as a string representation of a rule.
func ParseRule(v string) (RuleSpec, error) {
	var filters, actions string
	if parts := strings.SplitN(v, ":", 2); len(parts) == 2 {
		filters, actions = parts[0], parts[1]
	} else {
		actions = v
	}

	f, err := ParseFilters(filters)
	if err != nil {
		return RuleSpec{}, fmt.Errorf("invalid filter list: %v", err)
	}

	a, err := ParseActions(actions)
	if err != nil {
		return RuleSpec{}, fmt.Errorf("invalid action list: %v", err)
	}

	return RuleSpec{
		Filters: f,
		Actions: a,
	}, nil
}

// ParseRules parses v as a string representation of a rule list.
func ParseRules(v string) ([]RuleSpec, error) {
	var rules []RuleSpec
	for i, r := range strings.Split(v, "|") {
		rule, err := ParseRule(r)
		if err != nil {
			return nil, fmt.Errorf("invalid rule %d: %v", i, err)
		}
		rules = append(rules, rule)
	}
	return rules, nil
}
