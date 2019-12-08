package main

import (
	"fmt"
	"github.com/mightymatth/advent-of-code-2019/tasks/day6/orbit"
)

func main() {
	space := orbit.NewSpace("tasks/day6/input.txt")
	fmt.Printf("Checksum: %v\n", space.CountOrbitSum())
	fmt.Printf("Diff orbital changes: %v\n", space.CalculateOrbitalChanges("YOU", "SAN"))
}
