package main

import (
	"fmt"
	"time"
)

const EU = "02.01.2006 15:04:05"

func main() {
	dateString := "The date is 02.10.2025 09:58:05"
	layout := "The date is " + EU

	parsedTime, err := time.Parse(layout, dateString)
	// если есть ошибка выводим ошибку
	if err != nil {
		fmt.Println("Parsing failed. Error:", err)
		return
	}

	fmt.Println("Original string:", dateString)
	fmt.Println("Object out of string:", parsedTime.Format(EU))
}
