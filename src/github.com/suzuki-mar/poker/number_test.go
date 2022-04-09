package poker

import (
	"strconv"
	"testing"
)

func TestNumberDisplay(t *testing.T) {
	tests := []struct {
		number   Number
		expected string
	}{
		{number: Number{1}, expected: "A"},
		{number: Number{2}, expected: "2"},
		{number: Number{10}, expected: "10"},
		{number: Number{11}, expected: "J"},
		{number: Number{12}, expected: "Q"},
		{number: Number{13}, expected: "K"},
	}
	for _, tt := range tests {
		number := tt.number
		value := strconv.Itoa(number.value)

		t.Run(value, func(t *testing.T) {
			// t.Fatalf でテストが失敗した場合でもクリーンアップ処理は呼び出される
			t.Cleanup(func() {
				t.Log("cleanup!")
			})

			if number.String() != tt.expected {
				t.Errorf("期待した値 %v 実際の値 %v", tt.expected, number.String())
			}
		})
	}
}

func TestNumberIsSame(t *testing.T) {
	tests := []struct {
		target   Number
		compare  Number
		expected bool
	}{
		{target: Number{1}, compare: Number{1}, expected: true},
		{target: Number{1}, compare: Number{2}, expected: false},
	}
	for _, tt := range tests {
		t.Run(tt.target.String(), func(t *testing.T) {
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
