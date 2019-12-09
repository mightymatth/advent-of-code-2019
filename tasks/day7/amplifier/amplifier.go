package amplifier

import (
	"github.com/mightymatth/advent-of-code-2019/tasks/day7/processor"
)

func Amplifier() int {
	p := processor.Processor{Input: make(chan int, 5), Output: make(chan int, 5)}
	pMemory := p.FileToMemoryData("tasks/day7/input.txt")

	inputsPerms := permutations([]int{0, 1, 2, 3, 4})

	maxSourceSignal := 0
	for _, phaseInputs := range inputsPerms {
		sourceSignal := 0
		for _, phaseInput := range phaseInputs {
			p.LoadMemoryWithData(pMemory)
			p.Input <- phaseInput
			p.Input <- sourceSignal
			p.Start()
			sourceSignal = <-p.Output
		}

		if sourceSignal > maxSourceSignal {
			maxSourceSignal = sourceSignal
		}
	}

	return maxSourceSignal
}

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	var res [][]int

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}
