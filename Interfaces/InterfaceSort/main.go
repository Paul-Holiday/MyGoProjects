package main

import (
	"fmt"
)

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 | ~string
}

// интерфейс с ограничениями по типу Ordered
type Sortable[T Ordered] interface {
	Len() int           // возвращает длину слайса
	Less(i, j int) bool // возвращает	меньше ли элемент i, чем j
	Swap(i, j int)      // меняет местами i и j элементы
}

type IntSlice []int

func (ints IntSlice) Len() int {
	return len(ints)
}

func (ints IntSlice) Less(i, j int) bool {
	return ints[i] < ints[j]
}

func (ints IntSlice) Swap(i, j int) {
	ints[i], ints[j] = ints[j], ints[i]
}

type StrSlice []string

func (str StrSlice) Len() int {
	return len(str)
}

func (str StrSlice) Less(i, j int) bool {
	return str[i] < str[j]
}

func (str StrSlice) Swap(i, j int) {
	str[i], str[j] = str[j], str[i]
}

type FloatSlice []float64

func (flts FloatSlice) Len() int {
	return len(flts)
}

func (flts FloatSlice) Less(i, j int) bool {
	return flts[i] < flts[j]
}

func (flts FloatSlice) Swap(i, j int) {
	flts[i], flts[j] = flts[j], flts[i]
}

type Person struct {
	Name string
	Age  int
}

type People []Person

func (ppl People) Len() int {
	return len(ppl)
}

func (ppl People) Less(i, j int) bool {
	return ppl[i].Age < ppl[j].Age
}

func (ppl People) Swap(i, j int) {
	ppl[i], ppl[j] = ppl[j], ppl[i]
}

func BubbleSort[T Ordered](collection Sortable[T]) {
	for i := 0; i < collection.Len(); i++ {
		for j := 0; j < collection.Len()-1-i; j++ {
			if collection.Less(j+1, j) {
				collection.Swap(j+1, j)
			}
		}
	}
}

func QuickSort[T Ordered](collection Sortable[T], low, high int) {
	if low >= high {
		return
	}
	pivot := LomutoPartition[float64](collection, low, high)
	QuickSort[T](collection, low, pivot-1)
	QuickSort[T](collection, pivot+1, high)

}

func LomutoPartition[T Ordered](collection Sortable[T], low, high int) int {
	pivot := high
	i := low - 1
	for j := low; j < high; j++ {
		if collection.Less(j, pivot) {
			i++
			collection.Swap(i, j)
		}
	}
	collection.Swap(i+1, pivot)
	return i + 1
}

func HoarPartition(collection []int, low, high int) int {
	pivot := collection[low]
	i := low
	j := high

	for true {
		for i < high && collection[i] < pivot {
			i++
		}
		for collection[j] > j {
			j--
		}
		if i >= j {
			break
		}
		collection[i], collection[j] = collection[j], collection[i]
		i++
		j--
	}
	collection[low], collection[j] = collection[j], collection[low]
	return j
}

func main() {
	ints := IntSlice{0, 55, 63, 0, 90, 2, 5, -5, -40}
	fmt.Println(ints)
	fmt.Println("Sorting...")
	BubbleSort[int](ints)
	fmt.Println(ints)
	fmt.Println()

	strs := StrSlice{"g", "c", "e", "f", "a", "z", "b", "x", "y", "d", "h"}
	fmt.Println(strs)
	fmt.Println("Sorting...")
	BubbleSort[int](strs)
	fmt.Println(strs)
	fmt.Println()

	floats := FloatSlice{0.01, 0.55, 0.32, 1.1, 53.02, 54.09, 0.32, -0.45}
	fmt.Println(floats)
	fmt.Println("QuickSorting...")
	QuickSort[int](floats, 0, floats.Len()-1)
	fmt.Println(floats)
	fmt.Println()

	ppl := People{
		{Name: "Brian",
			Age: 55},
		{Name: "Alice",
			Age: 21},
		{Name: "Stephen",
			Age: 25},
		{Name: "John",
			Age: 35},
		{Name: "Carl",
			Age: 32},
		{Name: "Connor",
			Age: 42},
	}
	fmt.Println(ppl)
	fmt.Println("Sorting...")
	BubbleSort[int](ppl)
	fmt.Println(ppl)

}

// main:

// intSlice := Slice[int]{1, 42, 53, 903, -66, 0, 3, 2}
// for i, val := range intSlice {
// 	fmt.Printf("%d: %v\n", i, val)
// }
// BubbleSort(intSlice)
// fmt.Println("Sorting...")
// fmt.Println("Sorted slice:
// for i, val := range intSlice {
// 	fmt.Printf("%d: %v\n", i, val)
// }

// strSlice := Slice[string]{"b", "f", "m", "k", "z", "a", "g"}
// for i, val := range strSlice {
// 	fmt.Printf("%d: %v\n", i, val)
// }
// strSlice.Sort()
// fmt.Println("Sorting...")
// fmt.Println("Sorted slice:")
// for i, val := range strSlice {
// 	fmt.Printf("%d: %v\n", i, val)
// }

// попытка решать через слайс произвольного типа

// type Slice[T Ordered] []T

// func (s Slice[T]) Len() int {
// 	return len(s)
// }

// func (s Slice[T]) Less(i, j int) bool {
// 	return s[i] < s[j]
// }

// func (s Slice[T]) Swap(i, j int) {
// 	s[i], s[j] = s[j], s[i]
// }

// func BubbleSort[T Ordered](s Sortable[T]) {
// 	for i := 0; i < s.Len()-1; i++ {
// 		for j := 0; j < s.Len()-1-i; j++ {
// 			if s.Less(j+1, j) {
// 				s.Swap(j, j+1)
// 			}
// 		}
// 	}
// }

// func QuickSort[T Ordered](s[] Sortable[T]) {
// 	// делим слайс на подслайсы и сортируем элементы во всем слайсе, помещая в левый те элементы, что меньше опорного, а в правый те, что больше опорного, затем рекурсивно повторяем
// 	pivot :=  / 2
// 	for _, v := range s {
// 		if v >= s[pivot] {

// 		}
// 	}
// }
