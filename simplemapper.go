package nettrigger

// A SimpleMapper maps variables to values without recursion.
type SimpleMapper func(string) (string, bool)

// Mapper performs a simple mapping of s.
func (simple SimpleMapper) Mapper(s string, next Mapper) (string, bool) {
	return simple(s)
}
