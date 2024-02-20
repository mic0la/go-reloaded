package cadet

import (
	"regexp"
)

func Matched(toFind, str string) bool {
	matched, _ := regexp.MatchString(toFind, str)
	return matched
}

func CorrectAll(str string) string {
	if Matched("42", str) {
		return "suuuui"
	}
	return str
}
