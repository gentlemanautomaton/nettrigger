package nettrigger

import (
	"fmt"
	"strings"
)

// FilterSpec is a parsed filter specification.
type FilterSpec struct {
	Type string
	Args []string
}

// Arg returns the argument with the nth index.
//
// An empty string is returned if the argument doesn't exist.
func (spec FilterSpec) Arg(n int) string {
	if n >= len(spec.Args) {
		return ""
	}
	return spec.Args[n]
}

// ParseFilter parses v as a string representation of a filter.
func ParseFilter(v string) (FilterSpec, error) {
	parts := strings.Fields(v)
	switch len(parts) {
	case 0:
		return FilterSpec{}, fmt.Errorf("missing definition: \"%s\"", v)
	case 1:
		return FilterSpec{Type: parts[0]}, nil
	default:
		return FilterSpec{Type: parts[0], Args: parts[1:]}, nil
	}
}

// ParseFilters parses v as a string representation of a filter list.
func ParseFilters(v string) ([]FilterSpec, error) {
	if v == "" {
		return nil, nil
	}
	var filters []FilterSpec
	for i, f := range strings.Split(v, ",") {
		filter, err := ParseFilter(f)
		if err != nil {
			return nil, fmt.Errorf("invalid filter %d: %v", i, err)
		}
		filters = append(filters, filter)
	}
	return filters, nil
}
