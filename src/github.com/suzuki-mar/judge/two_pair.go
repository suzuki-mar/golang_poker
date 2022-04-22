package judge

import "golang/src/github.com/suzuki-mar/hand"

type TwoPair struct {
}

func (tp TwoPair) Judge(cards [4]hand.Card) bool {
	pair := map[string]int{}
	c := cards[0]

	pair[c.NumberString()] = 1

	return true
}

func (tp TwoPair) Rank() PokerHand {
	return HIGH_CARD
}
