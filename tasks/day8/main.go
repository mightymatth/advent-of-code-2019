package main

import (
	"fmt"
	"github.com/mightymatth/advent-of-code-2019/tasks/day8/picture"
)

func main() {
	p := picture.LoadPicture("tasks/day8/input.txt", 25, 6)
	fmt.Println("Part 2:")
	p.Draw()
}
