package stationcalc

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Map struct {
	Width     int
	Height    int
	Asteroids []Asteroid
}

func NewMap(filePath string) Map {
	file, err := os.Open(filePath)

	if err != nil {
		panic("Cannot read the file!")
	}
	defer file.Close()

	asteroids := make([]Asteroid, 0)
	mapHeight := 0

	maxMapHeight := 0
	maxMapWidth := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		mapSymbols := strings.Split(strings.TrimSpace(scanner.Text()), "")

		for mapWidthIndex, symbol := range mapSymbols {
			switch MapSymbol(symbol) {
			case AsteroidSymbol:
				asteroids = append(asteroids, Asteroid{X: mapWidthIndex, Y: mapHeight})
			case SpaceSymbol:
			default:
				panic(fmt.Sprintf("Unknown map symbol: %v\n", symbol))
			}

			if mapWidthIndex+1 > maxMapWidth {
				maxMapWidth = mapWidthIndex + 1
			}
		}

		mapHeight++
		if mapHeight > maxMapHeight {
			maxMapHeight = mapHeight
		}
	}

	return Map{
		Width:     maxMapWidth,
		Height:    maxMapHeight,
		Asteroids: asteroids,
	}
}

func (m Map) Calc()  {

}

func (m *Map) CalcOptimalPosition() Asteroid {
	var optimalAsteroid Asteroid
	maxLOS := 0

	for i, asteroid := range m.Asteroids {
		m.Asteroids[i].LOS = m.CalcLOS(asteroid)

		if m.Asteroids[i].LOS > maxLOS {
			maxLOS = m.Asteroids[i].LOS
			optimalAsteroid = m.Asteroids[i]
		}
	}

	return optimalAsteroid
}

func (m Map) CalcLOS(ref Asteroid) int {
	shadeSpace := m.CalcShadeSpace(ref)
	los := 0

	for _, target := range m.Asteroids {
		if ref == target {
			continue
		}

		_, hit := shadeSpace[Position{X: target.X, Y: target.Y}.toString()]

		if !hit {
			los++
		}
	}

	return los
}

func (m Map) CalcShadeSpace(ref Asteroid) ShadeSpace {
	shadeSpace := make(ShadeSpace)
	for _, target := range m.Asteroids {
		if ref == target {
			continue
		}

		m.fillShadeSpace(ref, target, shadeSpace)
	}

	return shadeSpace
}

type ShadeSpace map[string]bool

func (m Map) fillShadeSpace(ref Asteroid, target Asteroid, shadeSpace ShadeSpace) {
	stepX, stepY := findStep(ref, target)

	for i := 1; ; i++ {
		nextPosition := Position{
			X: target.X + i*stepX,
			Y: target.Y + i*stepY,
		}

		if m.isOut(nextPosition) {
			break
		}

		shadeSpace[nextPosition.toString()] = true
	}
}

func findStep(ref Asteroid, target Asteroid) (stepX, stepY int) {
	x1 := ref.X
	y1 := ref.Y
	x2 := target.X
	y2 := target.Y

	stepX = (x2 - x1) / gcd(abs(x2-x1), abs(y2-y1))
	stepY = (y2 - y1) / gcd(abs(x2-x1), abs(y2-y1))

	return
}

// greatest common divisor (GCD) via Euclidean algorithm
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (m Map) isOut(p Position) bool {
	return p.X < 0 || p.X >= m.Width || p.Y < 0 || p.Y >= m.Height
}

type MapSymbol string

const (
	AsteroidSymbol MapSymbol = "#"
	SpaceSymbol              = "."
)
