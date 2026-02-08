package anagram

import (
	"strings"
	"unicode"
)

func Detect(subject string, candidates []string) []string {
	subjRunes := getRuneMap(subject)
	res := []string{}
	for _, candidate := range candidates {
		if len(subject) != len(candidate) {
			continue
		}
		if strings.EqualFold(subject, candidate) {
			continue
		}
		candRunes := getRuneMap(candidate)
		if mapsEqual(subjRunes, candRunes) {
			res = append(res, candidate)
		}
	}

	return res
}

func mapsEqual(a, b map[rune]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}

func getRuneMap(subject string) map[rune]int {
	res := make(map[rune]int)
	for _, r := range subject {
		res[unicode.ToLower(r)]++
	}

	return res
}
