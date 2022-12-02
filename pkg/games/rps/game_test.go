package rps

import "testing"

func TestRoundScore(t *testing.T) {
	type args struct {
		opponent string
		you      string
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{opponent: "A", you: "Y"},
			want: 8,
		},
		{
			args: args{opponent: "B", you: "X"},
			want: 1,
		},
		{
			args: args{opponent: "C", you: "Z"},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.args.opponent+tt.args.you, func(t *testing.T) {
			if got := RoundScore(tt.args.opponent, tt.args.you); got != tt.want {
				t.Errorf("RoundScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
