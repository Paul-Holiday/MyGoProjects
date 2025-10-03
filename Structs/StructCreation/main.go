package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	Email string
}

func main() {
	// создаю экземпляр структуры 1 способом
	id1 := Person{
		Name:  "Ivan",
		Age:   15,
		Email: "ivan@randommailbox.com",
	}
	// создаю экземпляр структуры 2 способом
	id2 := Person{"Jake", 17, "jake@anothermailbox.com"}
	// создаю экземпляр структуры 3 способом
	var id3 Person
	id3.Name = "Tasha"
	id3.Age = 22
	id3.Email = "beautifulTasha@mailmail.com"
	fmt.Printf("Person's name is %s, age is %d, email is %s.\n", id1.Name, id1.Age, id1.Email)
	fmt.Printf("Person's name is %s, age is %d, email is %s.\n", id2.Name, id2.Age, id2.Email)
	fmt.Printf("Person's name is %s, age is %d, email is %s.\n", id3.Name, id3.Age, id3.Email)
}
