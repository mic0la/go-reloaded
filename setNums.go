package reloaded

import (
	"fmt"
	"regexp"
	"strconv"
)

func SetNums(re *regexp.Regexp, str string, numTypeInt int) string {
	return re.ReplaceAllStringFunc(str, func(arr string) string {
		fmt.Println("catched", arr)
		connector := ""
		connectorEnd := ""
		switch arr[0] {
		case '\n':
			connector = "\n"
		case '\r':
			connector = "\r"
		case ' ':
			connector = " "
		}
		arr = theTrimSpace(arr)
		for i := 0; i < len(arr); i++ {
			switch arr[i] {
			case '(':
				arr = arr[:i]
			case ',':
				arr = arr[:i] + arr[i+1:]
				connectorEnd = ","
				i--
			case '.':
				arr = arr[:i] + arr[i+1:]
				connectorEnd = "."
				i--
			case '?':
				arr = arr[:i] + arr[i+1:]
				connectorEnd = "?"
				i--
			case '!':
				arr = arr[:i] + arr[i+1:]
				connectorEnd = "!"
				i--
			case ')':
				arr = arr[:i] + arr[i+1:]
				connectorEnd = ")"
				i--
			case ']':
				arr = arr[:i] + arr[i+1:]
				connectorEnd = "]"
				i--
			case '}':
				arr = arr[:i] + arr[i+1:]
				connectorEnd = "}"
				i--
			}
		}
		decDigit, _ := strconv.ParseInt(arr, numTypeInt, 64)
		result := strconv.Itoa(int(decDigit))
		return connector + result + connectorEnd
	})
}
