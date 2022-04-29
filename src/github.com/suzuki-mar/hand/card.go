package hand

import (
	"math/rand"
	"strconv"
	"time"
)

type Card struct {
	number number
	suit   Suit
}

func NewCard(suit Suit, numberValue int) Card {

	number := number{}
	number.value = numberValue

	return Card{suit: suit, number: number}
}

func NewRandCard() Card {
	rand.Seed(time.Now().UnixNano())

	return NewCard(newRandSuit(), newRandNumber())
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

func newRandNumber() int {
	var value int

	for {
		value = rand.Intn(13)
		if value > 0 {
			break
		}
	}

	return value
}
