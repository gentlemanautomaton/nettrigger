package nettrigger

import (
	"errors"
	"fmt"

	"github.com/gobwas/glob"
)

type patternFilter struct {
	subject string
	pattern glob.Glob
}

func (f patternFilter) Filter(env Environment) bool {
	return f.pattern.Match(env.Expand(f.subject))
}

// PatternBuilder constructs pattern filters from filter specifications.
func PatternBuilder(spec FilterSpec) (Filter, error) {
	switch spec.Type {
	case "pattern", "pat":
	default:
		return nil, nil
	}

	switch len(spec.Args) {
	case 0, 1:
		return nil, errors.New("pattern filter requires a subject and a pattern")
	case 2:
		subject, pattern := spec.Args[0], spec.Args[1]
		g, err := glob.Compile(pattern)
		if err != nil {
			return nil, fmt.Errorf("invalid pattern \"%s\": %v", pattern, err)
		}
		return patternFilter{
			subject: subject,
			pattern: g,
		}.Filter, nil
	default:
		return nil, errors.New("pattern filter has %d arguments when two are needed")
	}
}
