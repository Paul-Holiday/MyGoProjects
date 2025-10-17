package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

type Human struct {
	Name string
}

func (d Dog) Speak() string {
	return "Собака говорит: гав!"
}

func (c Cat) Speak() string {
	return "Кошка говорит: мяу!"
}

func (h Human) Speak() string {
	return "Человек говорит: привет!"
}

func main() {
	dog := Dog{"Шарик"}
	cat := Cat{"Борис"}
	human := Human{"Алексей"}

	speakers := []Speaker{dog, cat, human}

	for _, speaker := range speakers {
		fmt.Println(speaker.Speak())
	}
}
