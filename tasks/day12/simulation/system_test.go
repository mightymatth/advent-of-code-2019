package simulation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSystem_ApplyGravityAndVelocity(t *testing.T) {
	system := System{}
	system.AddObject(NewObject("Callisto", Position{X: 5, Y: -3, Z: 8}))
	system.AddObject(NewObject("Ganymede", Position{X: 3, Y: 12, Z: 8}))
	system.AddObject(NewObject("Europa", Position{X: 13, Y: 2, Z: 4}))
	system.ApplyGravity()

	assert.Equal(t, Velocity{X: 0, Y: 2, Z: -1}, system.Objects[0].Velocity)
	assert.Equal(t, Velocity{X: 2, Y: -2, Z: -1}, system.Objects[1].Velocity)
	assert.Equal(t, Velocity{X: -2, Y: 0, Z: 2}, system.Objects[2].Velocity)

	system.ApplyVelocity()

	assert.Equal(t, Position{X: 5, Y: -1, Z: 7}, system.Objects[0].Position)
	assert.Equal(t, Position{X: 5, Y: 10, Z: 7}, system.Objects[1].Position)
	assert.Equal(t, Position{X: 11, Y: 2, Z: 6}, system.Objects[2].Position)
}

func TestSystem_Simulate(t *testing.T) {
	system := System{}
	system.AddObject(NewObject("Io", Position{X: -1, Y: 0, Z: 2}))
	system.AddObject(NewObject("Europa", Position{X: 2, Y: -10, Z: -7}))
	system.AddObject(NewObject("Ganymede", Position{X: 4, Y: -8, Z: 8}))
	system.AddObject(NewObject("Callisto", Position{X: 3, Y: 5, Z: -1}))
	snapshot := system.Simulate(10)

	assert.Equal(t, 11, len(snapshot))

	assert.Equal(t, Position{X: 2, Y: 1, Z: -3}, snapshot[10].Objects[0].Position)
	assert.Equal(t, Position{X: 1, Y: -8, Z: 0}, snapshot[10].Objects[1].Position)
	assert.Equal(t, Position{X: 3, Y: -6, Z: 1}, snapshot[10].Objects[2].Position)
	assert.Equal(t, Position{X: 2, Y: 0, Z: 4}, snapshot[10].Objects[3].Position)

	assert.Equal(t, Velocity{X: -3, Y: -2, Z: 1}, snapshot[10].Objects[0].Velocity)
	assert.Equal(t, Velocity{X: -1, Y: 1, Z: 3}, snapshot[10].Objects[1].Velocity)
	assert.Equal(t, Velocity{X: 3, Y: 2, Z: -3}, snapshot[10].Objects[2].Velocity)
	assert.Equal(t, Velocity{X: 1, Y: -1, Z: -1}, snapshot[10].Objects[3].Velocity)
}

func TestEnergy(t *testing.T) {
	system := System{}
	system.AddObject(NewObject("Io", Position{X: -1, Y: 0, Z: 2}))
	system.AddObject(NewObject("Europa", Position{X: 2, Y: -10, Z: -7}))
	system.AddObject(NewObject("Ganymede", Position{X: 4, Y: -8, Z: 8}))
	system.AddObject(NewObject("Callisto", Position{X: 3, Y: 5, Z: -1}))
	snapshot := system.Simulate(10)

	assert.Equal(t, 11, len(snapshot))
	assert.Equal(t, 179, snapshot[10].Energy())
}

func TestRealInputPart1(t *testing.T) {
	system := System{}
	system.AddObject(NewObject("Io", Position{X: -14, Y: -4, Z: -11}))
	system.AddObject(NewObject("Europa", Position{X: -9, Y: 6, Z: -7}))
	system.AddObject(NewObject("Ganymede", Position{X: 4, Y: 1, Z: 4}))
	system.AddObject(NewObject("Callisto", Position{X: 2, Y: -14, Z: -9}))
	snapshot := system.Simulate(1000)

	assert.Equal(t, 10028, snapshot[10].Energy())
}

func TestSystem_SimulateCyclePart2Dummy1(t *testing.T) {
	system := System{}
	system.AddObject(NewObject("Io", Position{X: -1, Y: 0, Z: 2}))
	system.AddObject(NewObject("Europa", Position{X: 2, Y: -10, Z: -7}))
	system.AddObject(NewObject("Ganymede", Position{X: 4, Y: -8, Z: 8}))
	system.AddObject(NewObject("Callisto", Position{X: 3, Y: 5, Z: -1}))
	counter := system.SimulateCycleCount()

	assert.Equal(t, 2772, counter)

}

func TestSystem_SimulateCyclePart2Dummy2(t *testing.T) {
	system := System{}
	system.AddObject(NewObject("Io", Position{X: -8, Y: -10, Z: 0}))
	system.AddObject(NewObject("Europa", Position{X: 5, Y: 5, Z: 10}))
	system.AddObject(NewObject("Ganymede", Position{X: 2, Y: -7, Z: 3}))
	system.AddObject(NewObject("Callisto", Position{X: 9, Y: -8, Z: -3}))
	counter := system.SimulateCycleCount()

	assert.Equal(t, 4686774924, counter)

}

func TestSystem_SimulateCyclePart2(t *testing.T) {
	system := System{}
	system.AddObject(NewObject("Io", Position{X: -14, Y: -4, Z: -11}))
	system.AddObject(NewObject("Europa", Position{X: -9, Y: 6, Z: -7}))
	system.AddObject(NewObject("Ganymede", Position{X: 4, Y: 1, Z: 4}))
	system.AddObject(NewObject("Callisto", Position{X: 2, Y: -14, Z: -9}))
	counter := system.SimulateCycleCount()

	assert.Equal(t, 314610635824376, counter)
}
