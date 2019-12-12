package main

import (
	"fmt"
	"github.com/mightymatth/advent-of-code-2019/tasks/day8/picture"
)

func main() {
	p := picture.LoadPicture("tasks/day8/inputDummy.txt", 3, 2)

	fmt.Println(p.Layers)
}
