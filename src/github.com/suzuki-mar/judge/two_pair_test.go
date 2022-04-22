package judge

import (
	"golang/src/github.com/suzuki-mar/hand"
	"testing"
)

func TestTwoPair_Judge(t *testing.T) {
	tests := []struct {
		name  string
		cards [5]hand.Card
		want  bool
	}{
		{
			name:  "同じカードが2枚ある場合",
			cards: buildPairCards(),
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tp := TwoPair{}
			if got := tp.Judge(tt.cards); got != tt.want {
				t.Errorf("Judge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func buildPairCards() [5]hand.Card {
	var cards [5]hand.Card

	cards[0] = hand.BuildCard(hand.CLOVER, 1)
	cards[1] = hand.BuildCard(hand.HEART, 1)
	cards[2] = hand.BuildCard(hand.CLOVER, 2)
	cards[3] = hand.BuildCard(hand.CLOVER, 3)
	cards[4] = hand.BuildCard(hand.CLOVER, 4)

	return cards
}
