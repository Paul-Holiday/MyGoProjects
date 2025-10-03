package main

import "fmt"

type Person struct {
	ID    int
	Name  string
	Age   int
	Email string
}

// конструктор принимает на вход три значения, которые запишет в поля структуры созданного экземпляра и вернёт указатель на экземпляр структуры
func NewPerson(id int, name string, age int, email string) *Person {
	return &Person{
		ID:    id,
		Name:  name,
		Age:   age,
		Email: email,
	}
}

func main() {
	students := []Person{}
	id := 0
	name := ""
	age := 0
	email := ""

	for range 5 {
		// последовательно вводим все данные на ученика
		fmt.Printf("id: %d. Введите имя ученика: ", id)
		fmt.Scanln(&name)

		fmt.Printf("id: %d. Имя: %s. Введите возраст ученика: ", id, name)
		fmt.Scanln(&age)

		fmt.Printf("id: %d. Имя: %s. Возраст: %d. Введите e-mail ученика: ", id, name, age)
		fmt.Scanln(&email)

		// сообщение что запись об ученике успешно создана
		fmt.Printf("Запись об ученике создана! id: %d. Имя: %s. Возраст: %d. E-mail:%s \n\n", id, name, age, email)

		students = append(students, *NewPerson(id, name, age, email))
		// меняем id для каждого последующего ученика
		id++
	}

	fmt.Println(students)

}
