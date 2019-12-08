package main

import (
	"fmt"
	"github.com/mightymatth/advent-of-code-2019/tasks/day3/fuelsystem"
)

func main() {
	wires := fuelsystem.FileToWires("tasks/day3/input.txt")

	wire1 := wires[0]
	wire2 := wires[1]

	crossPoints := fuelsystem.CrossPoints(wire1, wire2)
	shortestStepSumCross := fuelsystem.ShortestCrossPointForPathSum(crossPoints)

	fmt.Printf("Shortest step sum cross: %v\n", shortestStepSumCross)
	fmt.Printf("Shortest step: %v\n", shortestStepSumCross.StepSum)
}
