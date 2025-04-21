// упражнение Exercise: Slices с платформы "A tour of Go"

package main

import (
	"fmt"

	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	matrix := make([][]uint8, dy)
	for i := range matrix {
		matrix[i] = make([]uint8, dx)
	}

	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			res := 1
			pow := 1
			for f := j; f > 0; {
				if f%2 == 1 {
					res *= pow
				}
				pow *= i
				f = f >> 1
			}
			matrix[i][j] = uint8(res)
			fmt.Printf("x=%d y=%d x^y=%d", i, j, res)
		}
	}
	return (matrix)
}

func main() {
	pic.Show(Pic)
}
