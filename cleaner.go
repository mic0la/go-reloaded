package reloaded

func clean(str string) string {
	for i := 0; i < len(str); i++ {
		if str[i] == '(' {
			switch str[i+1 : i+5] {
			case "hex)":
				str = str[:i] + str[i+5:]
			case "bin)":
				str = str[:i] + str[i+5:]
			}
		}
	}
	return str
}
