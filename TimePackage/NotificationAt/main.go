package main

import (
	"fmt"
	"time"
)

func main() {
	// создаю тикер, отложенный .Stop для него и канал stop для остановки бесконечного цикла
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	stop := make(chan bool)

	layout := "2.1.2006 15:04:05"
	notifTime := "2.10.2025 17:57:00"

	notifText := "Message!"

	parsedTime, err := time.ParseInLocation(layout, notifTime, time.Local)
	if err != nil {
		fmt.Println("Parsing failed. Error: ", err)
		return
	}

	// горутина ожидающая Enter чтобы остановить программу
	go func() {
		fmt.Scanln()
		stop <- true
	}()

	fmt.Println("Notification set for:", parsedTime.Format(layout), "Press Enter to cancel!")

	for {
		select {
		case <-ticker.C:
			now := time.Now()
			fmt.Printf("\rActual local time is %v", now.Truncate(time.Second))

			if parsedTime.Truncate(time.Second).Equal(now.Truncate(time.Second)) {

				fmt.Printf("\nNotification: %s \nNotification time: %v", notifText, parsedTime.Format(layout))
				return
			}
		case <-stop:
			fmt.Println("Notification cancelled.")
			return
		}
	}
}
