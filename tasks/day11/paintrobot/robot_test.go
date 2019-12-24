package paintrobot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcPaintedFields(t *testing.T) {
	robot := NewRobot("../input.txt")
	robot.StartPainting()
	paintedFieldCnt := robot.CalcPaintedFields()

	assert.Equal(t, 2238, paintedFieldCnt)
}
