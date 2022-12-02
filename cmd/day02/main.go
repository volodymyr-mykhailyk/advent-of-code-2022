package main

import (
	"github.com/volodymyr-mykhailyk/advent-of-code-2022/pkg/games/rps"
	"github.com/volodymyr-mykhailyk/advent-of-code-go-tools/pkg/tasks"
	"github.com/volodymyr-mykhailyk/advent-of-code-go-tools/pkg/transformations"
)

func main() {
	tasks.Announce("Day02")
	moves := readData()
	tasks.Communicate("Predicting Rock Paper Scissors Game with %v moves", len(moves))
	score := 0
	for _, move := range moves {
		score += rps.RoundScore(move[0], move[1])
	}
	tasks.Communicate("Predicted score for the game is: %v", score)
}

func readData() [][]string {
	return transformations.SplitBySpace(tasks.ReadLines("cmd/day02/input.txt"))
}
