package main

import "fmt"

type Car struct {
	Name  string
	Cost  int
	Speed float64
}

func (s *Car) Acceleration() {
	s.Speed += 0.1
}

func main() {
	// i := 10
	//var p *int
	// var p *int = &i
	// fmt.Println(*p)
	// *p = 20
	// fmt.Println(i)

	Struct1 := Car{
		Name:  "Ford Focus",
		Cost:  4000,
		Speed: 60.5}
	Struct1.Acceleration()

	fmt.Println(Struct1.Speed)
}
