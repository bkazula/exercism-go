package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(phrase string) bool {
	var usedLetters string

	for _, char := range phrase {
		letter := strings.ToLower(string(char))
		if unicode.IsLetter(char) && strings.ContainsAny(letter, usedLetters) {
			return false
		}
		usedLetters += letter
	}

	return true
}
