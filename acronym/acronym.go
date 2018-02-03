package acronym

import (
	"strings"
	"unicode"
)

func Abbreviate(s string) string {
	var ac string
	words := strings.FieldsFunc(s, shouldSplit)
	for _, word := range words {
		ac += string(word[0])
	}
	return strings.ToUpper(ac)
}

func shouldSplit(r rune) bool {
	return unicode.IsPunct(r) || unicode.IsSpace(r)
}
