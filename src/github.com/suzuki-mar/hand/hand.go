package hand

import "strings"

type Hand struct {
	cards [5]Card
}

func (h Hand) String() string {
	str := ""

	for _, card := range h.cards {
		str += card.String() + " "
	}

	return strings.TrimRight(str, " ")
}

type HandParam struct {
	suit        Suit
	numberValue int
}

func buildHandWithParams(params [5]HandParam) Hand {

	var cards []Card

	for _, param := range params {
		num := number{value: param.numberValue}
		c := Card{suit: param.suit, number: num}
		cards = append(cards, c)
	}

	var c = [5]Card{}
	copy(c[:], cards)
	return Hand{cards: c}
}

func BuildHand() Hand {
	var cards []Card

	for i := 0; i < 5; i++ {
		cards = append(cards, BuildRandCard())
	}

	var c = [5]Card{}
	copy(c[:], cards)
	return Hand{cards: c}
}
