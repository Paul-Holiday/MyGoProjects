// упражнение Exercise: Slices с платформы "A tour of Go"

package main

import (
	"golang.org/x/tour/pic"
	// "fmt"
)

func Pic(dx, dy int) [][]uint8 {
	matrix := make([][]uint8, dy)
	for i := range matrix {
		matrix[i] = make([]uint8, dx)
	}

	for x := 0; x < dy; x++ {
		for y := 0; y < dx; y++ {
			res := 1
			base := x
			for f := y; f > 0; {
				if f%2 == 1 {
					res *= base
				}
				base *= base
				f = f >> 1
			}
			matrix[x][y] = uint8(res)
			// fmt.Printf("x=%d y=%d x^y=%d\n", x, y, uint64(res)) // снять комментирование, чтобы убедиться, что возведение работает верно, но ведет себя неадекватно из-за переполнения переменной (пока не знаю что в таких случаях делать)
		}
	}
	return (matrix)
}

func main() {
	pic.Show(Pic)
}
