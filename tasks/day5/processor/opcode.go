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
)

type OpCodeOperation int

const (
	AddOperation OpCodeOperation = iota
	MultiplyOperation
	StoreInputOperation
	PrintInputOperation
	HaltOperation
)

func getMode(char uint8) OpCodeMode {
	if char == "1"[0] {
		return ImmediateMode
	} else {
		return PositionMode
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
	case "99":
		return HaltOperation
	default:
		panic(fmt.Sprintf("Unknown operation in operation code %v\n", opCode))
	}
}
