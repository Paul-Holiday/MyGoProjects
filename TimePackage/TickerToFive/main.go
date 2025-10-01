package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	stop := make(chan bool)

	go func() {
		fmt.Scanln()
		stop <- true
	}()

	tick := 0

	fmt.Println("Ticker launched! Press Enter to.")
	for {
		select {
		case <-ticker.C:
			tick++
			fmt.Printf("\rTicker: %v.", tick)

			if tick >= 200 {
				fmt.Printf("\nTicker finished. %v ticks passed.", tick)
				return
			}
		case <-stop:
			fmt.Printf("\nTicker was stopped. %v ticks passed .", tick)
			return
		}

	}
}
