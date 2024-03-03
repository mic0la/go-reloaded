package reloaded

import (
	"regexp"
	"strings"
	"unicode"
	// "golang.org/x/text/cases"
	// "golang.org/x/text/language"
)

func fixPunc(re *regexp.Regexp, str string) string {
	return re.ReplaceAllStringFunc(str, func(arr string) string {
		arr = theTrimSpace(arr)
		arr = trimBaclSlashR(arr)
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

func trimBaclSlashR(str string) string {
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
		arr = headSpacesCut(arr)
		headCuttedStr := headSpacesCut(arr[1:])
		return "'" + tailSpacesCut(headCuttedStr[:len(headCuttedStr)-1]) + "'"
	})
}

func headSpacesCut(str string) string {
	if str[0] == ' ' || str[0] == '\n' || str[0] == '\r' {
		return headSpacesCut(str[1:])
	}
	return str
}

func tailSpacesCut(str string) string {
	if str[len(str)-1] == ' ' || str[len(str)-1] == '\n' || str[len(str)-1] == '\r' {
		return tailSpacesCut(str[:len(str)-1])
	}
	return str
}

func cluMinus(re *regexp.Regexp, str string) string {
	return re.ReplaceAllString(str, "")
}

func splitString(s string) []string {
	return strings.FieldsFunc(s, func(r rune) bool {
		return unicode.IsSpace(r)
	})
}

func CorrectAll(str string) string {
	reHex := regexp.MustCompile(`\b[ ]*[a-fA-F0-9]+[\s,!.\[\]{}():;']*\(hex\)`)
	reBin := regexp.MustCompile(`\b[ ]*[0-1]+[\s,!.\[\]{}():;']*\(bin\)`)
	reCap := regexp.MustCompile(`[\w'\[\]{}]+[\s,!'.:;]*\((cap|Cap|CAP)\)`)
	reLow := regexp.MustCompile(`[\w'\[\]{}]+[\s,!'.:;]*\((low|Low|LOW)\)`)
	reUp := regexp.MustCompile(`[\w\'[\]{}]+[\s,!'.:;]*\((up|UP|Up)\)`)
	reMultipleChars := regexp.MustCompile(`.*(\((cap|CAP|Cap|UP|up|Up|Low|LOW|low),\s(\d+)\)\s*){2,}`)
	reCapMany := regexp.MustCompile(`(.|\n)*\((cap|Cap|CAP),\s(\d+)\)`)
	reUpMany := regexp.MustCompile(`(.|\n)*\((up|UP|Up),\s(\d+)\)`)
	reLowMany := regexp.MustCompile(`(.|\n)*\((low|Low|LOW),\s(\d+)\)`)
	rePunc := regexp.MustCompile(`[\s^.?!]*[.,,,!,?,:;]\s*`)
	rePunc2 := regexp.MustCompile(`[?!.]\s*[?!.]\s*[?!.]\s*`)
	reQuotes := regexp.MustCompile(`\s+'\s*[^']*\s*'`)
	reAn := regexp.MustCompile(`\s[Aa]\s+\w\w+`)
	reCluMinus := regexp.MustCompile(`\s{0,1}\((cap|low|up),\s*-(\d+)\)`)

	words := splitString(str)
	switch words[0] {
	case "(low,":
		str = str[9:]
	case "(up,":
		str = str[8:]
	case "(cap,":
		str = str[9:]
	case "(low)":
		str = str[6:]
	case "(up)":
		str = str[5:]
	case "(cap)":
		str = str[6:]
	}
	lowManyCount := 0
	upManyCount := 0
	capManyCount := 0
	for _, v := range words {
		switch v {
		case "(low,":
			lowManyCount++
		case "(up,":
			upManyCount++
		case "(cap,":
			capManyCount++
		}
	}
	lowCount := 0
	upCount := 0
	capCount := 0
	for _, v := range words {
		switch v {
		case "(low)":
			lowCount++
		case "(up)":
			upCount++
		case "(cap)":
			capCount++
		}
	}
	result := MultipleChars(reMultipleChars, str, reLowMany, reUpMany, reCapMany)
	result = cluMinus(reCluMinus, result)
	result = SetNums(reBin, result, 2)
	result = SetNums(reHex, result, 16)
	for a := 0; a <= upCount; a++ {
		result = SetChars(reUp, result, "up")
	}
	for a := 0; a <= lowCount; a++ {
		result = SetChars(reLow, result, "low")
	}
	for a := 0; a <= capCount; a++ {
		result = SetChars(reCap, result, "cap")
	}
	result = SetChars(reCap, result, "cap")
	result = SetChars(reLow, result, "low")
	result = SetChars(reUp, result, "up")
	result = fixQuote(reQuotes, result)
	result = fixPunc(rePunc, result)
	result = FixAn(reAn, result)
	for i := 0; i <= upManyCount; i++ {
		result = SetCharsMany(reUpMany, result, "up")
	}
	for j := 0; j <= capManyCount; j++ {
		result = SetCharsMany(reCapMany, result, "cap")
	}
	for k := 0; k <= lowManyCount; k++ {
		result = SetCharsMany(reLowMany, result, "low")
	}
	for a := 0; a <= upCount; a++ {
		result = SetChars(reUp, result, "up")
	}
	for a := 0; a <= lowCount; a++ {
		result = SetChars(reLow, result, "low")
	}
	for a := 0; a <= capCount; a++ {
		result = SetChars(reCap, result, "cap")
	}
	for a := 0; a <= upCount; a++ {
		result = SetChars(reUp, result, "up")
	}
	for a := 0; a <= lowCount; a++ {
		result = SetChars(reLow, result, "low")
	}
	for a := 0; a <= capCount; a++ {
		result = SetChars(reCap, result, "cap")
	}
	result = fixPunc2(rePunc2, result)
	result = EmptyCheck(result)

	return result
}
