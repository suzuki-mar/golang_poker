package poker

import (
	"math/rand"
	"strconv"
	"time"
)

type Card struct {
	number number
	suit   Suit
}

func (c Card) String() string {
	return c.suit.string() + c.number.string()
}

func (c Card) IsSameSuit(compare Card) bool {
	return c.suit.isSame(compare.suit)
}

func (c Card) IsSameNumber(compare Card) bool {
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

func (s Suit) string() string {
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
	return s.string() == compare.string()
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

func (n number) isSame(compare number) bool {
	return n.value == compare.value
}

func BuildRandCard() Card {
	rand.Seed(time.Now().UnixNano())

	return Card{suit: buildRandSuit(), number: buildRandNumber()}
}

func buildRandNumber() number {
	var value int

	for {
		value := rand.Intn(13)
		if value > 0 {
			break
		}
	}

	return number{value: value}
}

func buildRandSuit() Suit {
	var value int

	for {
		value := rand.Intn(4)
		if value > 0 {
			break
		}
	}

	return Suit(value)
}
