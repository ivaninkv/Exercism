package bob

import (
	"strings"
	"unicode"
)

func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	if remark == "" {
		return "Fine. Be that way!"
	}

	isQuestion := strings.HasSuffix(remark, "?")
	hasLetters := strings.IndexFunc(remark, unicode.IsLetter) >= 0
	isShout := hasLetters && strings.ToUpper(remark) == remark

	if isShout && isQuestion {
		return "Calm down, I know what I'm doing!"
	}
	if isShout {
		return "Whoa, chill out!"
	}
	if isQuestion {
		return "Sure."
	}
	return "Whatever."
}
