package nettrigger

import "fmt"

// A Filter matches some particular condition of its environment.
type Filter func(Environment) bool

// BuildFilter uses builders to construct a filter from the specification.
func BuildFilter(spec FilterSpec, builders ...FilterBuilder) (Filter, error) {
	for _, builder := range builders {
		filter, err := builder(spec)
		if err != nil {
			return filter, err
		}
		if filter != nil {
			return filter, nil
		}
	}
	return nil, fmt.Errorf("unknown filter type \"%s\"", spec.Type)
}

// BuildFilters converts the given specifications into a filter list.
func BuildFilters(specs []FilterSpec, builders ...FilterBuilder) ([]Filter, error) {
	var filters []Filter
	for i, a := range specs {
		filter, err := BuildFilter(a, builders...)
		if err != nil {
			return nil, fmt.Errorf("invalid filter %d: %v", i, err)
		}
		filters = append(filters, filter)
	}
	return filters, nil
}
