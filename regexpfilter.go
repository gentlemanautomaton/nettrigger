package nettrigger

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type regexpFilter struct {
	subject string
	re      *regexp.Regexp
}

func (f regexpFilter) Filter(env Environment) bool {
	return f.re.MatchString(env.Expand(f.subject))
}

// RegexpBuilder constructs regular expression filters from filter
// specifications.
func RegexpBuilder(spec FilterSpec) (Filter, error) {
	switch strings.ToLower(spec.Type) {
	case "regexp", "regex", "re":
	default:
		return nil, nil
	}

	switch len(spec.Args) {
	case 0, 1:
		return nil, errors.New("regular expression filter requires a subject and an expression")
	case 2:
		subject, expression := spec.Args[0], spec.Args[1]

		// Force case-insensitive matching
		if !strings.HasPrefix(expression, "(?i)") {
			expression = "(?i)" + expression
		}

		re, err := regexp.Compile(expression)
		if err != nil {
			return nil, fmt.Errorf("invalid expression \"%s\": %v", expression, err)
		}

		return regexpFilter{
			subject: subject,
			re:      re,
		}.Filter, nil
	default:
		return nil, errors.New("regular expression filter has %d arguments when two are needed")
	}
}
