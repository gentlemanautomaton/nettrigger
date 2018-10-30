package nettrigger

import "strings"

// ArgMap maps variable names to argument indices.
type ArgMap map[string]int

// Value returns the named argument from the given argument list.
func (m ArgMap) Value(name string, args ...string) (value string, ok bool) {
	if index, ok := m[strings.ToLower(name)]; ok {
		if index < len(args) {
			return args[index], true
		}
	}
	return "", false
}

// Map returns a mapper that maps variable names in m to indices in args.
func (m ArgMap) Map(args ...string) Mapper {
	return func(s string, next Mapper) (value string, ok bool) {
		if m == nil {
			return "", false
		}
		return m.Value(s, args...)
	}
}
