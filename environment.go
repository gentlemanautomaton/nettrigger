package nettrigger

// Environment provides the ability to expand environment variables.
type Environment interface {
	// Expand replaces ${var} or $var in the string according to the
	// state of the current environment.
	Expand(s string) string
}
