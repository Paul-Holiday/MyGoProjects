package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// -- Структуры--
// Поля Book
type Book struct {
	title  string
	author string
	isbn   string
	year   int
}

// Поля Library
type Library struct {
	book []*Book
	name string
}

// --КОНСТРУКТОРЫ--
// Конструктор для Book
func NewBook(title, author, isbn string, year int) *Book {
	return &Book{title, author, isbn, year}
}

// Конструктор для Library
func NewLibrary(b []*Book, name string) *Library {
	return &Library{b, name}
}

// --МЕТОДЫ ДЛЯ Book--
// Изменение параметров книги (для администрирования системы)
func (b *Book) ChangeBookAttribute() {
	// обработка нила для безопасного поведения
	if b == nil {
		fmt.Println("Книга, которую пытались изменить == nil. (Книга не найдена) Ошибка!")
		return
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("--МЕНЮ ИЗМЕНЕНИЯ ПАРАМЕТРОВ КНИГИ--\n\n Введите:\n1 - Чтобы изменить название книги.\n2 - Чтобы изменить автора книги.\n3 - Чтобы изменить ISBN книги.\n4 - Чтобы изменить год публикации книги.\n")
	var parameter int
	_, err := fmt.Scanln(&parameter)
	if err != nil {
		fmt.Println("Ошибка ввода!")
		return
	} else {
		switch parameter {
		case 1:
			fmt.Printf("Введите новое название книги: ")
			input, _ := reader.ReadString('\n')
			b.title = strings.TrimSpace(input)
		case 2:
			fmt.Printf("Введите нового автора книги: ")
			input, _ := reader.ReadString('\n')
			b.author = strings.TrimSpace(input)
		case 3:
			fmt.Printf("Введите новый ISBN книги: ")
			input, _ := reader.ReadString('\n')
			b.isbn = strings.TrimSpace(input)
		case 4:
			fmt.Printf("Введите новый год публикации книги: ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			b.year, err = strconv.Atoi(input)
			if err != nil {
				fmt.Println("Ошибка ввода! Год должен быть числом!")
				return
			}
		default:
			fmt.Println("Ошибка ввода! Для выбора изменяемого параметра нажмите 1-4!")
			return
		}
		fmt.Printf("\nПараметр %d изменён.\n", parameter)
	}
}

// --МЕТОДЫ ДЛЯ Library--
// Добавить книгу
func (lib *Library) AddBook(newbook *Book) {
	lib.book = append(lib.book, newbook)
}

// Найти книгу по названию или ISBN
func (lib *Library) FindBookBy() *Book {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("\n--МЕНЮ ПОИСКА КНИГИ--\n\nВведите:\n1 - Поиск по названию.\n2 - Поиск по isbn.\n")
	var parameter int
	_, err := fmt.Scanln(&parameter)

	// обработка пустого ввода
	if err != nil {
		fmt.Println("Ошибка ввода!")
		return nil
	} else {
		switch parameter {
		// Поиск по названию
		case 1:
			fmt.Printf("\nВведите название: ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			// обработка пустого ввода названия
			if len(input) == 0 {
				fmt.Println("Ошибка ввода! Нужно ввести название книги!")
				return nil
			} else {
				// цикл поиска по названию, флаг чтобы был вывод, если книги нет
				flag := false
				for _, book := range lib.book {
					if strings.EqualFold(book.title, input) {
						fmt.Printf("Книга %v найдена!\n", book)
						flag = true
						return book
					}
				}
				if !flag {
					fmt.Println("Книга не найдена!")
					return nil
				}
			}
		// Поиск по ISBN
		case 2:
			fmt.Printf("\nВведите ISBN в формате xxx-x-xxxxx-xxx-x: ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			// обработка пустого ввода
			if len(input) == 0 {
				fmt.Println("Ошибка ввода! Нужно ввести ISBN!")
				return nil
			} else {
				// цикл поиска по isbn, флаг чтобы был вывод, если книги нет
				flag := false
				for _, book := range lib.book {
					if book.isbn == input {
						fmt.Printf("Книга %v найдена!", book)
						flag = true
						return book
					}
				}
				if !flag {
					fmt.Println("Книга не найдена!")
					return nil
				}
			}
		// Если ввели не 1 или 2, то ошибка
		default:
			fmt.Println("Ошибка ввода!")
			return nil
		}
	}
	return nil
}

// Удалить книгу по названию
func (lib *Library) DeleteBookByTitle() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите название книги для её удаления, введите 'отмена', чтобы ничего не удалять.")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Обработка отмены
	if strings.EqualFold(input, "отмена") {
		fmt.Println("Удаление было прервано пользователем.")
		return
	}

	// Обработка пустого ввода названия
	if len(input) == 0 {
		fmt.Println("Ошибка ввода! Нужно ввести название книги!")
		return
	} else {
		// Поиск книги и удалене из коллекции, если найдена, если нет, то выводим сообщение о неудаче
		fmt.Printf("\nУдаление книги %s из библиотеки...", input)
		flag := false
		for index, book := range lib.book {
			if strings.EqualFold(book.title, input) {
				lib.book = append(lib.book[:index], lib.book[index+1:]...)
				fmt.Printf("\nКнига %s была удалена из библиотеки.\n", book.title)
				flag = true
				break
			}
		}
		if !flag {
			fmt.Printf("\nКнига %s не найдена. Невозможно удалить.\n", input)
			return
		}
	}
}

// Сортировать книги по названию
func (lib *Library) SortByTitle() {
	fmt.Printf("\n--МЕНЮ СОРТИРОВКИ КНИГ--\n\nВведите:\n1 - Сортировка по алфавиту.\n2 - Сортировка в обратном порядке.\n")
	var parameter int
	_, err := fmt.Scan(&parameter)

	// обработка пустого ввода
	if err != nil {
		fmt.Println("Ошибка ввода!")
	} else {
		switch parameter {
		case 1:
			sort.Slice(lib.book, func(i, j int) bool {
				return lib.book[i].title < lib.book[j].title
			})
			fmt.Println("Сортировка выполнена успешно!")
		case 2:
			sort.Slice(lib.book, func(i, j int) bool {
				return lib.book[i].title > lib.book[j].title
			})
			fmt.Println("Сортировка выполнена успешно!")
		default:
			fmt.Println("Ошибка ввода!")
		}
	}
}

// Геттер для Book
func (b *Book) GetBook() *Book {
	return b
}

// Вывод всех книг
func (lib *Library) DisplayAllBooks() {
	fmt.Println("\n--ВСЕ КНИГИ В БИБЛИОТЕКЕ--\n\n")
	for i, b := range lib.book {
		fmt.Printf("%d. \"%s\" - %s (%d) [ISBN: %s]\n",
			i+1, b.title, b.author, b.year, b.isbn)
	}
}

func main() {
	// демонстрация работы конструктора *Book
	book1 := NewBook("Алиса в стране чудес", "Льюис Кэрролл", "111-2-33333-444-5", 2007)

	shelf1 := []*Book{
		{"Винни-Пух", "Алан Александр Милн", "222-3-44444-555-6", 1973},
		{"Унесенные ветром", "Маргарет Митчелл", "978-5-38917-583-9", 1992},
		{"Бедные люди", "Федор Достоевский", "978-5-38909-710-0", 1961},
		{"Война и мир. Том 1", "Лев Толстой", "978-5-04108-738-8", 2011},
	}
	// демонстрация работы конструктора *Library
	lib1 := NewLibrary(shelf1, "Классика")

	// демонстрация изменения параметров книги
	// изменим произвольный isbn на реально существующий 978-5-35311-344-7
	shelf1[0].ChangeBookAttribute()

	// добавим книгу на полку
	lib1.AddBook(book1)

	// можем внести изменения и в библиотеке для этого по названию найдем книгу
	// "Алиса в стране чудес" и поменяем ISBN на реальный 978-5-00132-357-0
	lib1.FindBookBy().ChangeBookAttribute()

	// удалим книгу Унесенные ветром из библиотеки
	lib1.DeleteBookByTitle() // не реализовал меню в консоли так как не хватило мотивации и сил, просто ввожу название в качестве аргумента функции

	//выводим одну книгу
	fmt.Println(lib1.FindBookBy().GetBook())

	// выводим все книги
	lib1.DisplayAllBooks()

	// сортируем книги по названию
	lib1.SortByTitle()

	//выводим все книги
	lib1.DisplayAllBooks()
}
