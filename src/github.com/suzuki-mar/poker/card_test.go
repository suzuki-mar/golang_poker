package poker

import (
	"testing"
)

func TestCardDisplayString(t *testing.T) {

	tests := []struct {
		card     Card
		expected string
		name     string
	}{
		{card: Card{Number{2}, Clover}, expected: "♣2", name: "クローバーの2"},
		{card: Card{Number{1}, Diamond}, expected: "◆A", name: "ダイアモンドのエース"},
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
		{target: Card{Number{2}, Clover}, compare: Card{Number{2}, Clover}, expected: true, name: "同じSuit"},
		{target: Card{Number{2}, Clover}, compare: Card{Number{2}, Diamond}, expected: false, name: "違うSuit"},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			// t.Fatalf でテストが失敗した場合でもクリーンアップ処理は呼び出される
			t.Cleanup(func() {
				t.Log("cleanup!")
			})

			if tt.target.isSameSuit(tt.compare) != tt.expected {
				t.Errorf("期待した値 %v 実際の値 %v", tt.expected, tt.target.isSameSuit(tt.compare))
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
		{target: Card{Number{2}, Clover}, compare: Card{Number{2}, Clover}, expected: true, name: "同じNumber"},
		{target: Card{Number{2}, Clover}, compare: Card{Number{3}, Diamond}, expected: false, name: "違うNumber"},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			// t.Fatalf でテストが失敗した場合でもクリーンアップ処理は呼び出される
			t.Cleanup(func() {
				t.Log("cleanup!")
			})

			if tt.target.isSameNumber(tt.compare) != tt.expected {
				t.Errorf("期待した値 %v 実際の値 %v", tt.expected, tt.target.isSameNumber(tt.compare))
			}
		})
	}
}
