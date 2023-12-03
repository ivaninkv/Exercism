package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var result int
	switch card {
	case "ace":
		result = 11
	case "two":
		result = 2
	case "three":
		result = 3
	case "four":
		result = 4
	case "five":
		result = 5
	case "six":
		result = 6
	case "seven":
		result = 7
	case "eight":
		result = 8
	case "nine":
		result = 9
	case "ten", "jack", "queen", "king":
		result = 10
	default:
		result = 0
	}

	return result
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	var result string
	cardSum := ParseCard(card1) + ParseCard(card2)
	dealerCardValue := ParseCard(dealerCard)

	switch {
	case cardSum == 22:
		result = "P"
	case cardSum == 21:
		if dealerCardValue >= 10 {
			result = "S"
		} else {
			result = "W"
		}
	case cardSum >= 17 && cardSum <= 20:
		result = "S"
	case cardSum >= 12 && cardSum <= 16:
		if dealerCardValue >= 7 {
			result = "H"
		} else {
			result = "S"
		}
	case cardSum <= 11:
		result = "H"
	}
	return result
}
