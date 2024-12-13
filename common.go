package directive

import (
	"regexp"
)

func isMatch(pattern, name string) bool {
	return regexp.MustCompile(pattern).MatchString(name)
}
