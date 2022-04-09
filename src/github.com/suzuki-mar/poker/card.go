package poker

import "strconv"

type Card struct {
	number number
	suit   Suit
}

func (c Card) String() string {
	return c.suit.String() + c.number.String()
}

func (c Card) isSameSuit(compare Card) bool {
	return c.suit.isSame(compare.suit)
}

func (c Card) isSameNumber(compare Card) bool {
	return c.number.isSame(compare.number)
}

func (c Card) numberValue() int {
	return c.number.value
}

type Suit int

const (
	Spade Suit = iota
	Heart
	Clover
	Diamond
)

func (s Suit) String() string {
	switch s {
	case Spade:
		return "♠"
	case Heart:
		return "♥"
	case Clover:
		return "♣"
	case Diamond:
		return "◆"
	default:
		return "error"
	}
}

func (s Suit) isSame(compare Suit) bool {
	return s.String() == compare.String()
}

type number struct {
	value int
}

func (n number) String() string {

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

func (n number) isSame(compare number) bool {
	return n.value == compare.value
}
