package nettrigger

// FilterBuilder builds filters from specifications.
type FilterBuilder func(FilterSpec) (Filter, error)
