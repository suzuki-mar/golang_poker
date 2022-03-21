package poker

import (
	"testing"
)

func TestSuitDisplayString(t *testing.T) {
	tests := []struct {
		suit     Suit
		expected string
		name     string
	}{
		{suit: Spade, expected: "♠", name: "スペード"},
		{suit: Heart, expected: "♥", name: "ハート"},
		{suit: Diamond, expected: "◆", name: "ダイヤモンド"},
		{suit: Clover, expected: "♣", name: "クローバー"},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			// t.Fatalf でテストが失敗した場合でもクリーンアップ処理は呼び出される
			t.Cleanup(func() {
				t.Log("cleanup!")
			})

			if tt.suit.displayString() != tt.expected {
				t.Errorf("期待した値 %v 実際の値 %v", tt.expected, tt.suit.displayString())
			}
		})
	}
}

func TestSuitIsSame(t *testing.T) {
	tests := []struct {
		target   Suit
		compare  Suit
		expected bool
	}{
		{target: Clover, compare: Clover, expected: true},
		{target: Clover, compare: Diamond, expected: false},
	}
	for _, tt := range tests {
		t.Run(tt.target.displayString(), func(t *testing.T) {
			// t.Fatalf でテストが失敗した場合でもクリーンアップ処理は呼び出される
			t.Cleanup(func() {
				t.Log("cleanup!")
			})

			if tt.target.isSame(tt.compare) != tt.expected {
				t.Errorf("期待した値 %v 実際の値 %v", tt.expected, tt.target.isSame(tt.compare))
			}
		})
	}
}
