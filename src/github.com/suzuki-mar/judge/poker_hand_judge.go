package judge

import (
	"golang/src/github.com/suzuki-mar/hand"
)

func PockerHandJudge(cards [5]hand.Card) PokerHand {

	var hasPair = false

	for _, target := range cards {

		for _, compare := range cards {

			if !target.IsSameValue(compare) && target.IsSameNumber(compare) {
				println("hoge")
				hasPair = true
				println(hasPair)
			}
		}
	}

	println(hasPair)

	if hasPair {
		return TWO_PAIR
	} else {
		return HIGH_CARD
	}

}
