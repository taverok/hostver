package grep

import (
	"regexp"
	"strings"
)

func ConcatFirst(text string, r *regexp.Regexp, sep string) string {
	out := MatchFirst(text, r)

	return strings.Join(out, sep)
}

// MatchFirst returns the first matches of a regular expression in a string.
func MatchFirst(text string, r *regexp.Regexp) []string {
	out := r.FindAllStringSubmatch(text, -1)
	if len(out) == 0 || len(out[0]) < 2 {
		return nil
	}

	return out[0][1:]
}
