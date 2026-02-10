package wordcount

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	freq := make(Frequency)

	lower := strings.ToLower(phrase)

	start := -1
	for i, ch := range lower {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) || ch == '\'' {
			if start == -1 {
				start = i
			}
		} else {
			if start != -1 {
				word := lower[start:i]
				word = strings.Trim(word, "'")
				if len(word) > 0 {
					freq[word]++
				}
				start = -1
			}
		}
	}

	if start != -1 {
		word := lower[start:]
		word = strings.Trim(word, "'")
		if len(word) > 0 {
			freq[word]++
		}
	}

	return freq
}
