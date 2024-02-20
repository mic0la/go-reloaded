package cadet

import "strconv"

func checkDigBin(frstdig, secdig byte) bool {
	if frstdig >= '0' && frstdig <= '1' {
		if secdig >= '0' && secdig <= '1' {
			return true
		}
	}
	return false
}

func BinHandler(s []byte) string {
	var result string
	for i := 0; i < len(s); i++ {
		if i+1 == len(s) {
			result += string(s[i])
			break
		}
		if checkDigBin(s[i], s[i+1]) {
			if string(s[i+2:i+8]) == " (bin)" {
				decDig, _ := strconv.ParseInt(string(s[i])+string(s[i+1]), 2, 64)
				result += strconv.Itoa(int(decDig))
				i += 7
				continue
			}
		}
		result += string(s[i])
	}
	return result
}
