package rps

import "strings"

var moveScores = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
}

var winningScores = map[string]int{
	"AA": 3,
	"AB": 6,
	"AC": 0,
	"BA": 0,
	"BB": 3,
	"BC": 6,
	"CA": 6,
	"CB": 0,
	"CC": 3,
}

var moveNormalization = map[string]string{
	"X": "A",
	"Y": "B",
	"Z": "C",
	"A": "A",
	"B": "B",
	"C": "C",
}

type Move struct {
	opponent string
	you      string
}

func (move Move) String() string {
	return move.opponent + " " + move.you
}

func ParseMove(line string) Move {
	moves := strings.Fields(line)
	return Move{opponent: moves[0], you: moves[1]}
}

func RoundScore(move Move) int {
	move = normalizeMove(move)
	return moveScore(move) + winningScore(move)
}

func normalizeMove(move Move) Move {
	return Move{opponent: move.opponent, you: moveNormalization[move.you]}
}

func winningScore(move Move) int {
	return winningScores[move.opponent+move.you]
}

func moveScore(move Move) int {
	return moveScores[move.you]
}
