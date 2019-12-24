package simulation

import (
	"fmt"

	"github.com/jinzhu/copier"
)

type System struct {
	Objects []Object
}

type Snapshot []System

func (s *System) AddObject(object Object) {
	s.Objects = append(s.Objects, object)
}

func (s *System) Simulate(iterations int) Snapshot {
	snapshot := make(Snapshot, 0)

	snapshot = append(snapshot, s.getSystemCopy())
	for i := 0; i < iterations; i++ {
		s.ApplyGravity()
		s.ApplyVelocity()
		snapshot = append(snapshot, s.getSystemCopy())
	}

	return snapshot
}

func (s *System) SimulateCycleCount() int {
	initStateX := s.checkSum(XAxis)
	initStateY := s.checkSum(YAxis)
	initStateZ := s.checkSum(ZAxis)

	var periodX, periodY, periodZ int
	var foundX, foundY, foundZ bool
	for i := 1; ; i++ {
		s.ApplyGravity()
		s.ApplyVelocity()

		if initStateX == s.checkSum(XAxis) && !foundX {
			periodX = i
			foundX = true
		}

		if initStateY == s.checkSum(YAxis) && !foundY {
			periodY = i
			foundY = true
		}

		if initStateZ == s.checkSum(ZAxis) && !foundZ {
			periodZ = i
			foundZ = true
		}

		if foundX && foundY && foundZ {
			break
		}
	}

	return lcm(periodX, lcm(periodY, periodZ))
}

// least common multiplier (LCM)
func lcm(a, b int) int {
	return abs(a*b) / gcd(a, b)
}

// greatest common divisor (GCD) via Euclidean algorithm
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func (s System) checkSum(axis SystemAxis) string {
	var data string

	for _, object := range s.Objects {
		switch axis {
		case XAxis:
			data += fmt.Sprintf("%v, %v|", object.Position.X, object.Velocity.X)
		case YAxis:
			data += fmt.Sprintf("%v, %v|", object.Position.Y, object.Velocity.Y)
		case ZAxis:
			data += fmt.Sprintf("%v, %v|", object.Position.Z, object.Velocity.Z)
		}
	}

	return data
}

type SystemAxis int

const (
	XAxis SystemAxis = iota
	YAxis
	ZAxis
)

// ApplyGravity applies gravity that changes objects' velocities
func (s *System) ApplyGravity() {
	newVelocities := make([]Velocity, 0, len(s.Objects))

	for _, ref := range s.Objects {
		var deltaX, deltaY, deltaZ int

		for _, target := range s.Objects {
			if ref == target {
				continue
			}

			switch {
			case ref.Position.X > target.Position.X:
				deltaX--
			case ref.Position.X < target.Position.X:
				deltaX++
			}

			switch {
			case ref.Position.Y > target.Position.Y:
				deltaY--
			case ref.Position.Y < target.Position.Y:
				deltaY++
			}

			switch {
			case ref.Position.Z > target.Position.Z:
				deltaZ--
			case ref.Position.Z < target.Position.Z:
				deltaZ++
			}
		}

		newVelocity := Velocity{
			X: ref.Velocity.X + deltaX,
			Y: ref.Velocity.Y + deltaY,
			Z: ref.Velocity.Z + deltaZ,
		}
		newVelocities = append(newVelocities, newVelocity)
	}

	for i := range s.Objects {
		s.Objects[i].Velocity = newVelocities[i]
	}
}

// ApplyVelocity applies velocity that changes objects' positions
func (s *System) ApplyVelocity() {
	for i := range s.Objects {
		s.Objects[i].Position.X += s.Objects[i].Velocity.X
		s.Objects[i].Position.Y += s.Objects[i].Velocity.Y
		s.Objects[i].Position.Z += s.Objects[i].Velocity.Z
	}
}

func (s System) getSystemCopy() System {
	var newSystem System
	err := copier.Copy(&newSystem, &s)

	if err != nil {
		panic(fmt.Sprintf("Cannot deep copy %v\n", s))
	}

	return newSystem
}

func (s System) Energy() (energy int) {
	for _, obj := range s.Objects {
		pot := abs(obj.Position.X) + abs(obj.Position.Y) + abs(obj.Position.Z)
		kin := abs(obj.Velocity.X) + abs(obj.Velocity.Y) + abs(obj.Velocity.Z)

		energy += pot * kin
	}

	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type Object struct {
	Name     string
	Position Position
	Velocity Velocity
}

func NewObject(name string, position Position) Object {
	return Object{
		Name:     name,
		Position: position,
		Velocity: Velocity{},
	}
}

type Position struct {
	X, Y, Z int
}

type Velocity struct {
	X, Y, Z int
}
