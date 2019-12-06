package processor

import (
	"fmt"
	"log"
)

type Instruction struct {
	index int
	opCode OpCode
	memory []int
}

func NewInstruction(opCodeIndex int, memory []int) Instruction {
	return Instruction{index: opCodeIndex, opCode: NewOpCode(memory[opCodeIndex]), memory: memory}
}

func (in *Instruction) Execute() (offset int, halt bool) {
	switch in.opCode.Operation {
	case AddOperation:
		return in.add(), false
	case MultiplyOperation:
		return in.multiply(), false
	case StoreInputOperation:
		return in.storeInput(), false
	case PrintInputOperation:
		return in.printInput(), false
	case HaltOperation:
		return 0, true
	default:
		log.Fatal("Unknown operation")
		return 0, true
	}
}

func (in *Instruction) add() (offset int) {
	val1 := in.valueForParameter(First)
	val2 := in.valueForParameter(Second)
	val3 := in.memory[in.index + 3]

	in.memory[val3] = val1 + val2

	return 4
}

func (in *Instruction) multiply() (offset int) {
	val1 := in.valueForParameter(First)
	val2 := in.valueForParameter(Second)
	val3 := in.memory[in.index + 3]

	in.memory[val3] = val1 * val2

	return 4
}

func (in *Instruction) storeInput() int {
	val1 := in.memory[in.index + 1]

	in.memory[val1] = 1 // user input

	return 2
}

func (in *Instruction) printInput() int {
	val1 := in.memory[in.memory[in.index + 1]]

	fmt.Println(val1)

	return 2
}

func (in *Instruction) valueForParameter(position ParamPosition) int {
	switch position {
	case First:
		return in.valueForMode(1, in.opCode.ModeParam1)
	case Second:
		return in.valueForMode(2, in.opCode.ModeParam2)
	case Third:
		return in.valueForMode(3, in.opCode.ModeParam3)
	default:
		panic("Unknown parameter position.")
	}
}

func (in *Instruction) valueForMode(offset int, mode OpCodeMode) int {
	switch mode {
	case ImmediateMode:
		return in.memory[in.index + offset]
	case PositionMode:
		return in.memory[in.memory[in.index + offset]]
	default:
		panic("Unknown operation code mode.")
	}
}

type ParamPosition int
const (
	First ParamPosition = iota
	Second
	Third
)
