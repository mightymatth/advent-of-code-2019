package processorv2

import (
	"fmt"
	"log"
)

type Instruction struct {
	index  int
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
	case JumpIfTrueOperation:
		return in.jumpIfTrue(), false
	case JumpIfFalseOperation:
		return in.jumpIfFalse(), false
	case LessThanOperation:
		return in.lessThan(), false
	case EqualsOperation:
		return in.equals(), false
	case HaltOperation:
		return 0, true
	default:
		log.Fatal("Unknown operation")
		return 0, true
	}
}

func (in *Instruction) add() (offset int) {
	val1 := in.memory[in.valueForParameter(First)]
	val2 := in.memory[in.valueForParameter(Second)]

	in.memory[in.valueForParameter(Third)] = val1 + val2

	return 4
}

func (in *Instruction) multiply() (offset int) {
	val1 := in.memory[in.valueForParameter(First)]
	val2 := in.memory[in.valueForParameter(Second)]

	in.memory[in.valueForParameter(Third)] = val1 * val2

	return 4
}

func (in *Instruction) storeInput() int {
	val1 := in.valueForParameter(First)

	in.memory[val1] = 5 // user input

	return 2
}

func (in *Instruction) printInput() int {
	val1 := in.memory[in.valueForParameter(First)]

	fmt.Println(val1)

	return 2
}

func (in *Instruction) jumpIfTrue() int {
	val1 := in.memory[in.valueForParameter(First)]
	val2 := in.memory[in.valueForParameter(Second)]

	if val1 != 0 {
		return in.jumpOffset(val2)
	}

	return 3
}

func (in *Instruction) jumpIfFalse() int {
	val1 := in.memory[in.valueForParameter(First)]
	val2 := in.memory[in.valueForParameter(Second)]

	if val1 == 0 {
		return in.jumpOffset(val2)
	}

	return 3
}

func (in *Instruction) lessThan() int {
	val1 := in.memory[in.valueForParameter(First)]
	val2 := in.memory[in.valueForParameter(Second)]
	val3 := in.valueForParameter(Third)

	if val1 < val2 {
		in.memory[val3] = 1
	} else {
		in.memory[val3] = 0
	}

	return 4
}

func (in *Instruction) equals() int {
	val1 := in.memory[in.valueForParameter(First)]
	val2 := in.memory[in.valueForParameter(Second)]
	val3 := in.valueForParameter(Third)

	if val1 == val2 {
		in.memory[val3] = 1
	} else {
		in.memory[val3] = 0
	}

	return 4
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
		return in.index + offset
	case PositionMode:
		return in.memory[in.index+offset]
	default:
		panic("Unknown operation code mode.")
	}
}

func (in *Instruction) jumpOffset(jumpIndex int) int {
	return jumpIndex - in.index
}

type ParamPosition int

const (
	First ParamPosition = iota
	Second
	Third
)
