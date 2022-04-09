package poker

type Card struct {
	number Number
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
