package processorv2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Processor struct {
	memory []int
}

func (p *Processor) Start() {
	i := 0
	for ; i < len(p.memory); {
		instruction := NewInstruction(i, p.memory)
		offset, end := instruction.Execute()

		if end {
			fmt.Println("Program halts!")
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
	return p.memory
}

func (p *Processor) LoadMemory(filePath string) {
	data := p.FileToMemoryData(filePath)

	p.memory = data
}

func (p *Processor) LoadMemoryWithData(data []int) {
	dataCopy := make([]int, len(data))
	copy(dataCopy, data)

	p.memory = dataCopy
}
