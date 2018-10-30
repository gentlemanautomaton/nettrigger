package nettrigger

import "strings"

// Lower returns the lower case of variable values taken from next.
func Lower(s string, next Mapper) (value string, ok bool) {
	if inner, ok := ParseFunction("LOWER", s); ok && len(inner) > 0 {
		if value, ok := next(inner, next); ok {
			return strings.ToLower(value), true
		}
	}
	return "", false
}
