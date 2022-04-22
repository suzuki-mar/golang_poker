package judge

import "testing"

func TestTwoPair_Judge(t *testing.T) {
	type args struct {
		cards [5]hand.Card
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hc := TwoPair{}
			if got := hc.Judge(tt.args.cards); got != tt.want {
				t.Errorf("Judge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTwoPair_Rank(t *testing.T) {
	tests := []struct {
		name string
		want PokerHand
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hc := TwoPair{}
			if got := hc.Rank(); got != tt.want {
				t.Errorf("Rank() = %v, want %v", got, tt.want)
			}
		})
	}
}
