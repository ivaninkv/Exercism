package isogram

import "unicode"

func IsIsogram(word string) bool {
	var arr [26]int
	for _, c := range word {
		if unicode.IsLetter(c) {
			if c >= 'A' && c <= 'Z' {
				c |= 32
			}
			if arr[c-'a'] == 1 {
				return false
			}
			arr[c-'a'] = 1
		}
	}
	return true
}
