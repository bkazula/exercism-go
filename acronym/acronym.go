package acronym

import "strings"

func Abbreviate(s string) string {
	var ac string
	s = strings.Replace(s, "-", " ", -1)
	words := strings.Split(s, " ")
	for _, word := range words {
		ac += string(word[0])
	}
	return strings.ToUpper(ac)
}
