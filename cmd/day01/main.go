package main

import (
	"github.com/volodymyr-mykhailyk/advent-of-code-2022/pkg/science"
	"github.com/volodymyr-mykhailyk/advent-of-code-2022/pkg/tasks"
	"github.com/volodymyr-mykhailyk/advent-of-code-2022/pkg/transformations"
)

func main() {
	tasks.Announce("Day01")
	rations := transformations.GroupOver(tasks.ReadLines("cmd/day01/input.txt"), "")
	tasks.Communicate("Calculating Rations for %v Elfs\n", len(rations))
	maxCalories := 0
	for _, ration := range rations {
		currentCalories := science.SumElements(transformations.ParseIntegers(ration))
		if maxCalories < currentCalories {
			maxCalories = currentCalories
		}
	}
	tasks.Communicate("Maximum ration across Elfs is: %v", maxCalories)
}
