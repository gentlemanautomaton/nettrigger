package nettrigger

import "os"

// MapperSet maps variables to values by iterating through a set of mappers.
type MapperSet []Mapper

// Expand replaces ${var} or $var in the string based on the mapper set.
func (ms MapperSet) Expand(s string) string {
	return os.Expand(s, ms.Resolve)
}

// value returns the value for the variable s.
func (ms MapperSet) value(s string, next Mapper) (value string, ok bool) {
	for _, mapper := range ms {
		if value, ok = mapper(s, next); ok {
			return
		}
	}
	return
}

// Resolve returns the value for the variable s. It returns an empty string
// if the s can't be resolved.
func (ms MapperSet) Resolve(s string) string {
	value, _ := ms.value(s, ms.value)
	return value
}
