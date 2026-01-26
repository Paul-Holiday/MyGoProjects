package main

import "fmt"

type Ordered interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64 | string
}

func Compare[T Ordered](a, b T) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	} else {
		return 0
	}
}

func main() {
	fmt.Println("Hello!")
}
