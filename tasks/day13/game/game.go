package game

import (
	"fmt"

	"github.com/mightymatth/advent-of-code-2019/tasks/day9/processor"
)

type Game struct {
	Screen    Tiles
	processor processor.Processor
}

func NewGame(filePath string) Game {
	proc := processor.NewProcessor()
	proc.LoadMemory(filePath)
	proc.AppendFreeMemory(1000)

	return Game{
		Screen:    make(Tiles),
		processor: proc,
	}
}

func (g *Game) Start() {
	go g.processor.Start(nil)
	g.drawScreen(g.processor.Output)
}

func (g *Game) drawScreen(procOutput <-chan int) {
	for {
		val1, opened := <-procOutput
		if !opened {
			break
		}

		val2, opened := <-procOutput
		if !opened {
			break
		}

		val3, opened := <-procOutput
		if !opened {
			break
		}

		position := Position{X: val1, Y: val2}
		tile := Tile(val3)

		g.Screen[position.toStr()] = tile
	}
}

type Tiles map[string]Tile

type Tile int

const (
	EmptyTile            Tile = 0
	WallTile             Tile = 1
	BlockTile            Tile = 2
	HorizontalPaddleTile Tile = 3
	BallTile             Tile = 4
)

type Position struct {
	X, Y int
}

func (p Position) toStr() string {
	return fmt.Sprintf("%v,%v", p.X, p.Y)
}
