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

func Path(from Point, pathDef string) []Point {
	path := []Point{from}

	moves := strings.Split(pathDef, ",")

	for _, move := range moves {
		re := regexp.MustCompile("^([LURD])([0-9]+)$")
		match := re.FindStringSubmatch(move)

		dir := match[1]
		steps, _ := strconv.Atoi(match[2])

		lastPoint := path[len(path)-1]

		for step := 0 ; step < steps; step++ {
			var newPoint Point

			switch dir {
			case "L":
				newPoint = Point{X:lastPoint.X-1, Y:lastPoint.Y}
			case "U":
				newPoint = Point{X:lastPoint.X, Y:lastPoint.Y+1}
			case "R":
				newPoint = Point{X:lastPoint.X+1, Y:lastPoint.Y}
			default:
				newPoint = Point{X:lastPoint.X, Y:lastPoint.Y-1}
			}

			path = append(path, newPoint)
			lastPoint = newPoint
		}
	}

	return path
}

