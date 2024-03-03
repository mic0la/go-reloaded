package reloaded

import (
	"regexp"
	"strconv"
)

func SetNums(re *regexp.Regexp, str string, numTypeInt int) string {
	return re.ReplaceAllStringFunc(str, func(arr string) string {
		connector := ""
		switch arr[0] {
		case '\n':
			connector = "\n"
		case '\r':
			connector = "\r"
		case ' ':
			connector = " "
		}
		arr = theTrimSpace(arr)
		for i, v := range arr {
			if v == '(' {
				arr = arr[:i]
			}
		}
		decDigit, _ := strconv.ParseInt(arr, numTypeInt, 64)
		result := strconv.Itoa(int(decDigit))
		return connector + result
	})
}
