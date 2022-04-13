package hand

import (
	"testing"
)

func TestCardString(t *testing.T) {

	tests := []struct {
		card     Card
		expected string
		name     string
	}{
		{card: Card{number{1}, CLOVER}, expected: "♣A", name: "クローバーのエース"},
		{card: Card{number{2}, DIAMOND}, expected: "◆2", name: "ダイアモンドの2"},
		{card: Card{number{10}, HEART}, expected: "♥10", name: "ハートの10"},
		{card: Card{number{11}, SPADE}, expected: "♠J", name: "スペードのジャック"},
		{card: Card{number{12}, SPADE}, expected: "♠Q", name: "スペードのクイーン"},
		{card: Card{number{13}, SPADE}, expected: "♠K", name: "スペードのキング"},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			// t.Fatalf でテストが失敗した場合でもクリーンアップ処理は呼び出される
			t.Cleanup(func() {
				t.Log("cleanup!")
			})

			if tt.card.String() != tt.expected {
				t.Errorf("期待した値 %v 実際の値 %v", tt.expected, tt.card.String())
			}
		})
	}
}

func TestCardIsSameSuit(t *testing.T) {

	tests := []struct {
		target   Card
		compare  Card
		expected bool
		name     string
	}{
		{target: Card{number{2}, CLOVER}, compare: Card{number{2}, CLOVER}, expected: true, name: "同じSuit"},
		{target: Card{number{2}, CLOVER}, compare: Card{number{2}, DIAMOND}, expected: false, name: "違うSuit"},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			// t.Fatalf でテストが失敗した場合でもクリーンアップ処理は呼び出される
			t.Cleanup(func() {
				t.Log("cleanup!")
			})

			if tt.target.IsSameSuit(tt.compare) != tt.expected {
				t.Errorf("期待した値 %v 実際の値 %v", tt.expected, tt.target.IsSameSuit(tt.compare))
			}
		})
	}
}

func TestCardIsSameNumber(t *testing.T) {

	tests := []struct {
		target   Card
		compare  Card
		expected bool
		name     string
	}{
		{target: Card{number{2}, CLOVER}, compare: Card{number{2}, CLOVER}, expected: true, name: "同じNumber"},
		{target: Card{number{2}, CLOVER}, compare: Card{number{3}, DIAMOND}, expected: false, name: "違うNumber"},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			// t.Fatalf でテストが失敗した場合でもクリーンアップ処理は呼び出される
			t.Cleanup(func() {
				t.Log("cleanup!")
			})

			if tt.target.IsSameNumber(tt.compare) != tt.expected {
				t.Errorf("期待した値 %v 実際の値 %v", tt.expected, tt.target.IsSameNumber(tt.compare))
			}
		})
	}
}
