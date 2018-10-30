package nettrigger

import "strings"

// Upper returns the upper case of variable values taken from next.
func Upper(s string, next Mapper) (value string, ok bool) {
	if inner, ok := ParseFunction("UPPER", s); ok && len(inner) > 0 {
		if value, ok := next(inner, next); ok {
			return strings.ToUpper(value), true
		}
	}
	return "", false
}
