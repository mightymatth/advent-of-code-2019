package paintrobot

import (
	"github.com/mightymatth/advent-of-code-2019/tasks/day9/processor"
)

type Robot struct {
	Proc           processor.Processor
	PaintingBlocks PaintingBlocks
	PaintHistory   chan PositionState
}

type PaintingBlocks map[string]PaintingBlock

type PaintingBlock struct {
	Position Position
	Color    PaintColor
}

func NewRobot(filePath string) Robot {
	proc := processor.NewProcessor()
	proc.LoadMemory(filePath)
	proc.AppendFreeMemory(1000)

	return Robot{
		Proc:           proc,
		PaintingBlocks: make(PaintingBlocks),
		PaintHistory:   make(chan PositionState),
	}
}

func (r *Robot) StartPainting() {
	go r.Proc.Start(nil)
	go r.painting()
}

// CalcPaintedFields prints how many fields are painted at least once.
func (r *Robot) CalcPaintedFields() int {
	paintedFieldMap := make(map[string]bool)

	for state := range r.PaintHistory {
		paintedFieldMap[state.Position.toStr()] = true
	}

	return len(paintedFieldMap)
}

func (r *Robot) painting() {
	currentPosition := Position{X: 0, Y: 0}
	currentFacing := FacingUp

	for {
		paintingBlock := r.getPaintingBlock(currentPosition)
		r.Proc.Input <- int(paintingBlock.Color)

		val1, opened := <-r.Proc.Output
		if !opened {
			break
		}

		val2, opened := <-r.Proc.Output
		if !opened {
			break
		}

		newPaintColor := PaintColor(val1)
		nextDirection := TurnDirection(val2)

		newPaintingBlock := r.paintBlock(currentPosition, newPaintColor)

		r.PaintHistory <- PositionState{
			Position: currentPosition,
			Facing:   currentFacing,
			OldColor: paintingBlock.Color,
			NewColor: newPaintingBlock.Color,
		}

		currentPosition, currentFacing = nextState(currentPosition, currentFacing, nextDirection)
	}

	close(r.PaintHistory)
}

func (r Robot) getPaintingBlock(position Position) PaintingBlock {
	paintingBlock, exists := r.PaintingBlocks[position.toStr()]

	if !exists {
		paintingBlock = PaintingBlock{
			Position: position,
			Color:    PaintBlack,
		}
	}

	return paintingBlock
}

func (r *Robot) paintBlock(position Position, color PaintColor) PaintingBlock {
	newPaintingBlock := PaintingBlock{
		Position: position,
		Color:    color,
	}

	r.PaintingBlocks[newPaintingBlock.Position.toStr()] = newPaintingBlock

	return newPaintingBlock
}

type PaintColor int

const (
	PaintBlack PaintColor = 0
	PaintWhite PaintColor = 1
)

//TurnDirection shows direction of rotation by 90 degrees.
type TurnDirection int

const (
	TurnLeft  TurnDirection = 0
	TurnRight TurnDirection = 1
)

type RobotFacing int

const (
	FacingUp RobotFacing = iota
	FacingDown
	FacingLeft
	FacingRight
)

func nextState(currPosition Position, currentFacing RobotFacing,
	moveDirection TurnDirection) (nextPosition Position, nextFacing RobotFacing) {

	nextFacing = calcNextFacing(currentFacing, moveDirection)
	switch nextFacing {
	case FacingUp:
		nextPosition = Position{X: currPosition.X, Y: currPosition.Y + 1}
	case FacingDown:
		nextPosition = Position{X: currPosition.X, Y: currPosition.Y - 1}
	case FacingLeft:
		nextPosition = Position{X: currPosition.X - 1, Y: currPosition.Y}
	case FacingRight:
		nextPosition = Position{X: currPosition.X + 1, Y: currPosition.Y}
	}

	return
}

func calcNextFacing(currentFacing RobotFacing, direction TurnDirection) (nextFacing RobotFacing) {
	switch currentFacing {
	case FacingUp:
		switch direction {
		case TurnLeft:
			nextFacing = FacingLeft
		case TurnRight:
			nextFacing = FacingRight
		}
	case FacingDown:
		switch direction {
		case TurnLeft:
			nextFacing = FacingRight
		case TurnRight:
			nextFacing = FacingLeft
		}
	case FacingRight:
		switch direction {
		case TurnLeft:
			nextFacing = FacingUp
		case TurnRight:
			nextFacing = FacingDown
		}
	case FacingLeft:
		switch direction {
		case TurnLeft:
			nextFacing = FacingDown
		case TurnRight:
			nextFacing = FacingUp
		}
	}

	return
}

type PositionState struct {
	Position Position
	Facing   RobotFacing
	OldColor PaintColor
	NewColor PaintColor
}
