package main

import (
	"fmt"
	"time"
)

type Person struct {
	ID    int
	Name  string
	Age   int
	Email string
}

// Этот метод принимает указатель на экземпляр структуры *Person
// и выводит в консоль данные об ученике
func (p *Person) Introduce() {
	fmt.Printf("🎓 Student #%d: %s, %d years old\n", p.ID, p.Name, p.Age)
	fmt.Printf("   📧 Contact: %s\n", p.Email)
}

// Этот метод увеличивает возраст студента на 1 год
func (p *Person) Birthday() {
	p.Age++
}

// Этот метод изменяет email студента на новый,
// который попросит ввести в консоль
func (p *Person) ChangeEmail() {
	var newEmail string
	fmt.Printf("Введите новый email для студента #%d: %s, %d years old\n", p.ID, p.Name, p.Age)
	fmt.Scanln(&newEmail)
	p.Email = newEmail
}

// конструктор принимает на вход три значения, которые запишет
// в поля структуры созданного экземпляра и вернёт указатель на экземпляр структуры
func NewPerson(id int, name string, age int, email string) *Person {
	return &Person{
		ID:    id,
		Name:  name,
		Age:   age,
		Email: email,
	}
}

func main() {
	students := []*Person{}
	id := 0
	name := ""
	age := 0
	email := ""

	for range 2 {
		// последовательно вводим все данные на ученика
		fmt.Printf("id: %d. Введите имя ученика: ", id)
		fmt.Scanln(&name)

		fmt.Printf("id: %d. Имя: %s. Введите возраст ученика: ", id, name)
		fmt.Scanln(&age)

		fmt.Printf("id: %d. Имя: %s. Возраст: %d. Введите e-mail ученика: ", id, name, age)
		fmt.Scanln(&email)

		// сообщение что запись об ученике успешно создана
		fmt.Printf("Запись об ученике создана! id: %d. Имя: %s. Возраст: %d. E-mail:%s \n\n", id, name, age, email)

		students = append(students, NewPerson(id, name, age, email))
		// меняем id для каждого последующего ученика
		id++
	}

	// применим методы Birthday и ChangeEmail для 1 в слайсе ученика
	students[0].Birthday()
	students[0].ChangeEmail()

	for _, student := range students {
		student.Introduce()
		time.Sleep(2 * time.Second)
	}

}
