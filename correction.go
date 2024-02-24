package reloaded

import (
	"regexp"
	"strconv"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func convertNums(re *regexp.Regexp, str string, numTypeName string, numTypeInt int) string {
	return re.ReplaceAllStringFunc(str, func(arr string) string {
		last6Runes := len(arr) - 6
		arr = arr[:last6Runes]
		decDigit, _ := strconv.ParseInt(arr, numTypeInt, 64)
		result := strconv.Itoa(int(decDigit))
		return result
	})
}

func convertChars(re *regexp.Regexp, result string, charType string) string {
	return re.ReplaceAllStringFunc(result, func(arr string) string {
		lastRunes := len(arr) - 6
		if len(charType) == 2 {
			lastRunes++
		}
		arr = arr[:lastRunes]
		var caser cases.Caser
		switch charType {
		case "cap":
			caser = cases.Title(language.English)
		case "low":
			caser = cases.Lower(language.English)
		case "up":
			caser = cases.Upper(language.English)
		}
		arr = caser.String(arr)
		return arr
	})
}

func CorrectAll(str string) string {
	reHex := regexp.MustCompile(`[A-F0-9]+.\(hex\)`)
	reBin := regexp.MustCompile(`[0-1]+.\(bin\)`)
	reCap := regexp.MustCompile(`[a-zA-Z]+[\s,]\(cap\)`)
	reLow := regexp.MustCompile(`[a-zA-Z]+[\s,]\(low\)`)
	reUp := regexp.MustCompile(`[a-zA-Z]+[\s,]\(up\)`)
	// reCapMany := regexp.MustCompile(`(\s*.*)+\(cap,\s\d+\)`)

	result := convertNums(reBin, str, "(bin)", 2)
	result = convertNums(reHex, result, "(hex)", 16)
	result = convertChars(reCap, result, "cap")
	result = convertChars(reLow, result, "low")
	result = convertChars(reUp, result, "up")

	// result = reCapMany.ReplaceAllStringFunc(result, func(arr string) string {
	// 	words := strings.Split(arr, " ")
	// 	fmt.Println(words)

	// 	caser := cases.Title(language.English)
	// 	arr = caser.String(arr)
	// 	return arr
	// })

	return result
}
