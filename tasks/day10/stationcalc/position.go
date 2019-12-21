package stationcalc

import (
	"fmt"
	"math"
)

type Position struct {
	X int
	Y int
}

func (p Position) toString() string {
	return fmt.Sprintf("%v,%v", p.X, p.Y)
}

func measureAngle(p1 Position, p2 Position) float64 {
	y := float64(p2.Y - p1.Y)
	x := float64(p2.X - p1.X)

	return math.Mod((math.Atan2(y, x)-3*math.Pi/2+2*math.Pi+2*math.Pi)*(180/math.Pi), 360)
}
