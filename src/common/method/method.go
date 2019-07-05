package method

import (
	"regexp"
)

func RegexFind(pat, str string) []string {
	var patRegex = regexp.MustCompile(pat)
	strs := patRegex.FindStringSubmatch(str)
	return strs
}
