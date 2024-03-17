package reloaded

import (
	"strings"
	"unicode"
)

func cap(str string, index int) string {
	endOfCap := index + 5
	for i := index - 1; i > 0; i-- {
		if str[i] == ' ' || str[i] == '\n' {
			index = i
		} else {
			break
		}
	}
	str = str[:index] + str[endOfCap:]
	if len(str) > index {
		if str[index] == ' ' || str[index] == '\n' {
			for i := index; i < len(str); i++ {
				if str[i] == ' ' || str[i] == '\n' {
					endOfCap = i
				} else {
					break
				}
			}
			str = str[:index] + str[endOfCap+1:]
		}
		if index != 0 {
			str = str[:index] + " " + str[index:]
		}
	}
	for i := index - 1; i >= 0; i-- {
		if i == 0 {
			str = strings.ToUpper(string(str[0])) + strings.ToLower(str[1:index]) + str[index:]
			break
		}
		if str[i] == ' ' || str[i] == '\n' {
			if unicode.IsLetter(rune(str[i+1])) || unicode.IsDigit(rune(str[i+1])) {
				str = str[:i+1] + strings.ToUpper(string(str[i+1])) + strings.ToLower(str[i+2:index]) + str[index:]
			} else {
				str = str[:i+2] + strings.ToUpper(string(str[i+2])) + strings.ToLower(str[i+3:index]) + str[index:]
			}
			break
		}
	}
	return str
}

func low(str string, index int) string {
	endOfCap := index + 5
	for i := index - 1; i > 0; i-- {
		if str[i] == ' ' || str[i] == '\n' {
			index = i
		} else {
			break
		}
	}
	str = str[:index] + str[endOfCap:]
	if len(str) > index {
		if str[index] == ' ' || str[index] == '\n' {
			for i := index; i < len(str); i++ {
				if str[i] == ' ' || str[i] == '\n' {
					endOfCap = i
				} else {
					break
				}
			}
			str = str[:index] + str[endOfCap+1:]
		}
		if index != 0 {
			str = str[:index] + " " + str[index:]
		}
	}
	for i := index - 1; i >= 0; i-- {
		if i == 0 {
			str = strings.ToLower(str[:index]) + str[index:]
			break
		}
		if str[i] == ' ' || str[i] == '\n' {
			str = str[:i] + strings.ToLower(str[i:index]) + str[index:]
			break
		}
	}

	return str
}

func up(str string, index int) string {
	endOfCap := index + 4
	for i := index - 1; i > 0; i-- {
		if str[i] == ' ' || str[i] == '\n' {
			index = i
		} else {
			break
		}
	}
	str = str[:index] + str[endOfCap:]

	if len(str) > index {
		if str[index] == ' ' || str[index] == '\n' {
			for i := index; i < len(str); i++ {
				if str[i] == ' ' || str[i] == '\n' {
					endOfCap = i
				} else {
					break
				}
			}
			str = str[:index] + str[endOfCap+1:]
		}
		if index != 0 {
			str = str[:index] + " " + str[index:]
		}
	}
	for i := index - 1; i >= 0; i-- {
		if i == 0 {
			str = strings.ToUpper(str[:index]) + str[index:]
			break
		}
		if str[i] == ' ' || str[i] == '\n' {
			str = str[:i] + strings.ToUpper(str[i:index]) + str[index:]
			break
		}
	}
	return str
}

// func HandleClu(str string) string {
// 	for i := 0; i < len(str); i++ {
// 		if str[i] == '(' {
// 			if i+4 > len(str) {
// 				break
// 			}
// 			switch str[i : i+4] {
// 			case "(up)":
// 				str = up(str, i)
// 				if i >= 4 {
// 					i = i - 4
// 				}
// 			}
// 			if i+5 > len(str) {
// 				break
// 			}
// 			switch str[i : i+5] {
// 			case "(cap)":
// 				str = cap(str, i)
// 				if i >= 5 {
// 					i = i - 4
// 				}
// 			case "(low)":
// 				str = low(str, i)
// 				if i >= 5 {
// 					i = i - 4
// 				}
// 			}
// 			if i+4 > len(str) {
// 				break
// 			}
// 			switch str[i : i+4] {
// 			case "(UP)":
// 				str = up(str, i)
// 				if i >= 4 {
// 					i = i - 4
// 				}
// 			}
// 			if i+5 > len(str) {
// 				break
// 			}
// 			switch str[i : i+5] {
// 			case "(CAP)":
// 				str = cap(str, i)
// 				if i >= 5 {
// 					i = i - 4
// 				}
// 			case "(LOW)":
// 				str = low(str, i)
// 				if i >= 5 {
// 					i = i - 4
// 				}
// 			}
// 			if i+4 > len(str) {
// 				break
// 			}
// 			switch str[i : i+4] {
// 			case "(Up)":
// 				str = up(str, i)
// 				if i >= 4 {
// 					i = i - 4
// 				}
// 			}
// 			if i+5 > len(str) {
// 				break
// 			}
// 			switch str[i : i+5] {
// 			case "(Cap)":
// 				str = cap(str, i)
// 				if i >= 5 {
// 					i = i - 4
// 				}
// 			case "(Low)":
// 				str = low(str, i)
// 				if i >= 5 {
// 					i = i - 4
// 				}
// 			}
// 		}
// 	}
// 	return str
// }
