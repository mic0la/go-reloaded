package reloaded

func EmptyCheck(result string) string {
	for i := 0; i < len(result)-1; i++ {
		if result[i] == '(' {
			switch {
			case result[i+1] == 'b' || result[i+1] == 'B':
				if result[i+2] == 'i' || result[i+2] == 'I' {
					if i+5 < len(result)-1 {
						if result[i+5] == ' ' {
							result = result[:i-1] + result[i+5:]
							continue
						}
						result = result[:i] + result[i+5:]
					}
				}
			case result[i+1] == 'h' || result[i+1] == 'H':
				if result[i+2] == 'e' || result[i+2] == 'E' {
					if i+5 < len(result)-1 {
						if result[i+5] == ' ' {
							result = result[:i-1] + result[i+5:]
							continue
						}
					}
					result = result[:i] + result[i+5:]
				}
			case result[i+1] == 'c' || result[i+1] == 'C':
				if result[i+2] == 'a' || result[i+2] == 'A' {
					result = result[:i] + result[i+5:]
				}
			case result[i+1] == 'l' || result[i+1] == 'L':
				if result[i+2] == 'o' || result[i+2] == 'O' {
					result = result[:i] + result[i+5:]
				}
			case result[i+1] == 'u' || result[i+1] == 'U':
				if result[i+2] == 'p' || result[i+2] == 'P' {
					result = result[:i] + result[i+4:]
				}
			}
		}
	}
	return result
}
