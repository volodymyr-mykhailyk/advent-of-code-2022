package main

import (
	"github.com/volodymyr-mykhailyk/advent-of-code-2022/pkg/science"
	"github.com/volodymyr-mykhailyk/advent-of-code-2022/pkg/tasks"
	"github.com/volodymyr-mykhailyk/advent-of-code-2022/pkg/transformations"
	"sort"
)

func main() {
	tasks.Announce("Day01")
	rations := readData()
	tasks.Communicate("Calculating Rations for %v Elfs", len(rations))
	sort.Sort(sort.Reverse(sort.IntSlice(rations)))
	tasks.Communicate("Maximum ration across Elfs is: %v", rations[0])
	tasks.Communicate("Maximum ration across 3 top Elfs is: %v", science.SumElements(rations[0:3]))
}

func readData() []int {
	stringRations := transformations.GroupOver(tasks.ReadLines("cmd/day01/input.txt"), "")
	rations := make([]int, len(stringRations))
	for i, ration := range stringRations {
		rations[i] = science.SumElements(transformations.ParseIntegers(ration))
	}
	return rations
}
