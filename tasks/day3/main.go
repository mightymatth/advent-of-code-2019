package main

import (
	"fmt"
	"github.com/mightymatth/advent-of-code-2019/tasks/day3/fuelsystem"
)

func main() {
	wires := fuelsystem.FileToWires("tasks/day3/input.txt")

	fmt.Println(wires)
}
