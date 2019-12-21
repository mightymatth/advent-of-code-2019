package stationcalc

type Asteroid struct {
	Position Position
	LOS int
}

func NewAsteroid(x, y int) Asteroid {
	return Asteroid{
		Position: Position{X: x, Y: y},
		LOS:      0,
	}
}
