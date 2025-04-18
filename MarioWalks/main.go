package main

import "fmt"

type Mario struct {
	X, Y int
}

func (coor *Mario) MoveY(dir bool) {
	if dir {
		coor.Y++
	} else {
		coor.Y--
	}
}
func (coor *Mario) MoveX(dir bool) {
	if dir {
		coor.X++
	} else {
		coor.X--
	}
}

func main() {
	MarioDest := Mario{
		X: 4500,
		Y: 7000}

	MarioDest.MoveY(true)
	MarioDest.MoveX(false)

	fmt.Println(MarioDest)
}
