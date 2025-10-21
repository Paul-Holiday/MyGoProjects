package main

import (
	"fmt"
	"time"
)

// Структуры

type Human struct {
	Name string
}

type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

type Lion struct {
	Name string
}

// Методы

func (h *Human) MakeSound() {
	fmt.Println("Человек говорит: Привет!")
}

func (d *Dog) MakeSound() {
	fmt.Println("Собака говорит: Гав!")
}

func (c *Cat) MakeSound() {
	fmt.Println("Кошка говорит: Мяу!")
}

func (l *Lion) MakeSound() {
	fmt.Println("Лев говорит: АРРРР!")
}

// Интерфейс

type Speaker interface {
	MakeSound()
}

func main() {
	h := &Human{"Артём"}
	d := &Dog{"Бобик"}
	c := &Cat{"Василиса"}
	l := &Lion{"Алекс"}

	speakers := []Speaker{h, d, c, l}

	for _, speaker := range speakers {
		speaker.MakeSound()
		time.Sleep(1 * time.Second)
	}
}
