package luhn

import (
	"strings"
	"unicode/utf8"
)

func Valid(id string) bool {
	if !checkString(id) {
		return false
	}

	sum := 0
	needMulty := false
	for len(id) > 0 {
		r, size := utf8.DecodeLastRuneInString(id)
		if r == ' ' {
			id = id[:len(id)-size]
			continue
		}
		if needMulty {
			s := int(r-'0') * 2
			if s > 9 {
				sum += s/10 + s%10
			} else {
				sum += s
			}
			needMulty = false
		} else {
			sum += int(r - '0')
			needMulty = true
		}
		id = id[:len(id)-size]
	}

	return sum%10 == 0
}

func checkString(str string) bool {
	for _, r := range str {
		if (r < '0' || r > '9') && r != ' ' {
			return false
		}
	}
	return utf8.RuneCountInString(str)-strings.Count(str, " ") > 1
}
