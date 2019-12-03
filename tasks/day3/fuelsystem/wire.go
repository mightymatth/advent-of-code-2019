package fuelsystem

import (
	"bufio"
	"os"
	"strings"
)

type Wire struct {
	PathPoints []PathPoint
}

func NewWire(definition string) Wire {
	return Wire{PathPoints: Path(Point{X: 0, Y: 0}, definition)}
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
	for _, pathPoint1 := range wire1.PathPoints {
		for _, pathPoint2 := range wire2.PathPoints {
			if pathPoint1.Point == pathPoint2.Point {
				crossPoints = append(crossPoints, NewCrossPoint(pathPoint1, pathPoint2))
			}
		}
	}

	return crossPoints[1:]
}

type CrossPoint struct {
	Point    Point
	Distance int
	StepSum  int
}

func NewCrossPoint(pathP1 PathPoint, pathP2 PathPoint) CrossPoint {
	return CrossPoint{
		Point:    pathP1.Point,
		Distance: Distance(Point{X: 0, Y: 0}, pathP1.Point),
		StepSum:  pathP1.Step + pathP2.Step,
	}
}
