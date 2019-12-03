package day2

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Processor struct {
	memory []int
}

func (p Processor) Start() ProcessorOutput {
	for i := 0; i < len(p.memory); i += 4 {
		pin := Instruction{operatorIndex: i, memory: p.memory}
		p.processOperation(pin)
	}

	return ProcessorOutput{
		Result: p.memory[0],
		Noun:   p.memory[1],
		Verb:   p.memory[2],
	}
}

func (p Processor) processOperation(instr Instruction) {
	operate := p.getOperator(instr.memory[instr.operatorIndex])

	if operate == nil {
		return
	}

	oprA := instr.memory[instr.memory[instr.operatorIndex+1]]
	oprB := instr.memory[instr.memory[instr.operatorIndex+2]]
	destIndex := instr.memory[instr.operatorIndex+3]

	resultCode := operate(oprA, oprB)

	instr.memory[destIndex] = resultCode
}

func (p Processor) getOperator(operatorCode int) func(val1 int, val2 int) int {
	switch operatorCode {
	case 1:
		return p.add
	case 2:
		return p.multiply
	case 99:
		return nil
	default:
		return nil
	}
}

func (p *Processor) SetNounVerb(noun int, verb int) {
	p.memory[1] = noun
	p.memory[2] = verb
}

func (p Processor) add(val1 int, val2 int) int {
	return val1 + val2
}

func (p Processor) multiply(val1 int, val2 int) int {
	return val1 * val2
}

type Instruction struct {
	operatorIndex int
	memory        []int
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

func (p *Processor) LoadMemory(filePath string) {
	data := p.FileToMemoryData(filePath)

	p.memory = data
}

func (p *Processor) LoadMemoryWithData(data []int) {
	dataCopy := make([]int, len(data))
	copy(dataCopy, data)

	p.memory = dataCopy
}

type ProcessorOutput struct {
	Result int
	Noun   int
	Verb   int
}
