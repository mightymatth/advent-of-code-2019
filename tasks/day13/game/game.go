package game

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/mightymatth/advent-of-code-2019/tasks/day9/processor"
)

type Game struct {
	Screen    Tiles
	processor processor.Processor
	score     int
}

func NewGame(filePath string) Game {
	proc := processor.NewProcessor()
	proc.LoadMemory(filePath)
	proc.AppendFreeMemory(1000)

	return Game{
		Screen:    Tiles{m: make(map[Position]Tile)},
		processor: proc,
	}
}

func (g *Game) StartFreeMode() {
	g.processor.Memory[0] = 2

	go g.processor.Start(nil)

	refreshScreen := make(chan bool)
	go g.createScreenElements(g.processor.Output, refreshScreen)
	go g.drawScreen(refreshScreen)
	g.computerPlayer()
}

func (g *Game) Start() {
	go g.processor.Start(nil)
	g.createScreenElements(g.processor.Output, nil)
}

func (g *Game) createScreenElements(procOutput <-chan int, refreshScreen chan<- bool) {
	scorePosition := Position{X: -1, Y: 0}
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

		if position == scorePosition {
			g.score = val3
		} else {
			tile := Tile(val3)
			g.Screen.Lock()
			g.Screen.m[position] = tile
			g.Screen.Unlock()
		}

		if refreshScreen != nil {
			refreshScreen <- true
		}
	}
}

func (g *Game) drawScreen(refreshScreen <-chan bool) {
	for range refreshScreen {
		g.Screen.Lock()
		var width, height int
		for position := range g.Screen.m {
			if position.X > width {
				width++
			}
			if position.Y > height {
				height++
			}
		}

		width++
		height++

		fmt.Printf("Score: %v\n", g.score)
		for j := 0; j < height; j++ {
			for i := 0; i < width; i++ {
				tile, found := g.Screen.m[Position{X: i, Y: j}]

				if found {
					switch tile {
					case EmptyTile:
						fmt.Print(" ")
					case WallTile:
						fmt.Print("█")
					case BlockTile:
						fmt.Print("░")
					case HorizontalPaddleTile:
						fmt.Print("▒")
					case BallTile:
						fmt.Print("●")
					}
				} else {
					fmt.Print("X")
				}
			}
			fmt.Println("")
		}
		fmt.Println("")
		g.Screen.Unlock()
	}
}

func (g *Game) computerPlayer() {
	fmt.Println("Press Enter to start!")
	reader := bufio.NewReader(os.Stdin)
	_, _, err := reader.ReadRune()

	if err != nil {
		fmt.Println(err)
	}
	ticker := time.NewTicker(200 * time.Millisecond)

	var lastBallPosition Position
	for {
		<-ticker.C

		var ballPos Position
		var ballFound bool
		var paddlePos Position
		var paddleFound bool

		g.Screen.Lock()
		for position, tile := range g.Screen.m {
			if ballFound && paddleFound {
				return
			}
			if tile == BallTile && !ballFound {
				ballPos = position
			}

			if tile == HorizontalPaddleTile && !paddleFound {
				paddlePos = position
			}
		}
		g.Screen.Unlock()

		if lastBallPosition.X == 0 && lastBallPosition.Y == 0 {
			g.processor.Input <- 0
		} else {
			ballDirection := getBallDirection(lastBallPosition, ballPos)

			switch ballDirection {
			case LeftDirection:
				if paddlePos.X < ballPos.X {
					if paddlePos.X+1 == ballPos.X {
						if paddlePos.Y-ballPos.Y == 1 {
							g.processor.Input <- 1
						} else {
							g.processor.Input <- -1
						}
					} else {
						g.processor.Input <- 0
					}
				} else if paddlePos.X == ballPos.X {
					g.processor.Input <- -1
				} else {
					g.processor.Input <- -1
				}
			case RightDirection:
				if paddlePos.X > ballPos.X {
					if paddlePos.X == ballPos.X+1 {
						if paddlePos.Y-ballPos.Y == 1 {
							g.processor.Input <- -1
						} else {
							g.processor.Input <- 1
						}
					} else {
						g.processor.Input <- 1
					}
				} else if paddlePos.X == ballPos.X {
					g.processor.Input <- 1
				} else {
					g.processor.Input <- 1
				}
			}
		}

		lastBallPosition = ballPos
	}
}

type Tiles struct {
	sync.RWMutex
	m map[Position]Tile
}

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

func getBallDirection(last, current Position) BallDirection {
	if current.X > last.X {
		return RightDirection
	} else {
		return LeftDirection
	}
}

type BallDirection int

const (
	LeftDirection BallDirection = iota
	RightDirection
)
