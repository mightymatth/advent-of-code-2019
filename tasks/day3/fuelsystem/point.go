package fuelsystem

import (
	"math/rand"
	"regexp"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

func (p *Point) StringVal() string {
	return strconv.Itoa(p.X) + "," + strconv.Itoa(p.Y)
}

type WirePoint struct {
	Point     Point
	SerialNum int
	Step      int
}

func Path(from Point, pathDef string) []WirePoint {
	wireSerialNum := rand.Int()
	stepCounter := 0
	path := []WirePoint{{Point: from, SerialNum: wireSerialNum, Step: stepCounter}}

	moves := strings.Split(pathDef, ",")

	for _, move := range moves {
		re := regexp.MustCompile("^([LURD])([0-9]+)$")
		match := re.FindStringSubmatch(move)

		dir := match[1]
		steps, _ := strconv.Atoi(match[2])

		lastPathPoint := path[len(path)-1]

		for step := 0; step < steps; step++ {
			var newWirePoint WirePoint
			lastPoint := lastPathPoint.Point
			stepCounter++

			switch dir {
			case "L":
				newWirePoint = WirePoint{Point: Point{X: lastPoint.X - 1, Y: lastPoint.Y},
					SerialNum: wireSerialNum, Step: stepCounter}
			case "U":
				newWirePoint = WirePoint{Point: Point{X: lastPoint.X, Y: lastPoint.Y + 1},
					SerialNum: wireSerialNum, Step: stepCounter}
			case "R":
				newWirePoint = WirePoint{Point: Point{X: lastPoint.X + 1, Y: lastPoint.Y},
					SerialNum: wireSerialNum, Step: stepCounter}
			default:
				newWirePoint = WirePoint{Point: Point{X: lastPoint.X, Y: lastPoint.Y - 1},
					SerialNum: wireSerialNum, Step: stepCounter}
			}

			path = append(path, newWirePoint)
			lastPathPoint = newWirePoint
		}
	}

	return path
}
