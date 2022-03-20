package poker

import (
	"strconv"
)

type Number struct {
	value int
}

func (n Number) displayString() string {

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

func (n Number) isSame(compare Number) bool {
	return n.value == compare.value
}
