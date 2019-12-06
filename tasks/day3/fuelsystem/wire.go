package fuelsystem

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
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
	defer func(timeStart time.Time) { fmt.Printf("Time in CrossPoints spent: %v\n", time.Since(timeStart)) }(time.Now())

	crossPoints := make([]CrossPoint, 0)
	crossPointsMap := make(map[string][]PathPoint)
	wirePoints := append(wire1.PathPoints, wire2.PathPoints...)

	for _, wirePoint := range wirePoints {
		val, exists := crossPointsMap[wirePoint.Point.StringVal()]
		if exists {
			val = append(val, wirePoint)
			crossPointsMap[wirePoint.Point.StringVal()] = val
		} else {
			crossPointsMap[wirePoint.Point.StringVal()] = []PathPoint{wirePoint}
		}

		val = crossPointsMap[wirePoint.Point.StringVal()]
	}

	startPoint := Point{X: 0, Y: 0}
	delete(crossPointsMap, startPoint.StringVal())

	for _, mapCrossPoints := range crossPointsMap {
		if len(mapCrossPoints) == 2 {
			a := NewCrossPoint(mapCrossPoints[0], mapCrossPoints[1])
			crossPoints = append(crossPoints, a)
		}
	}

	return crossPoints
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
