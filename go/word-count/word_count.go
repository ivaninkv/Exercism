package wordcount

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	freq := make(Frequency)

	phrase = strings.ToLower(phrase)

	var word strings.Builder
	for _, ch := range phrase {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) || ch == '\'' {
			word.WriteRune(ch)
		} else {
			if word.Len() > 0 {
				w := word.String()

				w = strings.Trim(w, "'")
				if len(w) > 0 {
					freq[w]++
				}
				word.Reset()
			}
		}
	}

	if word.Len() > 0 {
		w := word.String()
		w = strings.Trim(w, "'")
		if len(w) > 0 {
			freq[w]++
		}
	}

	return freq
}
