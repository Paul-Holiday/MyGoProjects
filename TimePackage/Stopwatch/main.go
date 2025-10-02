package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	seconds := 0
	fmt.Printf("Stopwatch started!\n")

	for range ticker.C {
		seconds++
		fmt.Printf("\rStopwatch: %v seconds passed.", time.Duration(seconds)*time.Second)
	}
}
