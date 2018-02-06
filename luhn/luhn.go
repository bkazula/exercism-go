package luhn

import "strings"

func Valid(input string) bool {
	var sum int
	var temp int
	input = strings.Replace(input, " ", "", -1)

	// Do not run iteration if value will not be divisible
	if len(input) == 1 {
		return false
	}
	doubled := false
	for i := len(input) - 1; i >= 0; i-- {
		n := int(input[i]) - '0'
		if !doubled {
			temp = n
		} else {
			temp = n * 2
			if temp > 9 {
				temp -= 9
			}
		}
		sum += temp
		doubled = !doubled
	}

	return sum%10 == 0
}
