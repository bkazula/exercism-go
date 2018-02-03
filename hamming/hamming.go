package hamming

import (
	"errors"
	"strings"
)

func Distance(a, b string) (int, error) {
	count := 0
	if len(a) != len(b) {
		return -1, errors.New("Codes are not equal")
	}
	if strings.Compare(a, b) == 0 {
		return 0, nil
	}
	for index, char := range b {
		if string(a[index]) != string(char) {
			count++
		}
	}
	return count, nil
}
