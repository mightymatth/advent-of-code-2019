package processor

import (
	"fmt"
	"strconv"
)

type OpCode struct {
	ModeParam1 OpCodeMode
	ModeParam2 OpCodeMode
	ModeParam3 OpCodeMode
	Operation  OpCodeOperation
}

func NewOpCode(opCode int) OpCode {
	opCodeStr := strconv.Itoa(opCode)
	opCodeStr = fmt.Sprintf("%05s", opCodeStr)

	return OpCode{
		ModeParam1: getMode(opCodeStr[2]),
		ModeParam2: getMode(opCodeStr[1]),
		ModeParam3: getMode(opCodeStr[0]),
		Operation:  getOpCode(opCodeStr[3:5]),
	}
}

type OpCodeMode int

const (
	ImmediateMode OpCodeMode = iota
	PositionMode
	RelativeMode
)

type OpCodeOperation int

const (
	AddOperation OpCodeOperation = iota
	MultiplyOperation
	StoreInputOperation
	PrintInputOperation
	JumpIfTrueOperation
	JumpIfFalseOperation
	LessThanOperation
	EqualsOperation
	AdjustRelativeBaseOperation
	HaltOperation
)

func getMode(char uint8) OpCodeMode {
	switch char {
	case "0"[0]:
		return PositionMode
	case "1"[0]:
		return ImmediateMode
	case "2"[0]:
		return RelativeMode
	default:
		panic(fmt.Sprintf("Unknown operation code mode %v\n", char))
	}
}

func getOpCode(opCode string) OpCodeOperation {
	switch opCode {
	case "01":
		return AddOperation
	case "02":
		return MultiplyOperation
	case "03":
		return StoreInputOperation
	case "04":
		return PrintInputOperation
	case "05":
		return JumpIfTrueOperation
	case "06":
		return JumpIfFalseOperation
	case "07":
		return LessThanOperation
	case "08":
		return EqualsOperation
	case "09":
		return AdjustRelativeBaseOperation
	case "99":
		return HaltOperation
	default:
		panic(fmt.Sprintf("Unknown operation in operation code %v\n", opCode))
	}
}