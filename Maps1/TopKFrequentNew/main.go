package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{0, 0, 0, 4, 4, 4, 4, 7, 7, 7, 7, 7, 11, 22, 11, 5, 11, 6}
	fmt.Println(TopKFrequentNew(numbers, 3))
}

// Функция принимает на вход слайс с интами и переменную типа
// инт - количество самых встречаемых элементов в слайсе.

// Функция возвращает слайс - топ k сымых встречаемых элементов.

// Если функция приняла пустой слайс, то вернёт пустой слайс
// если количество уникальных элементов в слайсе меньше, чем k, то вернёт все элементы что есть.

func TopKFrequentNew(nums []int, k int) []int {
	// Проверка на пустой слайс
	if len(nums) == 0 {
		return []int{}
	}
	// Создание мапы для подсчёта повторений (частот) элементов
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	// Наполняем новый слайс уникальными элементами из мапы
	uniqueNums := make([]int, 0, len(freqMap))
	for num := range freqMap {
		uniqueNums = append(uniqueNums, num)
	}

	// Сортируем слайс с уникальными элементами, в качестве условия сортировки
	// обращаемся к соответствующей каждому элементу частоте в мапе.
	sort.Slice(uniqueNums, func(i, j int) bool {
		return freqMap[uniqueNums[i]] > freqMap[uniqueNums[j]]
	})

	// Если k больше, чем количество уникальных элемнтов, то возвращаем весь слайс
	if k > len(uniqueNums) {
		return uniqueNums
	}

	// Возвращаем первые k элементов из отсортированного слайса
	return uniqueNums[:k]
}
