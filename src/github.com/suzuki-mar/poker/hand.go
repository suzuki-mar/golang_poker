package poker

type Hand struct {
	cards [5]Card
}

func (h Hand) string() string {
	str := ""

	for _, card := range h.cards {
		str += card.string()
	}

	return str
}

type HandParam struct {
	suit        Suit
	numberValue int
}

func buildHand(params [5]HandParam) Hand {

	cards := []Card{}

	for _, param := range params {
		num := Number{value: param.numberValue}
		c := Card{suit: param.suit, number: num}
		cards = append(cards, c)
	}

	var c = [5]Card{}
	copy(c[:], cards)
	return Hand{cards: c}
}
