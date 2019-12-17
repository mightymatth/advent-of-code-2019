package processor

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Processor struct {
	Memory []int
	Input  chan int
	Output chan int
}

func (p *Processor) Start() {
	i := 0
	for i < len(p.Memory) {
		instruction := NewInstruction(i, *p)
		offset, end := instruction.Execute()

		if end {
			break
		}

		i += offset
	}
}

func (p Processor) FileToMemoryData(filePath string) []int {
	file, err := os.Open(filePath)

	if err != nil {
		panic("Cannot read the file!")
	}
	defer file.Close()

	finalCodes := make([]int, 0)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		codes := strings.Split(strings.TrimSpace(scanner.Text()), ",")

		for _, codeStr := range codes {
			code, err := strconv.Atoi(codeStr)

			if err != nil {
				panic("Cannot convert text to integer!")
			}

			finalCodes = append(finalCodes, code)
		}
	}

	return finalCodes
}

func (p *Processor) GetMemory() []int {
	return p.Memory
}

func (p *Processor) LoadMemory(filePath string) {
	data := p.FileToMemoryData(filePath)

	p.Memory = data
}

func (p *Processor) LoadMemoryWithData(data []int) {
	dataCopy := make([]int, len(data))
	copy(dataCopy, data)

	p.Memory = dataCopy
}
