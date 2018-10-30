package nettrigger

import "os"

// Mapper converts variables and functions into values.
type Mapper func(s string, next Mapper) (value string, ok bool)

// Expand replaces ${var} or $var in the string with its value.
func (m Mapper) Expand(s string) string {
	return os.Expand(s, func(v string) string {
		value, _ := m(v, nil)
		return value
	})
}
