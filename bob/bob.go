package bob

import (
	"regexp"
	"strings"
)

func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	isUpper := strings.ToUpper(remark) == remark
	isWord, _ := regexp.MatchString("[a-zA-Z]+", remark)

	if len(remark) == 0 {
		return "Fine. Be that way!"
	} else if strings.HasSuffix(remark, "?") && isUpper && isWord {
		return "Calm down, I know what I'm doing!"
	} else if strings.HasSuffix(remark, "?") {
		return "Sure."
	} else if !isWord {
		return "Whatever."
	} else if isUpper {
		return "Whoa, chill out!"
	}
	return "Whatever."
}
