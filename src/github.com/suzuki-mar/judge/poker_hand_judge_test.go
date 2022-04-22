package judge

import (
	"golang/src/github.com/suzuki-mar/hand"
	"testing"
)

func Test_Judge(t *testing.T) {
	tests := []struct {
		name  string
		cards [5]hand.Card
		want  PokerHand
	}{
		{
			name:  "ハイカード",
			cards: buildHighCards(),
			want:  HIGH_CARD,
		},

		{
			name:  "ワンペアー",
			cards: buildOnePairCards(),
			want:  ONE_PAIR,
		},
		{
			name:  "ツーペアー",
			cards: buildTwoPairCards(),
			want:  TWO_PAIR,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if PockerHandJudge(tt.cards) != tt.want {
				t.Errorf("want %v actual %v ", tt.want, PockerHandJudge(tt.cards))
			}
		})
	}
}

func buildOnePairCards() [5]hand.Card {
	var cards [5]hand.Card

	cards[0] = hand.BuildCard(hand.CLOVER, 1)
	cards[1] = hand.BuildCard(hand.HEART, 1)
	cards[2] = hand.BuildCard(hand.CLOVER, 2)
	cards[3] = hand.BuildCard(hand.CLOVER, 3)
	cards[4] = hand.BuildCard(hand.CLOVER, 4)

	return cards
}

func buildTwoPairCards() [5]hand.Card {
	var cards [5]hand.Card

	cards[0] = hand.BuildCard(hand.CLOVER, 1)
	cards[1] = hand.BuildCard(hand.HEART, 1)
	cards[2] = hand.BuildCard(hand.CLOVER, 2)
	cards[3] = hand.BuildCard(hand.SPADE, 2)
	cards[4] = hand.BuildCard(hand.CLOVER, 4)

	return cards
}

func buildHighCards() [5]hand.Card {
	var cards [5]hand.Card

	cards[0] = hand.BuildCard(hand.CLOVER, 1)
	cards[1] = hand.BuildCard(hand.HEART, 9)
	cards[2] = hand.BuildCard(hand.CLOVER, 7)
	cards[3] = hand.BuildCard(hand.CLOVER, 3)
	cards[4] = hand.BuildCard(hand.CLOVER, 4)

	return cards
}
