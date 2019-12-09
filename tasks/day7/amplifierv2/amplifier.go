package amplifierv2

import (
	"github.com/mightymatth/advent-of-code-2019/tasks/day7/processorv2"
	"sync"
)

func Amplifier() int {
	p := processorv2.Processor{Input: make(chan int, 50), Output: make(chan int, 50)}
	pMemory := p.FileToMemoryData("tasks/day7/input.txt")

	inputsPerms := permutations([]int{5, 6, 7, 8, 9})

	maxSourceSignal := 0
	for _, phaseInputs := range inputsPerms {
		sourceSignal := 0

		var wg sync.WaitGroup
		ampA := processorv2.Processor{Input: make(chan int, 50), Output: make(chan int, 50)}
		ampA.LoadMemoryWithData(pMemory)
		ampA.Input <- phaseInputs[0]
		ampA.Input <- sourceSignal
		wg.Add(1)
		go ampA.Start(&wg)

		ampB := processorv2.Processor{Input: make(chan int, 50), Output: make(chan int, 50)}
		ampB.LoadMemoryWithData(pMemory)
		ampB.Input <- phaseInputs[1]
		wg.Add(1)
		go ampB.Start(&wg)

		ampC := processorv2.Processor{Input: make(chan int, 50), Output: make(chan int, 50)}
		ampC.LoadMemoryWithData(pMemory)
		ampC.Input <- phaseInputs[2]
		wg.Add(1)
		go ampC.Start(&wg)

		ampD := processorv2.Processor{Input: make(chan int, 50), Output: make(chan int, 50)}
		ampD.LoadMemoryWithData(pMemory)
		ampD.Input <- phaseInputs[3]
		wg.Add(1)
		go ampD.Start(&wg)

		ampE := processorv2.Processor{Input: make(chan int, 50), Output: make(chan int, 50)}
		ampE.LoadMemoryWithData(pMemory)
		ampE.Input <- phaseInputs[4]
		wg.Add(1)
		go ampE.Start(&wg)

		go func() {
			for ; ; {
				select {
				case outputAmpA := <-ampA.Output:
					ampB.Input <- outputAmpA
				case outputAmpB := <-ampB.Output:
					ampC.Input <- outputAmpB
				case outputAmpC := <-ampC.Output:
					ampD.Input <- outputAmpC
				case outputAmpD := <-ampD.Output:
					ampE.Input <- outputAmpD
				case outputAmpE := <-ampE.Output:
					ampA.Input <- outputAmpE
				}
			}
		}()

		wg.Wait()

		sourceSignal = <-ampA.Input

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
