package reloaded

func EmptyCheck(result string) string {
	for i := 0; i < len(result); i++ {
		if i == len(result)-1 {
			break
		}
		if result[i] == '(' {
			switch {
			case result[i+1] == 'b' || result[i+1] == 'B':
				if result[i+2] == 'i' || result[i+2] == 'I' {
					result = result[:i] + result[i+5:]
				}
			case result[i+1] == 'h' || result[i+1] == 'H':
				if result[i+2] == 'e' || result[i+2] == 'E' {
					result = result[:i] + result[i+5:]
				}
			}
		}
	}
	return result
}
