package hand

import "strings"

type Hand struct {
	cards [5]Card
}

type Suit int

const (
	SPADE Suit = iota
	HEART
	CLOVER
	DIAMOND
)

func (h Hand) String() string {
	str := ""

	for _, card := range h.cards {
		str += card.String() + " "
	}

	return strings.TrimRight(str, " ")
}

func BuildHandWithCards(cards [5]Card) Hand {

	return Hand{cards: cards}
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
