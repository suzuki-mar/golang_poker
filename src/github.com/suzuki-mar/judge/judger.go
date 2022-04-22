package judge

import "golang/src/github.com/suzuki-mar/hand"

type PokerCard interface {
	String() string
	IsSameSuit(compare hand.Card) bool
	IsSameNumber(compare hand.Card) bool
}

type PokerHand int

const (
	HighCard PokerHand = iota
	OnePair
	TwoPair
	Diamond
)

type judge interface {
}
