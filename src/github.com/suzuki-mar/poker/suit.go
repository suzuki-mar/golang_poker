package poker

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
