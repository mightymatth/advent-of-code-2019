package main

import (
	"fmt"
	"github.com/mightymatth/advent-of-code-2019/tasks/day3/fuelsystem"
	"math"
)

func main() {
	wires := fuelsystem.FileToWires("tasks/day3/input.txt")

	wire1 := wires[0]
	wire2 := wires[1]

	crossPoints := fuelsystem.CrossPoints(wire1, wire2)

	var shortestStepSumCross fuelsystem.CrossPoint
	shortestStepSum := math.MaxInt32
	for _, crossPoint := range crossPoints {
		if crossPoint.StepSum < shortestStepSum {
			shortestStepSumCross = crossPoint
			shortestStepSum = crossPoint.StepSum
		}
	}

	fmt.Printf("Shortest step sum cross: %v\n", shortestStepSumCross)
}
