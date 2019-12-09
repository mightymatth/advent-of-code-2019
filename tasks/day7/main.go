package main

import (
	"fmt"
	"github.com/mightymatth/advent-of-code-2019/tasks/day7/amplifier"
	"github.com/mightymatth/advent-of-code-2019/tasks/day7/amplifierv2"
)

func main() {
	fmt.Printf("Part 1; Amplifier result: %v\n", amplifier.Amplifier())
	fmt.Printf("Part 2; Amplifier result: %v\n", amplifierv2.Amplifier())
}
