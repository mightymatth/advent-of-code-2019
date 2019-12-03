package main

import (
	"fmt"
	"github.com/mightymatth/advent-of-code-2019/tasks/day2"
	"math/rand"
)

func main() {
	p := day2.Processor{}
	memoryData := p.FileToMemoryData("tasks/day2/input.txt")
	var finalState day2.ProcessorOutput

	retries := 1
	for ; ; retries++ {
		p.LoadMemoryWithData(memoryData)
		randNoun := rand.Int() % 100
		randVerb := rand.Int() % 100

		p.SetNounVerb(randNoun, randVerb)

		state := p.Start()

		if state.Result == 19690720 {
			finalState = state
			break
		}
	}

	res := calculateResult(finalState)

	fmt.Printf("Result: %v\n", res)
	fmt.Printf("Tried %v times until calculated\n", retries)
}

func calculateResult(state day2.ProcessorOutput) int {
	return 100*state.Noun + state.Verb
}
