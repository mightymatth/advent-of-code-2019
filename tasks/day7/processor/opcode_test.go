package processor

import "testing"

func TestNewOpCode(t *testing.T) {
	test1OpCode := OpCode{
		ModeParam1: ImmediateMode,
		ModeParam2: PositionMode,
		ModeParam3: ImmediateMode,
		Operation:  MultiplyOperation,
	}

	if NewOpCode(10102) != test1OpCode {
		t.Fail()
	}
}
