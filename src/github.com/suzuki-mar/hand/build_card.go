package hand

import (
	"math/rand"
	"time"
)

func BuildCard(suit Suit, numberValue int) Card {

	number := number{}
	number.value = numberValue

	return Card{suit: suit, number: number}
}

func BuildRandCard() Card {
	rand.Seed(time.Now().UnixNano())

	return BuildCard(buildRandSuit(), buildRandNumber())
}

func buildRandNumber() int {
	var value int

	for {
		value = rand.Intn(13)
		if value > 0 {
			break
		}
	}

	return value
}

func buildRandSuit() Suit {
	var value int

	for {
		value = rand.Intn(4)
		if value > 0 {
			break
		}
	}

	return Suit(value)
}
