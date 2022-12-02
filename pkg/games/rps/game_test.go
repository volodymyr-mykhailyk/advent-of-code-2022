package rps

import (
	"reflect"
	"testing"
)

type testInput = struct {
	move Move
	want int
}

func TestRoundScore(t *testing.T) {
	tests := []testInput{
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

func TestPredictCorrectMove(t *testing.T) {
	tests := []testInput{
		{
			move: Move{opponent: "A", you: "Y"},
			want: 4,
		},
		{
			move: Move{opponent: "B", you: "X"},
			want: 1,
		},
		{
			move: Move{opponent: "C", you: "Z"},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.move.String(), func(t *testing.T) {
			if got := RoundScore(PredictCorrectMove(tt.move)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PredictCorrectMove() = %v, want %v", got, tt.want)
			}
		})
	}
}
