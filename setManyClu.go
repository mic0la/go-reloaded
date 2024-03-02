package reloaded

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func SetCharsMany(re *regexp.Regexp, str string, charType string) string {
	return re.ReplaceAllStringFunc(str, func(arr string) string {
		wordsToChange := 0
		startFrom := 0
		countSpace := 0
		var arrToChange string
		var i int
		var cutHere int
		outOfRange := false
		isCaseUp := 0

		for i = len(arr) - 1; i >= 0; i-- {
			if arr[i] == ' ' || arr[i] == '\n' {
				startFrom = i - 5
				wordsToChange, _ = strconv.Atoi(arr[i+1 : len(arr)-1])
				break
			}
		}

		for i = startFrom; i >= 0; i-- {
			if unicode.IsLetter(rune(arr[i])) {
				startFrom = i
				break
			}
		}
		for i = startFrom; i > 0; i-- {
			if arr[i] == ' ' {
				countSpace++
				for i > 0 {
					if unicode.IsLetter(rune(arr[i])) {
						break
					}
					i--
				}
			}
			if i == 1 || i == 0 {
				if countSpace == wordsToChange {
					arrToChange = arr[i+1:]
					break
				}
				arrToChange = arr
				outOfRange = true
				break
			}
			if countSpace == wordsToChange {
				arrToChange = arr[i+1:]
				break
			}
		}
		switch charType {
		case "up":
			isCaseUp++
			arrToChange = strings.ToUpper(arrToChange)
		case "cap":
			arrToChange = strings.ToLower(arrToChange)
			for i := 0; i < len(arrToChange); i++ {
				if unicode.IsLetter(rune(arrToChange[i])) {
					arrToChange = arrToChange[:i] + strings.ToUpper(string(arrToChange[i])) + arrToChange[i+1:]
					for i < len(arrToChange) {
						if arrToChange[i] == '\'' {
							i++
							continue
						}
						if !unicode.IsLetter(rune(arrToChange[i])) {
							break
						}
						i++
					}
				}
			}
		case "low":
			arrToChange = strings.ToLower(arrToChange)
		}
		for j := len(arrToChange) - 1; j >= 0; j-- {
			if arrToChange[j] == '(' {
				cutHere = j
				break
			}
		}
		if outOfRange {
			return arrToChange[:len(arrToChange)-9+isCaseUp]
		}
		return arr[:i+1] + arrToChange[:cutHere-1]
	})
}
