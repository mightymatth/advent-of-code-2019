package main

import (
	"github.com/mightymatth/advent-of-code-2019/tasks/day5/processor"
)

func main() {
	p := processor.Processor{}
	p.LoadMemory("tasks/day5/input.txt")

	p.Start()
}
