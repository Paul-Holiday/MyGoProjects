package main

import (
	"fmt"
	"time"
)

func main() {
	slice := []int{}
	fmt.Println("Cycle starts now!")
	beginning := time.Now()
	for i := 0; i < 500; i++ {
		slice = append(slice, i)
	}
	end := time.Now()
	duration := end.Sub(beginning)
	fmt.Printf("Cycle ran for %v seconds!", duration.Seconds())
}
