package fuelsystem

import (
	"regexp"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

type PathPoint struct {
	Point Point
	Step  int
}

func Path(from Point, pathDef string) []PathPoint {
	stepCounter := 0
	path := []PathPoint{{Point: from, Step: stepCounter}}

	moves := strings.Split(pathDef, ",")

	for _, move := range moves {
		re := regexp.MustCompile("^([LURD])([0-9]+)$")
		match := re.FindStringSubmatch(move)

		dir := match[1]
		steps, _ := strconv.Atoi(match[2])

		lastPathPoint := path[len(path)-1]

		for step := 0; step < steps; step++ {
			var newPathPoint PathPoint
			lastPoint := lastPathPoint.Point
			stepCounter++

			switch dir {
			case "L":
				newPathPoint = PathPoint{Point: Point{X: lastPoint.X - 1, Y: lastPoint.Y}, Step: stepCounter}
			case "U":
				newPathPoint = PathPoint{Point: Point{X: lastPoint.X, Y: lastPoint.Y + 1}, Step: stepCounter}
			case "R":
				newPathPoint = PathPoint{Point: Point{X: lastPoint.X + 1, Y: lastPoint.Y}, Step: stepCounter}
			default:
				newPathPoint = PathPoint{Point: Point{X: lastPoint.X, Y: lastPoint.Y - 1}, Step: stepCounter}
			}

			path = append(path, newPathPoint)
			lastPathPoint = newPathPoint
		}
	}

	return path
}
