package romannumerals

import (
	"fmt"
	"strings"
)

func ToRomanNumeral(input int) (string, error) {
	if input < 1 || input > 3999 {
		return "", fmt.Errorf("input must be between 1 and 3999")
	}

	result := strings.Builder{}
	for input > 0 {
		symbol, value := getNearArabic(input)
		result.WriteString(symbol)
		input -= value
	}

	return result.String(), nil
}

func getNearArabic(input int) (string, int) {
	switch {
	case input >= 1000:
		return "M", 1000
	case input >= 900:
		return "CM", 900
	case input >= 500:
		return "D", 500
	case input >= 400:
		return "CD", 400
	case input >= 100:
		return "C", 100
	case input >= 90:
		return "XC", 90
	case input >= 50:
		return "L", 50
	case input >= 40:
		return "XL", 40
	case input >= 10:
		return "X", 10
	case input >= 9:
		return "IX", 9
	case input >= 5:
		return "V", 5
	case input >= 4:
		return "IV", 4
	case input >= 1:
		return "I", 1
	default:
		return "", 0
	}
}
