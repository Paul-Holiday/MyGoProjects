// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	// текущее время
// 	now := time.Now()

// 	fmt.Println("=== ТЕКУЩЕЕ ВРЕМЯ ===")
// 	fmt.Printf("Полная дата: %v\n\n", now)

// 	// отдельные компоненты
// 	fmt.Println("=== КОМПОНЕНТЫ ===")
// 	fmt.Printf("Дата: %v\n", now.Date()) // возвращает (year, month, day)
// 	fmt.Printf("Год: %d\n", now.Year())
// 	fmt.Printf("Месяц: %s (%d)\n", now.Month(), now.Month()) // строкой и числом
// 	fmt.Printf("День: %d\n", now.Day())
// 	fmt.Printf("Час: %d\n", now.Hour())
// 	fmt.Printf("Минуты: %d\n", now.Minute())
// 	fmt.Printf("Секунды: %d\n", now.Second())
// 	fmt.Printf("Наносекунды: %d\n", now.Nanosecond())
// 	fmt.Printf("Unix мс: %d\n", now.UnixMilli())
// 	fmt.Printf("День недели: %s\n\n", now.Weekday())

// 	// стандартные форматы
// 	fmt.Println("=== СТАНДАРТНЫЕ ФОРМАТЫ ===")
// 	fmt.Printf("RFC1123: %s\n", now.Format(time.RFC1123))
// 	fmt.Printf("RFC3339: %s\n", now.Format(time.RFC3339))
// 	fmt.Printf("RFC822: %s\n", now.Format(time.RFC822))
// 	fmt.Printf("DateOnly: %s\n", now.Format(time.DateOnly))
// 	fmt.Printf("TimeOnly: %s\n", now.Format(time.TimeOnly))
// 	fmt.Printf("Kitchen: %s\n\n", now.Format(time.Kitchen))

//		// кастомные форматы
//		fmt.Println("=== КАСТОМНЫЕ ФОРМАТЫ ===")
//		fmt.Printf("US: %s\n", now.Format("01/02/2006"))
//		fmt.Printf("EU: %s\n", now.Format("02.01.2006"))
//		fmt.Printf("Полный: %s\n", now.Format("January 2, 2006 15:04:05"))
//		fmt.Printf("Время: %s\n", now.Format("15:04:05"))
//		fmt.Printf("Короткое время: %s\n", now.Format("15:04"))
//		fmt.Printf("AM/PM: %s\n", now.Format("3:04:05 PM"))
//	}
package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println("=== ТЕКУЩЕЕ ВРЕМЯ ===")
	fmt.Printf("Полная дата: %v\n\n", now)
	year, month, day := now.Date()
	fmt.Println("=== КОМПОНЕНТЫ ===")
	fmt.Printf("Дата: %v - %d - %v\n", year, month, day) // возвращает (year, month, day)
	fmt.Printf("Год: %d\n", now.Year())
	fmt.Printf("Месяц: %s (%d)\n", now.Month(), now.Month()) // строкой и числом
	fmt.Printf("День: %d\n", now.Day())
	fmt.Printf("Час: %d\n", now.Hour())
	fmt.Printf("Минуты: %d\n", now.Minute())
	fmt.Printf("Секунды: %d\n", now.Second())
	fmt.Printf("Наносекунды: %d\n", now.Nanosecond())
	fmt.Printf("Unix мс: %d\n", now.UnixMilli())
	fmt.Printf("День недели: %s\n\n", now.Weekday())

	fmt.Println("=== СТАНДАРТНЫЕ ФОРМАТЫ ===")
	fmt.Printf("RFC1123: %s\n", now.Format(time.RFC1123))
	fmt.Printf("RFC3339: %s\n", now.Format(time.RFC3339))
	fmt.Printf("RFC822: %s\n", now.Format(time.RFC822))
	fmt.Printf("DateOnly: %s\n", now.Format(time.DateOnly))
	fmt.Printf("TimeOnly: %s\n", now.Format(time.TimeOnly))
	fmt.Printf("Kitchen: %s\n\n", now.Format(time.Kitchen))

	fmt.Println("=== КАСТОМНЫЕ ФОРМАТЫ ===")
	fmt.Printf("US: %s\n", now.Format("01/02/2006"))
	fmt.Printf("EU: %s\n", now.Format("02.01.2006"))
	fmt.Printf("Полный: %s\n", now.Format("January 2, 2006 15:04:05"))
	fmt.Printf("Время: %s\n", now.Format("15:04:05"))
	fmt.Printf("Короткое время: %s\n", now.Format("15:04"))
	fmt.Printf("AM/PM: %s\n", now.Format("3:04:05 PM"))
}
