package main

import (
	"fmt"
	"time"
)

const EU = "02.01.2006 15:04:05"

func main() {
	now := time.Now()
	days := time.Hour * 3 * 24
	hours := time.Hour * 2

	fmt.Println("Current date: ", now.Format(EU))
	fmt.Println("Future date: ", now.Add(days+hours).Format(EU))
}
