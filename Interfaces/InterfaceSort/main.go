package main

import "fmt"

type Ordered interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64 | string
}

// Generic интерфейс для сортируемой коллекции
type Sortable[T Ordered] interface {
	Len() int           // возвращает длину слайса
	Less(i, j int) bool // возвращает	меньше ли элемент i, чем j
	Swap(i, j int)      // меняет местами i и j элементы
}

type Slice[T Ordered] []T

func (s Slice[T]) Len() int {
	return len(s)
}

func (s Slice[T]) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s Slice[T]) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// func Compare[T Ordered](a, b T) int {
// 	if a < b {
// 		return -1
// 	}
// 	if a > b {
// 		return 1
// 	} else {
// 		return 0
// 	}
// }

func (s Slice[T]) Sort() {
	for i := 0; i < s.Len()-1; i++ {
		for j := 0; j < s.Len()-1-i; j++ {
			if s.Less(j+1, j) {
				s.Swap(j, j+1)
			}
		}
	}
}

func main() {
	intSlice := Slice[int]{1, 42, 53, 903, -66, 0, 3, 2}
	for i, val := range intSlice {
		fmt.Printf("%d: %v\n", i, val)
	}
	intSlice.Sort()
	fmt.Println("Sorting...")
	fmt.Println("Sorted slice:")
	for i, val := range intSlice {
		fmt.Printf("%d: %v\n", i, val)
	}

	strSlice := Slice[string]{"b", "f", "m", "k", "z", "a", "g"}
	for i, val := range strSlice {
		fmt.Printf("%d: %v\n", i, val)
	}
	strSlice.Sort()
	fmt.Println("Sorting...")
	fmt.Println("Sorted slice:")
	for i, val := range strSlice {
		fmt.Printf("%d: %v\n", i, val)
	}
}
