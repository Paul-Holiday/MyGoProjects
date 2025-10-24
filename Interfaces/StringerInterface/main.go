package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Year   int
}

func (b Book) String() string {
	return fmt.Sprintf("Книга: %s / %s (%d)", b.Title, b.Author, b.Year)
}

func main() {
	book1 := Book{"Алиса в стране чудес", "Кэрролл Л.", 1865}
	fmt.Println(book1)
}
