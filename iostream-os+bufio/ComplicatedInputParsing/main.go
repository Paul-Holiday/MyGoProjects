package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Person struct {
	name string
	age  int
	city string
}

func NewPerson(name string, age int, city string) *Person {
	return &Person{name, age, city}
}

func GetPerson(p *Person) (name string, age int, city string) {
	return p.name, p.age, p.city
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите данные (Имя Возраст Город): ")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace("\n")
	data := strings.Split(input, " ")

	if len(data) != 3 {
		fmt.Println("Ошибка ввода!")
		return
	}

	age, err := strconv.Atoi(data[1])
	if err != nil {
		fmt.Println("Ошибка! Возраст должен быть числом!")
		return
	}

	person := NewPerson(data[0], age, data[2])
	fmt.Printf("%s, %d лет, %s\n", person.name, person.age, person.city)
}
