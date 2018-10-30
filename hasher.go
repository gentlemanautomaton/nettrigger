package nettrigger

import (
	"crypto/sha256"
	"fmt"
)

// Hasher performs hashing of variable values taken from next.
func Hasher(s string, next Mapper) (value string, ok bool) {
	if inner, ok := ParseFunction("SHA2_256", s); ok && len(inner) > 0 {
		if value, ok := next(inner, next); ok {
			return fmt.Sprintf("%x", sha256.Sum256([]byte(value))), true
		}
		return "", false
	}
	if inner, ok := ParseFunction("SHA2_64", s); ok && len(inner) > 0 {
		if value, ok := next(inner, next); ok {
			hashed := sha256.Sum256([]byte(value))
			return fmt.Sprintf("%x", hashed[0:8]), true
		}
		return "", false
	}
	return "", false
}
