package fuelsystem

import (
	"bufio"
	"os"
	"strings"
)

type Wire struct {
	Points []Point
}

func NewWire(definition string) Wire {
	return Wire{Points: Path(Point{X: 0, Y: 0}, definition)}
}

func FileToWires(filePath string) []Wire {
	file, err := os.Open(filePath)

	if err != nil {
		panic("Cannot read the file!")
	}
	defer file.Close()

	wires := make([]Wire, 0)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		wireDef := strings.TrimSpace(scanner.Text())
		wires = append(wires, NewWire(wireDef))
	}

	return wires
}

func Distance(p1 Point, p2 Point) int {
	return abs(p1.X - p2.X) + abs(p1.Y - p2.Y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
