package scrabble

import "strings"

var arrayScores = [26]int{1, 3, 3, 2, 1, 4, 2, 4, 1, 8, 5, 1, 3, 1, 1, 3, 10, 1, 1, 1, 1, 4, 4, 8, 4, 10}

func scoreSwitch(r rune) int {
	switch r {
	case 'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't':
		return 1
	case 'd', 'g':
		return 2
	case 'b', 'c', 'm', 'p':
		return 3
	case 'f', 'h', 'v', 'w', 'y':
		return 4
	case 'k':
		return 5
	case 'j', 'x':
		return 8
	case 'q', 'z':
		return 10
	default:
		return 0
	}
}

func Score(word string) int {
	score := 0
	for _, c := range strings.ToLower(word) {
		score += scoreSwitch(c)
	}
	return score
}

func ScoreBest(word string) int {
	score := 0
	for _, c := range strings.ToLower(word) {
		score += arrayScores[c-'a']
	}
	return score
}
