package stationcalc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMap(t *testing.T) {
	spaceMap := NewMap("../inputDummy1.txt")

	assert.Equal(t, 5, spaceMap.Width)
	assert.Equal(t, 5, spaceMap.Height)

	assert.Equal(t, 10, len(spaceMap.Asteroids))
	assert.Equal(t, NewAsteroid(1, 0), spaceMap.Asteroids[0])
	assert.Equal(t, NewAsteroid(4, 0), spaceMap.Asteroids[1])

	assert.Equal(t, NewAsteroid(0, 2), spaceMap.Asteroids[2])
	assert.Equal(t, NewAsteroid(1, 2), spaceMap.Asteroids[3])
	assert.Equal(t, NewAsteroid(2, 2), spaceMap.Asteroids[4])
	assert.Equal(t, NewAsteroid(3, 2), spaceMap.Asteroids[5])
	assert.Equal(t, NewAsteroid(4, 2), spaceMap.Asteroids[6])

	assert.Equal(t, NewAsteroid(4, 3), spaceMap.Asteroids[7])

	assert.Equal(t, NewAsteroid(3, 4), spaceMap.Asteroids[8])
	assert.Equal(t, NewAsteroid(4, 4), spaceMap.Asteroids[9])
}

func TestMap_CalcShadeSpace(t *testing.T) {
	spaceMap := NewMap("../inputDummy2.txt")

	// For Asteroid (1, 2)
	asteroidX1Y2 := spaceMap.Asteroids[0]
	shadeSpaceX1Y2 := spaceMap.CalcShadeSpace(asteroidX1Y2)
	assert.Equal(t, 2, len(shadeSpaceX1Y2))
	// With Target (1, 2) -> (2, 2)
	assert.True(t, shadeSpaceX1Y2[Position{X: 3, Y: 2}.toString()])
	assert.True(t, shadeSpaceX1Y2[Position{X: 4, Y: 2}.toString()])

	// For Asteroid (2, 2)
	asteroidX2Y2 := spaceMap.Asteroids[1]
	shadeSpaceX2Y2 := spaceMap.CalcShadeSpace(asteroidX2Y2)
	assert.Equal(t, 1, len(shadeSpaceX2Y2))
	// With Target (2, 2) -> (1, 2)
	assert.True(t, shadeSpaceX2Y2[Position{X: 0, Y: 2}.toString()])

	// For Asteroid (1, 4)
	asteroidX1Y4 := spaceMap.Asteroids[2]
	shadeSpaceX1Y4 := spaceMap.CalcShadeSpace(asteroidX1Y4)
	assert.Equal(t, 3, len(shadeSpaceX1Y4))
	// With Target (1, 4) -> (1, 2)
	assert.True(t, shadeSpaceX1Y4[Position{X: 1, Y: 1}.toString()])
	assert.True(t, shadeSpaceX1Y4[Position{X: 1, Y: 0}.toString()])
	// With Target (1, 4) -> (2, 2)
	assert.True(t, shadeSpaceX1Y4[Position{X: 3, Y: 0}.toString()])
}

func TestMap_CalcOptimalPosition(t *testing.T) {
	spaceMapDummy1 := NewMap("../inputDummy1.txt")
	bestAsteroidDummy1 := spaceMapDummy1.CalcOptimalPosition()
	assert.Equal(t, 8, bestAsteroidDummy1.LOS)

	spaceMap := NewMap("../input.txt")
	bestAsteroid := spaceMap.CalcOptimalPosition()
	assert.Equal(t, 260, bestAsteroid.LOS)
}

// Look at the picture: assets/measureAngleTest.jpg
func TestMeasureAngle(t *testing.T) {
	angle1 := measureAngle(Position{X: 7, Y: 9}, Position{X: 7, Y: 6})
	angle2 := measureAngle(Position{X: 7, Y: 9}, Position{X: 8, Y: 6})
	angle3 := measureAngle(Position{X: 7, Y: 9}, Position{X: 8, Y: 7})
	angle4 := measureAngle(Position{X: 7, Y: 9}, Position{X: 10, Y: 9})
	angle5 := measureAngle(Position{X: 7, Y: 9}, Position{X: 8, Y: 11})
	angle6 := measureAngle(Position{X: 7, Y: 9}, Position{X: 7, Y: 12})
	angle7 := measureAngle(Position{X: 7, Y: 9}, Position{X: 5, Y: 11})
	angle8 := measureAngle(Position{X: 7, Y: 9}, Position{X: 5, Y: 9})

	assert.True(t, angle2 > angle1)
	assert.True(t, angle3 > angle2)
	assert.True(t, angle4 > angle3)
	assert.True(t, angle5 > angle4)
	assert.True(t, angle6 > angle5)
	assert.True(t, angle7 > angle6)
	assert.True(t, angle8 > angle7)
}

func TestMap_LaserTargets(t *testing.T) {
	spaceMapDummy1 := NewMap("../inputDummy1.txt")
	laserDummy1 := spaceMapDummy1.CalcOptimalPosition()
	targetsDummy1 := spaceMapDummy1.LaserTargets(laserDummy1)

	assert.Equal(t, 9, len(targetsDummy1))
	assert.Equal(t, Position{X: 3, Y: 2}, targetsDummy1[0].Position)
	assert.Equal(t, Position{X: 4, Y: 0}, targetsDummy1[1].Position)
	assert.Equal(t, Position{X: 4, Y: 2}, targetsDummy1[2].Position)
	assert.Equal(t, Position{X: 4, Y: 3}, targetsDummy1[3].Position)
	assert.Equal(t, Position{X: 4, Y: 4}, targetsDummy1[4].Position)
	assert.Equal(t, Position{X: 0, Y: 2}, targetsDummy1[5].Position)
	assert.Equal(t, Position{X: 1, Y: 2}, targetsDummy1[6].Position)
	assert.Equal(t, Position{X: 2, Y: 2}, targetsDummy1[7].Position)
	assert.Equal(t, Position{X: 1, Y: 0}, targetsDummy1[8].Position)
}

func TestBet(t *testing.T) {
	assert.Equal(t, 802, calculateBet("../inputDummy3.txt"))
	assert.Equal(t, 608, calculateBet("../input.txt"))
}

func calculateBet(inputFilePath string) int {
	spaceMap := NewMap(inputFilePath)
	laser := spaceMap.CalcOptimalPosition()
	targets := spaceMap.LaserTargets(laser)
	aimedPosition := targets[199].Position
	result := aimedPosition.X*100 + aimedPosition.Y

	return result
}
