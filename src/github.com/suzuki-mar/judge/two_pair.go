package judge

import "golang/src/github.com/suzuki-mar/hand"

type HighCard struct {
}

func (hc HighCard) Judge(cards [5]hand.Card) PokerHand {
	
	for _, card := range h.cards {
		str += card.String() + " "
	}

	return c.suit.string() + c.number.string()
}

func (hc HighCard) Rank() PokerHand {
	return HIGH_CARD
}
