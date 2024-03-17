package reloaded

import (
	"regexp"
)

func fixPunc(re *regexp.Regexp, str string) string {
	return re.ReplaceAllStringFunc(str, func(arr string) string {
		arr = theTrimSpace(arr)
		arr = trimBackSlashR(arr)
		return arr + " "
	})
}

func theTrimSpace(str string) string {
	result := ""
	for _, v := range str {
		if v == ' ' || v == '\n' {
			continue
		}
		result += string(v)
	}
	return result
}

func trimBackSlashR(str string) string {
	result := ""
	for _, v := range str {
		if v == '\r' {
			continue
		}
		result += string(v)
	}
	return result
}

func fixPunc2(re *regexp.Regexp, str string) string {
	return re.ReplaceAllStringFunc(str, func(arr string) string {
		result := theTrimSpace(arr)
		return result + " "
	})
}

func fixQuote(re *regexp.Regexp, str string) string {
	return re.ReplaceAllStringFunc(str, func(arr string) string {
		quotesAfter := ""
		connector := ""
		if arr[0] != '\'' {
			connector = string(arr[0])
		}
		arr = headSpacesCut(arr)
		headCuttedStr := headSpacesCut(arr[1:])
		return connector + "'" + tailSpacesCut(headCuttedStr[:len(headCuttedStr)-1], quotesAfter) + "'"
	})
}

func headSpacesCut(str string) string {
	if str[0] == ' ' || str[0] == '\n' || str[0] == '\r' {
		return headSpacesCut(str[1:])
	}
	return str
}

func tailSpacesCut(str, quotesAfter string) string {
	if str[len(str)-1] == '\'' {
		return tailSpacesCut(str[:len(str)-1], quotesAfter+"'")
	}
	if str[len(str)-1] == ' ' || str[len(str)-1] == '\n' || str[len(str)-1] == '\r' {
		return tailSpacesCut(str[:len(str)-1], quotesAfter)
	}
	return str + quotesAfter
}

func fixQuoteEnd(re *regexp.Regexp, str string) string {
	return re.ReplaceAllStringFunc(str, func(arr string) string {
		return tailSpacesCut(arr[:len(arr)-1], "'")
	})
}

func CorrectAll(str string) string {
	reHex := regexp.MustCompile(`\b[ ]*[a-fA-F0-9]+[\s,!.\[\]{}:;']*\(hex\)`)
	reBin := regexp.MustCompile(`\b[ ]*[0-1]+[\s,!.\[\]{}:;']*\(bin\)`)
	rePunc := regexp.MustCompile(`[\s^.?!]*[.,,,!,?,:;]\s*`)
	rePunc2 := regexp.MustCompile(`([?!.]\s*)+`)
	reQuotes := regexp.MustCompile(`(\s)*'\s*.*\s*'`)
	reAn := regexp.MustCompile(`\s[Aa]\s+\w\w+`)
	reQuoteEnd := regexp.MustCompile(`'.*[!.,?]\s'`)

	result := SetNums(reBin, str, 2)
	result = SetNums(reHex, result, 16)
	result = SetNums(reBin, result, 2)
	result = fixQuote(reQuotes, result)
	result = FixAn(reAn, result)
	result = fixPunc(rePunc, result)
	result = fixPunc2(rePunc2, result)
	result = fixQuoteEnd(reQuoteEnd, result)
	result = HandleCluMany(result)
	//result = HandleClu(result)
	//result = headSpacesCut(tailSpacesCut(result, ""))
	result = SetNums(reBin, result, 2)
	result = SetNums(reHex, result, 16)
	result = SetNums(reBin, result, 2)
	result = clean(result)
	result = fixPunc(rePunc, result)
	result = fixPunc2(rePunc2, result)
	result = fixQuoteEnd(reQuoteEnd, result)
	return result
}
