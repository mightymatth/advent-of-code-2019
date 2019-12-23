package fuelsystem

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Wire struct {
	WirePoints []WirePoint
}

func NewWire(definition string) Wire {
	return Wire{WirePoints: Path(Point{X: 0, Y: 0}, definition)}
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

func CrossPoints(wire1 Wire, wire2 Wire) []CrossPoint {
	defer func(timeStart time.Time) { fmt.Printf("Time in CrossPoints spent: %v\n", time.Since(timeStart)) }(time.Now())

	crossPoints := make([]CrossPoint, 0)
	crossPointsMap := make(map[string][]WirePoint)
	wirePoints := append(wire1.WirePoints, wire2.WirePoints...)

	for _, wirePoint := range wirePoints {
		val, exists := crossPointsMap[wirePoint.Point.StringVal()]
		if exists {
			val = append(val, wirePoint)
			crossPointsMap[wirePoint.Point.StringVal()] = val
		} else {
			crossPointsMap[wirePoint.Point.StringVal()] = []WirePoint{wirePoint}
		}
	}

	startPoint := Point{X: 0, Y: 0}
	delete(crossPointsMap, startPoint.StringVal())

	for _, mapCrossPoints := range crossPointsMap {
		if len(mapCrossPoints) < 2 {
			continue
		}

		for i, mapCrossPoint1 := range mapCrossPoints {
			for j, mapCrossPoint2 := range mapCrossPoints {
				if i == j {
					continue
				}

				crossPoints = append(crossPoints, NewCrossPoint(mapCrossPoint1, mapCrossPoint2))
			}
		}
	}

	return crossPoints
}
