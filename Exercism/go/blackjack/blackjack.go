package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten":
		return 10
	case "jack":
		return 10
	case "queen":
		return 10
	case "king":
		return 10
	default:
		return 0

	}
	panic("Please implement the ParseCard function")
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {

	switch {
	case card1 == "ace" && card2 == "ace":
		return "P"
	case (ParseCard(card1)+ParseCard(card2)) == 21 && (ParseCard(dealerCard) != 10 && dealerCard != "ace"):
		return "W"
	case (ParseCard(card1)+ParseCard(card2)) == 21 && (ParseCard(dealerCard) == 10 && dealerCard == "ace"):
		return "S"
	case (ParseCard(card1) + ParseCard(card2)) == 21:
		return "S"
	case (ParseCard(card1)+ParseCard(card2)) >= 17 && (ParseCard(card1)+ParseCard(card2)) <= 20:
		return "S"
	case (ParseCard(card1)+ParseCard(card2)) >= 12 && (ParseCard(card1)+ParseCard(card2)) <= 16 && ParseCard(dealerCard) < 7:
		return "S"
	case (ParseCard(card1)+ParseCard(card2)) >= 12 && (ParseCard(card1)+ParseCard(card2)) <= 16 && ParseCard(dealerCard) >= 7:
		return "H"
	case (ParseCard(card1) + ParseCard(card2)) <= 11:
		return "H"
	default:
		return "A"
	}
	panic("Please implement the FirstTurn function")
}
