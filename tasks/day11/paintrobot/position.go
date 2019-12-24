package paintrobot

import "fmt"

type Position struct {
	X int
	Y int
}

func (p Position) toStr() string {
	return fmt.Sprintf("%v,%v", p.X, p.Y)
}
