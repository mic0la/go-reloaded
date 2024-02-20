package cadet

import "strconv"

func checkDigHex(frstdig, secdig byte) bool {
	if (frstdig >= '0' && frstdig <= '9') || (frstdig >= 'A' && frstdig <= 'F') {
		if (secdig >= '0' && secdig <= '9') || (secdig >= 'A' && secdig <= 'F') {
			return true
		}
	}
	return false
}

func HexHandler(s []byte) string {
	var result string
	for i := 0; i < len(s); i++ {
		if i+1 == len(s) {
			result += string(s[i])
			break
		}
		if checkDigHex(s[i], s[i+1]) {
			if string(s[i+2:i+8]) == " (hex)" {
				decDig, _ := strconv.ParseInt(string(s[i])+string(s[i+1]), 16, 64)
				result += strconv.Itoa(int(decDig))
				i += 7
				continue
			}
		}
		result += string(s[i])
	}
	return result
}
