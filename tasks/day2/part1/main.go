package main

import (
	"fmt"
	"github.com/mightymatth/advent-of-code-2019/tasks/day2"
)

func main() {
	p := day2.Processor{}
	p.LoadMemory("tasks/day2/input.txt")

	p.SetNounVerb(12, 2)

	res := p.Start()

	fmt.Println(res)
}
