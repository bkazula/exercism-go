package bob

import (
	"regexp"
	"strings"
)

func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	isUpper := strings.ToUpper(remark) == remark
	isWord, _ := regexp.MatchString("[a-zA-Z]+", remark)

	switch {
	case len(remark) == 0:
		return "Fine. Be that way!"
	case strings.HasSuffix(remark, "?") && isUpper && isWord:
		return "Calm down, I know what I'm doing!"
	case strings.HasSuffix(remark, "?"):
		return "Sure."
	case !isWord:
		return "Whatever."
	case isUpper:
		return "Whoa, chill out!"
	default:
		return "Whatever."
	}
}
