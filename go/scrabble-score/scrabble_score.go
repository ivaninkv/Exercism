package scrabble

import "strings"

var scores = map[string]int{
	"aeioulnrst": 1,
	"dg":         2,
	"bcmp":       3,
	"fhvwy":      4,
	"k":          5,
	"jx":         8,
	"qz":         10,
}

func Score(word string) int {
	score := 0

	for _, c := range strings.ToLower(word) {
		for pattern, value := range scores {
			if strings.ContainsRune(pattern, c) {
				score += value
				break
			}
		}
	}

	return score
}
