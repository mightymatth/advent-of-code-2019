package processor

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Processor struct {
	Memory       []int
	RelativeBase int
	Input        chan int
	Output       chan int
}

func NewProcessor() Processor {
	return Processor{
		Memory:       nil,
		RelativeBase: 0,
		Input:        make(chan int, 300),
		Output:       make(chan int, 300),
	}
}

func (p *Processor) Start(wg *sync.WaitGroup) {
	i := 0
	for i < len(p.Memory) {
		instruction := NewInstruction(i, p)
		offset, end := instruction.Execute()

		if end {
			if wg != nil {
				wg.Done()
			}
			close(p.Input)
			close(p.Output)
			break
		}

		i += offset
	}
}

func FileToMemory(filePath string) []int {
	file, err := os.Open(filePath)

	if err != nil {
		panic("Cannot read the file!")
	}
	defer file.Close()

	finalCodes := make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		newCodes := parseLine(strings.TrimSpace(scanner.Text()))
		finalCodes = append(finalCodes, newCodes...)
	}

	return finalCodes
}

func parseLine(codes string) []int {
	codesStr := strings.Split(codes, ",")
	var newCodes []int

	for _, codeStr := range codesStr {
		code, err := strconv.Atoi(codeStr)

		if err != nil {
			panic("Cannot convert text to integer!")
		}

		newCodes = append(newCodes, code)
	}

	return newCodes
}

func (p *Processor) LineToMemory(codes string) {
	codesForMemory := parseLine(codes)
	p.Memory = codesForMemory
}

func (p *Processor) GetMemory() []int {
	return p.Memory
}

func (p *Processor) LoadMemory(filePath string) {
	data := FileToMemory(filePath)

	p.Memory = data
}

func (p *Processor) LoadMemoryWithData(data []int) {
	dataCopy := make([]int, len(data))
	copy(dataCopy, data)

	p.Memory = dataCopy
}

func (p *Processor) AppendFreeMemory(size int) {
	newMemory := make([]int, size)
	p.Memory = append(p.Memory, newMemory...)
}
