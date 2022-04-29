package hand

import "math/rand"

type Suit int

const (
	SPADE Suit = iota
	HEART
	CLOVER
	DIAMOND
)

func (s Suit) string() string {
	switch s {
	case SPADE:
		return "♠"
	case HEART:
		return "♥"
	case CLOVER:
		return "♣"
	case DIAMOND:
		return "◆"
	default:
		return "error"
	}
}

func newRandSuit() Suit {
	var value int

	for {
		value = rand.Intn(4)
		if value > 0 {
			break
		}
	}

	return Suit(value)
}
