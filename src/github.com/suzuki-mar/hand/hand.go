package hand

import "strings"

type Hand struct {
	cards [5]Card
}

func (h Hand) String() string {
	var sb strings.Builder

	for _, card := range h.cards {
		sb.WriteString(card.String())
		sb.WriteString(" ")
	}

	str := sb.String()
	return strings.TrimRight(str, " ")
}

func NewHandWithCards(cards [5]Card) Hand {

	return Hand{cards: cards}
}

func NewHand() Hand {
	var cards []Card

	for i := 0; i < 5; i++ {
		cards = append(cards, NewRandCard())
	}

	var c = [5]Card{}
	copy(c[:], cards)
	return Hand{cards: c}
}
