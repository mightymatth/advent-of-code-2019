package main

import (
	"github.com/mightymatth/advent-of-code-2019/tasks/day5/processorv2"
)

func main() {
	p := processorv2.Processor{}
	p.LoadMemory("tasks/day5/input.txt")

	p.Start()
}
