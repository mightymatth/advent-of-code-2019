package paintrobot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcPaintedFields(t *testing.T) {
	robot := NewRobot("../input.txt")
	go robot.Paint()
	paintedFieldCnt := robot.CalcPaintedFields()

	assert.Equal(t, 2238, paintedFieldCnt)
}

func TestPrint(t *testing.T) {
	robot := NewRobot("../input.txt")
	robot.PaintBlock(Position{X: 0, Y: 0}, PaintWhite)
	robot.Paint()
	PaintBlocks(robot.PaintingBlocks) // PKFPAZRP
}
