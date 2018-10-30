package nettrigger

// Literal returns the value of quoted literals.
func Literal(s string, next Mapper) (value string, ok bool) {
	if len(s) < 2 {
		return "", false
	}

	startQuote := s[0]
	endQuote := s[len(s)-1]
	if startQuote != endQuote {
		return "", false
	}

	switch startQuote {
	case '"', '\'', '`':
		return s[1 : len(s)-1], true
	}

	return "", false
}
