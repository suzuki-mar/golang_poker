package poker

type Suit int

//* スペード(♠): spade
//* ハート(♥): heart
//* クローバー(♣): clover
//* ダイヤ(◆): diamond

const (
	Spade Suit = iota
	Heart
	Clover
	Diamond
)

func (s Suit) displayString() string {
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
	return s.displayString() == compare.displayString()
}
