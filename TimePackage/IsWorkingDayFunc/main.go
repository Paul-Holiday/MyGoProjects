package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	if IsWorkingDayFunc(now) {
		fmt.Printf("%s is a working day.\n", now.Weekday())
	} else {
		fmt.Printf("%s is a day off.\n", now.Weekday())
	}
}

// функция возвращает true/false в зависимости от того рабочий/выходной день
// получен на вход функции
func IsWorkingDayFunc(day time.Time) bool {
	weekday := day.Weekday()
	return weekday >= time.Monday && weekday <= time.Friday
}
