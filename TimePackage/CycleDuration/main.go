package main

import (
	"fmt"
	"time"
)

func main() {
	// измерение производительности с помощью отложенной функции
	defer func(start time.Time) {
		fmt.Printf("All operations took %v.\n", time.Since(start))
	}(time.Now())

	// решение известными мне инструментами
	slice := []int{}
	fmt.Println("Cycle starts now!")
	beginning := time.Now()
	for i := 0; i < 50000; i++ {
		slice = append(slice, i)
	}
	end := time.Now()
	duration := end.Sub(beginning)
	fmt.Printf("Cycle ran for %v microseconds!\n\n", duration.Microseconds())

	// решение с помощью time.Since()
	slice1 := []int{}
	fmt.Println("Cycle 2 starts now!")
	start := time.Now()
	for i := range 50000 {
		slice1 = append(slice1, i)
	}
	fmt.Printf("Cycle 2 of ran for %v microseconds!\n\n", time.Since(start).Microseconds())
}
