package reloaded

import (
	"strconv"
	"strings"
	"unicode"
)

func lowMany(str string, index int) (string, int) {
	endOfCap := 0
	end := 0
	countSpace := 0
	wordsToChange := 0
	for i := index + 4; i < len(str); i++ {
		if str[i] == ')' {
			endOfCap = i + 1
			end = endOfCap
			to := i - 1
			for str[i] != ',' {
				i--
			}
			wordsToChange, _ = strconv.Atoi(str[i+2 : to+1])
			break
		}
	}
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
	if wordsToChange <= 0 {
		return str, end
	}
	for i := index - 1; i >= 0; i-- {
		if i == 0 {
			str = strings.ToLower(str[:index]) + str[index:]
			break
		}
		if str[i] == ' ' || str[i] == '\n' {
			countSpace++
			for i >= 0 {
				if unicode.IsLetter(rune(str[i])) {
					break
				}
				i--
			}
			if countSpace == wordsToChange {
				str = str[:i+1] + strings.ToLower(str[i+1:index]) + str[index:]
				break
			}
		}
	}
	return str, end
}

func upMany(str string, index int) (string, int) {
	endOfCap := 0
	end := 0
	countSpace := 0
	wordsToChange := 0
	for i := index + 2; i < len(str); i++ {
		if str[i] == ')' {
			endOfCap = i + 1
			end = endOfCap
			to := i - 1
			for str[i] != ',' {
				i--
			}
			wordsToChange, _ = strconv.Atoi(str[i+2 : to+1])
			break
		}
	}
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
	if wordsToChange <= 0 {
		return str, end
	}
	for i := index - 1; i >= 0; i-- {
		if i == 0 {
			str = strings.ToUpper(str[:index]) + str[index:]
			break
		}
		if str[i] == ' ' || str[i] == '\n' {
			countSpace++
			for i >= 0 {
				if unicode.IsLetter(rune(str[i])) {
					break
				}
				i--
			}
			if countSpace == wordsToChange {
				str = str[:i+1] + strings.ToUpper(str[i+1:index]) + str[index:]
				break
			}
		}
	}
	return str, end
}

func capMany(str string, index int) (string, int) {
	endOfCap := 0
	end := 0
	countSpace := 0
	wordsToChange := 0
	for i := index + 4; i < len(str); i++ {
		if str[i] == ')' {
			endOfCap = i + 1
			end = endOfCap
			to := i - 1
			for str[i] != ',' {
				i--
			}
			wordsToChange, _ = strconv.Atoi(str[i+2 : to+1])
			break
		}
	}
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
	if wordsToChange <= 0 {
		return str, end
	}
	for i := index - 1; i >= 0; i-- {
		if i == 0 {
			//str = strings.ToTitle(str[:index]) + str[index:]
			str = strings.ToUpper(string(str[0])) + str[1:]
			for j := 0; j < index-1; j++ {
				if unicode.IsLetter(rune(str[j])) && j != 0 {
					str = str[:j-1] + strings.ToUpper(string(str[j-1])) + str[j:]
				}
				for j < index-1 {
					if unicode.IsSpace(rune(str[j])) {
						for j < index-1 {
							if unicode.IsLetter(rune(str[j])) || unicode.IsDigit(rune(str[j])) {
								break
							}
							j++
						}
						break
					}
					j++
				}
			}
			break
		}
		if str[i] == ' ' || str[i] == '\n' {
			countSpace++
			for i >= 0 {
				if unicode.IsLetter(rune(str[i])) {
					break
				}
				i--
			}
			if countSpace == wordsToChange {
				str = str[:i+1] + Capitalize(str[i+1:index]) + str[index:]
				break
			}
		}
	}
	return str, end
}

func HandleCluMany(str string) string {
	for i := 0; i < len(str); i++ {
		if str[i] == '(' {
			if i+4 > len(str) {
				break
			}
			var cut int
			switch str[i : i+4] {
			case "(up,":
				str, cut = upMany(str, i)
				if i >= (cut-i)+1 {
					i = i - (cut - i)
				}
			}
			if i+5 > len(str) {
				break
			}
			switch str[i : i+5] {
			case "(cap,":
				str, cut = capMany(str, i)
				if i >= (cut-i)+1 {
					i = i - (cut - i)
				}
			case "(low,":
				str, cut = lowMany(str, i)
				if i >= (cut-i)+1 {
					i = i - (cut - i)
				}
			}
		}
	}
	return str
}
