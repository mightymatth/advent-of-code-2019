package fuelsystem

import (
	"math"
)

type CrossPoint struct {
	Point          Point
	Wire1SerialNum int
	Wire2SerialNum int
	Distance       int
	StepSum        int
}

func NewCrossPoint(wireP1 WirePoint, wireP2 WirePoint) CrossPoint {
	return CrossPoint{
		Point:          wireP1.Point,
		Wire1SerialNum: wireP1.SerialNum,
		Wire2SerialNum: wireP2.SerialNum,
		Distance:       Distance(Point{X: 0, Y: 0}, wireP1.Point),
		StepSum:        wireP1.Step + wireP2.Step,
	}
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

func (cp CrossPoint) DifferentWires() bool {
	return cp.Wire1SerialNum != cp.Wire2SerialNum
}

func ShortestCrossPointForDistance(crossPoints []CrossPoint) CrossPoint {
	var shortestCross CrossPoint
	shortestCrossDistance := math.MaxInt32
	for _, crossPoint := range crossPoints {
		if crossPoint.Distance < shortestCrossDistance && crossPoint.DifferentWires() {
			shortestCross = crossPoint
			shortestCrossDistance = crossPoint.Distance
		}
	}

	return shortestCross
}

func ShortestCrossPointForPathSum(crossPoints []CrossPoint) CrossPoint {
	var shortestCross CrossPoint
	shortestStepSum := math.MaxInt32
	for _, crossPoint := range crossPoints {
		if crossPoint.StepSum < shortestStepSum && crossPoint.DifferentWires() {
			shortestCross = crossPoint
			shortestStepSum = crossPoint.StepSum
		}
	}

	return shortestCross
}
