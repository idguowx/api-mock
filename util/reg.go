package util

import "regexp"

func RegMatch(regStr string, matchStr string) bool {
	r, _ := regexp.Compile(regStr)
	return r.MatchString(matchStr)
}
