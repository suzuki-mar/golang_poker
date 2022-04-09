package poker

import (
	"testing"
)

func TestHandString(t *testing.T) {

	tests := []struct {
		hand     Hand
		expected string
	}{
		{hand: buildHand(buildHandParams()), expected: "♣2 ♣3 ♣4 ♣5 ♣6"},
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

func TestHandBuildRandHand(t *testing.T) {
	t.Run("buildBuildRandHand()", func(t *testing.T) {
		//作成できたらテストに成功したことにする
		BuildHand()
	})
}

func buildHandParams() [5]HandParam {
	var params [5]HandParam

	params[0] = HandParam{
		suit:        Clover,
		numberValue: 2,
	}
	params[1] = HandParam{
		suit:        Clover,
		numberValue: 3,
	}
	params[2] = HandParam{
		suit:        Clover,
		numberValue: 4,
	}
	params[3] = HandParam{
		suit:        Clover,
		numberValue: 5,
	}
	params[4] = HandParam{
		suit:        Clover,
		numberValue: 6,
	}

	return params
}
