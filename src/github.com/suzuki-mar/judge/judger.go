package judge

import "golang/src/github.com/suzuki-mar/hand"

type PokerCard interface {
	String() string
	IsSameSuit(compare hand.Card) bool
	IsSameNumber(compare hand.Card) bool
}

type Judger interface {
	Judge(cards [5]hand.Card) bool
	Rank() PokerHand
}

type PokerHand int

const (
	HIGH_CARD PokerHand = iota
	ONE_PAIR
	TWO_PAIR
	THREE_CARD
	STRAIGHT
	FOUR_OF_CARDS
	FLUSH
	STRAIGHT_FLUSH
	ROYAL_STRAIGHT_FLUSH
)
