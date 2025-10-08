package main

import "fmt"

type Person struct {
	ID    int
	Name  string
	Age   int
	Email string
}

// структура Person встраивается в структуру Employee
type Employee struct {
	Person
	Salary  float64
	Company string
}

// Этот метод принимает указатель на экземпляр структуры *Person
// и выводит в консоль данные о Person
func (p *Person) Introduce() {
	fmt.Printf("Person #%d: %s, %d years old\n", p.ID, p.Name, p.Age)
	fmt.Printf("📧 Contact: %s\n", p.Email)
}

// Этот метод увеличивает возраст Person на 1 год
func (p *Person) Birthday() {
	p.Age++
}

// Этот метод изменяет email Person на новый,
// который попросит ввести в консоль
func (p *Person) ChangeEmail() {
	var newEmail string
	fmt.Printf("Введите новый email для person #%d: %s, %d years old\n", p.ID, p.Name, p.Age)
	fmt.Scanln(&newEmail)
	p.Email = newEmail
}

// конструктор принимает на вход три значения, которые запишет
// в поля структуры созданного экземпляра и вернёт указатель на экземпляр Person
func NewPerson(id int, name string, age int, email string) *Person {
	return &Person{
		ID:    id,
		Name:  name,
		Age:   age,
		Email: email,
	}
}

func NewEmployee(p Person, salary float64, company string) *Employee {
	return &Employee{
		Person:  p,
		Salary:  salary,
		Company: company,
	}
}

func main() {
	emp := Employee{
		Person:  Person{1, "Ford", 64, "StanfordPines@mail"},
		Salary:  9000,
		Company: "GravityFallsInc",
	}

	emp.ChangeEmail()

	emp.Introduce()
	fmt.Printf("He's working at %s and earns %.0f$ per month.\n", emp.Company, emp.Salary)

	pers := NewPerson(2, "Marty", 18, "MartyMcFly@mail")
	emp2 := NewEmployee(*pers, 4000, "Hill Walley School")

	emp2.Birthday()

	emp2.Introduce()
	fmt.Printf("He's working at %s and earns %.0f$ per month.\n", emp2.Company, emp2.Salary)
}
