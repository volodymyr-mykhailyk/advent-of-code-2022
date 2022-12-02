package rps

import "testing"

func TestRoundScore(t *testing.T) {
	tests := []struct {
		move Move
		want int
	}{
		{
			move: Move{opponent: "A", you: "Y"},
			want: 8,
		},
		{
			move: Move{opponent: "B", you: "X"},
			want: 1,
		},
		{
			move: Move{opponent: "C", you: "Z"},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.move.String(), func(t *testing.T) {
			if got := RoundScore(tt.move); got != tt.want {
				t.Errorf("RoundScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
