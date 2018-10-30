package nettrigger

import (
	"fmt"
	"strings"
)

// ActionSpec is a parsed action specification.
type ActionSpec struct {
	Type string
	Args []string
}

// Arg returns the argument with the nth index.
//
// An empty string is returned if the argument doesn't exist.
func (spec ActionSpec) Arg(n int) string {
	if n >= len(spec.Args) {
		return ""
	}
	return spec.Args[n]
}

// ParseAction parses v as a string representation of an action.
func ParseAction(v string) (ActionSpec, error) {
	parts := strings.Fields(v)
	switch len(parts) {
	case 0:
		return ActionSpec{}, fmt.Errorf("missing definition: \"%s\"", v)
	case 1:
		return ActionSpec{Type: parts[0]}, nil
	default:
		return ActionSpec{Type: parts[0], Args: parts[1:]}, nil
	}
}

// ParseActions parses v as a string representation of an action list.
func ParseActions(v string) ([]ActionSpec, error) {
	if v == "" {
		return nil, nil
	}
	var actions []ActionSpec
	for i, a := range strings.Split(v, ",") {
		action, err := ParseAction(a)
		if err != nil {
			return nil, fmt.Errorf("invalid action %d: %v", i, err)
		}
		actions = append(actions, action)
	}
	return actions, nil
}
