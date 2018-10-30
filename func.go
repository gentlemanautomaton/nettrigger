package nettrigger

import (
	"strings"
)

// ParseFunction attemps to parse s as a function. If successful it
// returns the function's parameters and true.
func ParseFunction(name, s string) (params string, ok bool) {
	if len(s) < len(name)+2 {
		return "", false
	}
	if !strings.EqualFold(name, s[0:len(name)]) {
		return "", false
	}
	if s[len(name)] != '(' {
		return "", false
	}
	if s[len(s)-1] != ')' {
		return "", false
	}
	return s[len(name)+1 : len(s)-1], true
}
