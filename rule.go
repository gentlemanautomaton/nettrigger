package nettrigger

import (
	"fmt"
)

// A Rule defines zero or more triggers and one or more actions.
type Rule struct {
	Filters []Filter
	Actions []Action
}

// Match returns true if the rule's filters match the environment.
func (r Rule) Match(env Environment) bool {
	for _, filter := range r.Filters {
		if !filter(env) {
			return false
		}
	}
	return true
}

// BuildRule uses fb and ab to construct a rule from the specification.
func BuildRule(spec RuleSpec, fb []FilterBuilder, ab []ActionBuilder) (Rule, error) {
	filters, err := BuildFilters(spec.Filters, fb...)
	if err != nil {
		return Rule{}, err
	}
	actions, err := BuildActions(spec.Actions, ab...)
	if err != nil {
		return Rule{}, err
	}
	return Rule{
		Filters: filters,
		Actions: actions,
	}, nil
}

// BuildRules converts the given specifications into a rule list.
func BuildRules(specs []RuleSpec, fb []FilterBuilder, ab []ActionBuilder) ([]Rule, error) {
	var rules []Rule
	for i, r := range specs {
		rule, err := BuildRule(r, fb, ab)
		if err != nil {
			return nil, fmt.Errorf("bad rule #%d: %v", i, err)
		}
		rules = append(rules, rule)
	}
	return rules, nil
}
