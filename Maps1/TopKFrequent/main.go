package main

import (
	"fmt"
	"math/rand"
)

func main() {
	chisla := []int{1, 2, 2, 2, 2, 55, 65, 65, 65, 11, 42, 42, 0, 0, 0, 0, 0}
	fmt.Println(TopKFrequent(chisla, 3)) // должны быть 0, 2, 65
}

func TopKFrequent(nums []int, k int) []int {
	// проверка на пустой слайс
	if len(nums) == 0 {
		return []int{}
	}

	numsFreq := make(map[int]int) // мапа с частотой употребления числа в значении и самим числом в ключе
	uniqNums := []int{}           // слайс с уникальными числами
	freq := []int{}               // соответствующие частоты употребления уникальных чисел

	result := []int{}

	for _, num := range nums { // считаю сколько раз встречается каждое число и записываю в мапу
		numsFreq[num]++
	}

	// наполняем слайсы уникальными (неповторяющимися) элементами и соответствующими частотами
	for num, freqcy := range numsFreq {
		uniqNums = append(uniqNums, num)
		freq = append(freq, freqcy)
	}

	// нужно отсортировать слайс с частотами и так же передвигать и уникальные элементы (слайс с частотами должен оказаться отсортированным, а соответствие "число-частота употребления" по индексам в двух слайсах сохраниться)
	freq, uniqNums = quickSort(freq, uniqNums)

	// вывод пустого слайса в случае если уникальных элементов оказалось меньше, чем требуемое количество элементов в топе
	if k > len(freq) {
		return []int{}
	}

	// возвращение топ k элементов из слайса уникальных элементов
	for i := 0; i < k; i++ {
		result = append(result, uniqNums[i])
	}

	return result
}

// кастомная сортировка слайса arr2 по слайсу arr
func quickSort(arr []int, arr2 []int) ([]int, []int) {
	if len(arr) < 2 {
		return arr, arr2
	}

	left, right := 0, len(arr)-1

	// Выбираем опорный элемент (можно рандомизировать для улучшения производительности)
	pivotIndex := rand.Intn(len(arr))
	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]
	arr2[pivotIndex], arr2[right] = arr2[right], arr2[pivotIndex]

	// Разделение массива
	for i := range arr {
		if arr[i] > arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			arr2[i], arr2[left] = arr2[left], arr2[i]
			left++
		}
	}

	// Ставим опорный элемент на правильное место
	arr[left], arr[right] = arr[right], arr[left]
	arr2[left], arr2[right] = arr2[right], arr2[left]

	// Рекурсивно сортируем левую и правую части
	quickSort(arr[:left], arr2[:left])
	quickSort(arr[left+1:], arr2[left+1:])
	return arr, arr2
}
