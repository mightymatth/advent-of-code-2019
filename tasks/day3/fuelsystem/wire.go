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
	return abs(p1.X-p2.X) + abs(p1.Y-p2.Y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func CrossPoints(wire1 Wire, wire2 Wire) []CrossPoint {
	crossPoints := make([]CrossPoint, 0)
	for _, point1 := range wire1.Points {
		for _, point2 := range wire2.Points {
			if point1 == point2 {
				crossPoints = append(crossPoints, NewCrossPoint(point1))
			}
		}
	}

	return crossPoints[1:]
}

type CrossPoint struct {
	Point Point
	Distance int
}

func NewCrossPoint(point Point) CrossPoint {
	return CrossPoint{
		Point:    point,
		Distance: Distance(Point{X: 0, Y: 0}, point),
	}
}
