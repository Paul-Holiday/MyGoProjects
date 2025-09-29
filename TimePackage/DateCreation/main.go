package main

import (
	"fmt"
	"time"
)

func main() {
	// текущее время
	now := time.Now()

	// мой день рождения
	myBD := time.Date(2001, time.June, 24, 0, 0, 0, 0, time.Local)

	// день рождения моей жены
	myWifeBD := time.Date(2002, time.October, 7, 0, 0, 0, 0, time.Local)

	// день нашей свадьбы
	ourWD := time.Date(2023, time.August, 05, 0, 0, 0, 0, time.Local)

	// день космонавтики
	spaceDay := time.Date(1961, time.April, 12, 9, 7, 0, 0,
		time.FixedZone("MSK", 3*60*60))

	fmt.Printf("My bd was at the %v (%s) of %s %v!\n", myBD.Day(), myBD.Weekday(), myBD.Month(), myBD.Year())
	fmt.Printf("My wife's bd was at the %v (%s) of %s %v!\n", myWifeBD.Day(), myWifeBD.Weekday(), myWifeBD.Month(), myWifeBD.Year())
	fmt.Printf("Our wedding day was at the %v (%s) of %s %v!\n", ourWD.Day(), ourWD.Weekday(), ourWD.Month(), ourWD.Year())
	fmt.Printf("The space day was at the %v (%s) of %s %v!\n\n", spaceDay.Day(), spaceDay.Weekday(), spaceDay.Month(), spaceDay.Year())

	diff := now.Sub(myWifeBD)
	fmt.Printf("My wife is %v years old rn.\n", int(diff.Hours()/24/365.25))

	fmt.Printf("It is %v that I'm older than my wife.\n", myBD.Before(myWifeBD))
	fmt.Printf("Today's date is %s.\n", now.Format("02.01.2006"))
}
