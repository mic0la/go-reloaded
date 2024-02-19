package cadet

import "fmt"

func CheckChar(v byte) bool {
	if (v >= '0' && v <= '9') || (v >= 'A' && v <= 'F') {
		return true
	}
	return false
}

func HexHandler(s []byte) []byte {
	for i, v := range s {
		if CheckChar(v) && CheckChar(s[i+1]) {
			if string(s[i+2:i+8]) == " (hex)" {
				fmt.Println("reached")
				return []byte{55}
			}
		}
	}
	return s
}
