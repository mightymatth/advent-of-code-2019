package main

import (
	"fmt"
	"github.com/mightymatth/advent-of-code-2019/tasks/day6/orbit"
)

func main() {
	comObject := orbit.NewMap("tasks/day6/input.txt")
	fmt.Printf("Checksum: %v\n", comObject.CountOrbitSum())
}
