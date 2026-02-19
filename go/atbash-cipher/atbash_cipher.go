package atbash

import "strings"

func Atbash(s string) string {
	if len(s) == 0 {
		return ""
	}

	var result strings.Builder
	result.Grow(len(s) + len(s)/5 + 1)

	charCount := 0
	bytes := []byte(s)

	for i := range len(bytes) {
		b := bytes[i]
		var encoded byte

		if b >= 'a' && b <= 'z' {
			encoded = 'a' + ('z' - b)
		} else if b >= 'A' && b <= 'Z' {
			encoded = 'a' + ('Z' - b)
		} else if b >= '0' && b <= '9' {
			encoded = b
		} else {
			continue
		}

		if charCount > 0 && charCount%5 == 0 {
			result.WriteByte(' ')
		}
		result.WriteByte(encoded)
		charCount++
	}

	return result.String()
}
