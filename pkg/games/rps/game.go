package rps

var moveScores = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

var winningScores = map[string]int{
	"AX": 3,
	"AY": 6,
	"AZ": 0,
	"BX": 0,
	"BY": 3,
	"BZ": 6,
	"CX": 6,
	"CY": 0,
	"CZ": 3,
}

func RoundScore(opponent string, you string) int {
	return moveScore(you) + winningScore(opponent, you)
}

func winningScore(opponent string, you string) int {
	return winningScores[opponent+you]
}

func moveScore(you string) int {
	return moveScores[you]
}
