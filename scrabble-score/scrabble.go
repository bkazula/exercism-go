package scrabble

import "strings"

var characters = map[int]string{
	1:  "A, E, I, O, U, L, N, R, S, T",
	2:  "D, G",
	3:  "B, C, M, P",
	4:  "F, H, V, W, Y",
	5:  "K",
	8:  "J, X",
	10: "Q, Z",
}

func Score(word string) (score int) {
	if len(word) == 0 {
		return
	}
	word = strings.ToUpper(word)
	for _, char := range word {
		for points, chars := range characters {
			if strings.Contains(chars, string(char)) {
				score += points
			}
		}
	}

	return
}
