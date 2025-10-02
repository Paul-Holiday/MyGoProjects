package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== РАЗНЫЕ ВИДЫ ЗАДЕРЖЕК ===")

	// 1. Простой Sleep
	fmt.Println("1. Please wait 2 seconds!")
	time.Sleep(2 * time.Second)
	fmt.Println("   Time's up!")

	// 2. Sleep с переменной Duration
	duration := time.Second * 3
	fmt.Println("2. Please wait 3 more seconds!")
	time.Sleep(duration)
	fmt.Println("   Time's up!")

	// 3. time.After с каналом
	duration = time.Duration(2) * time.Second
	fmt.Println("3. Please wait 2 seconds (time.After)")
	<-time.After(duration)
	fmt.Println("   Time's up!")

	// 4. Таймер с ожиданием
	timer := time.NewTimer(2 * time.Second)
	fmt.Println("4. Wait 2 seconds (Timer)")
	<-timer.C
	fmt.Println("   Time's up!(timer)")

	// 5. Таймер с горутиной (без остановки)
	fmt.Println("5. Wait 2 seconds (Timer + goroutine)")
	timer = time.NewTimer(2 * time.Second)
	go func() {
		<-timer.C
		fmt.Println("   Time's up! (from goroutine)")
	}()
	time.Sleep(3 * time.Second) // Даем время горутине выполниться

	// 6. Таймер с остановкой (демонстрация)
	fmt.Println("6. Starting timer (will be stopped)")
	timer = time.NewTimer(5 * time.Second)
	go func() {
		<-timer.C
		fmt.Println("   This shouldn't print!")
	}()
	time.Sleep(1 * time.Second) // Ждем немного
	timer.Stop()
	fmt.Println("   Timer was stopped!")

	// 7. Интерактивный отсчет
	fmt.Println("\n7. Final countdown!(tootootoo-tooooo)")
	for i := 5; i > 0; i-- {
		fmt.Printf("   %d second(s) left!\n", i)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Time's up!")
}
