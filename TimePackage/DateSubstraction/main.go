package main

import (
	"fmt"
	"time"
)

func main() {
	myBD := time.Date(2001, 06, 24, 0, 0, 0, 0, time.Local)
	now := time.Now()
	diff := now.Sub(myBD)

	days := int(diff.Hours()) / 24
	hours := int(diff.Hours()) % 24
	minutes := int(diff.Minutes()) % 60

	fmt.Printf("My birthday was %d days, %d hours, %d minutes past.", days, hours, minutes)
}
