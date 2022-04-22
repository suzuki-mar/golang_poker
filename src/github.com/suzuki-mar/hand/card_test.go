package hand

import (
	"testing"
)

func Test_Card_String(t *testing.T) {

	tests := []struct {
		card     Card
		expected string
		name     string
	}{
		{card: BuildCard(CLOVER, 1), expected: "♣A", name: "クローバーのエース"},
		{card: BuildCard(DIAMOND, 2), expected: "◆2", name: "ダイアモンドの2"},
		{card: BuildCard(HEART, 10), expected: "♥10", name: "ハートの10"},
		{card: BuildCard(SPADE, 11), expected: "♠J", name: "スペードのジャック"},
		{card: BuildCard(SPADE, 12), expected: "♠Q", name: "スペードのクイーン"},
		{card: BuildCard(SPADE, 13), expected: "♠K", name: "スペードのキング"},
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

func Test_Card_IsSameValue(t *testing.T) {

	tests := []struct {
		target   Card
		compare  Card
		expected bool
		name     string
	}{
		{target: BuildCard(CLOVER, 1), compare: BuildCard(CLOVER, 1), expected: true, name: "同じ値のカード"},
		{target: BuildCard(CLOVER, 1), compare: BuildCard(CLOVER, 2), expected: false, name: "違う値のカード"},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			// t.Fatalf でテストが失敗した場合でもクリーンアップ処理は呼び出される
			t.Cleanup(func() {
			})

			if tt.target.IsSameValue(tt.compare) != tt.expected {
				t.Errorf("期待した値 %v 実際の値 %v", tt.expected, tt.target.IsSameValue(tt.compare))
			}
		})
	}
}

func Test_Card_IsSameSuit(t *testing.T) {

	tests := []struct {
		target   Card
		compare  Card
		expected bool
		name     string
	}{
		{target: BuildCard(CLOVER, 2), compare: BuildCard(CLOVER, 3), expected: true, name: "同じSuit"},
		{target: BuildCard(CLOVER, 2), compare: BuildCard(HEART, 3), expected: false, name: "違うSuit"},
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

func Test_Card_IsSameNumber(t *testing.T) {

	tests := []struct {
		target   Card
		compare  Card
		expected bool
		name     string
	}{
		{target: BuildCard(CLOVER, 2), compare: BuildCard(HEART, 2), expected: true, name: "同じNumber"},
		{target: BuildCard(CLOVER, 2), compare: BuildCard(CLOVER, 3), expected: false, name: "違うNumber"},
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
