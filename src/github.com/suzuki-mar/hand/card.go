package hand

import (
	"strconv"
)

type Card struct {
	number number
	suit   Suit
}

func (c Card) String() string {
	return c.suit.string() + c.number.string()
}

func (c Card) IsSameSuit(compare Card) bool {
	return c.suit.string() == compare.suit.string()
}

func (c Card) IsSameNumber(compare Card) bool {
	return c.number.value == compare.number.value
}

func (c Card) IsSameValue(compare Card) bool {
	return c.IsSameNumber(compare) && c.IsSameSuit(compare)
}

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

type number struct {
	value int
}

func (n number) string() string {

	switch n.value {
	case 1:
		return "A"
	case 11:
		return "J"
	case 12:
		return "Q"
	case 13:
		return "K"
	default:
		return strconv.Itoa(n.value)
	}
}
