package nettrigger

import (
	"strings"
)

// Concat concatenates variable values taken from next.
func Concat(s string, next Mapper) (value string, ok bool) {
	if inner, ok := ParseFunction("concat", s); ok && len(inner) > 0 {
		args := strings.Split(inner, "+")
		output := ""
		for _, arg := range args {
			if value, ok := next(arg, next); ok {
				output += value
			} else {
				return "", false
			}
		}
		return output, true
	}
	return "", false
}
