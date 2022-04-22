package judge

import (
	"golang/src/github.com/suzuki-mar/hand"
)

func PockerHandJudge(cards [5]hand.Card) PokerHand {

	switch calcPairCount(cards) {
	case 1:
		return ONE_PAIR
	case 2:
		return TWO_PAIR
	}

	return HIGH_CARD
}

func calcPairCount(cards [5]hand.Card) int {
	var pairCount = 0.0

	for _, target := range cards {

		for _, compare := range cards {

			if !target.IsSameValue(compare) && target.IsSameNumber(compare) {
				// 2回ペアで1になる
				pairCount += 0.5
			}
		}
	}

	return int(pairCount)
}
