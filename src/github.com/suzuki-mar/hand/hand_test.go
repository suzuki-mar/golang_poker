package hand

import (
	"testing"
)

func TestHandString(t *testing.T) {

	var cards [5]Card

	cards[0] = BuildCard(CLOVER, 2)
	cards[1] = BuildCard(CLOVER, 3)
	cards[2] = BuildCard(CLOVER, 4)
	cards[3] = BuildCard(CLOVER, 5)
	cards[4] = BuildCard(CLOVER, 6)

	tests := []struct {
		hand     Hand
		expected string
	}{
		{hand: BuildHandWithCards(cards), expected: "♣2 ♣3 ♣4 ♣5 ♣6"},
	}
	for _, tt := range tests {
		t.Run("string()", func(t *testing.T) {
			// t.Fatalf でテストが失敗した場合でもクリーンアップ処理は呼び出される
			t.Cleanup(func() {
				t.Log("cleanup!")
			})

			if tt.hand.String() != tt.expected {
				t.Errorf("期待した値 %v 実際の値 %v", tt.expected, tt.hand.String())
			}
		})
	}
}

func TestHandBuildHand(t *testing.T) {
	t.Run("buildBuildRandHand()", func(t *testing.T) {
		BuildHand()
	})
}
