package reloaded

import "unicode"

func Capitalize(str string) string {
	// Handle empty string
	if str == "" {
		return str
	}

	// Convert string to rune slice for easier manipulation
	runes := []rune(str)

	// isPreviousLetter checks if the previous rune was a letter
	isPreviousLetter := false

	// Loop through each rune in the string
	for i, r := range runes {
		// Check if the current rune is a letter and the previous wasn't
		if unicode.IsLetter(r) && !isPreviousLetter {
			// Check if previous rune is an apostrophe before capitalizing
			if i > 0 && runes[i-1] == rune('\'') {
				runes[i] = r // Don't capitalize after apostrophe in contractions
			} else {
				runes[i] = unicode.ToUpper(r) // Capitalize the first letter of each word
			}
		}
		isPreviousLetter = unicode.IsLetter(r) // Update previous letter check
	}

	// Convert the rune slice back to string
	return string(runes)
}
