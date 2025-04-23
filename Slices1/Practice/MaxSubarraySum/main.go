package main

import (
	"fmt"
	"math"
)

func main() {
	slice := []int{0, 3, 4, 5, 6, 6, 9, 9, 9, 0, 8, 5, 3, 0}
	k := 3
	result, err := MaxSubarraySum(slice, k)
	fmt.Println(result, err) // 27
}

func MaxSubarraySum(slice []int, k int) (int, error) {
	if len(slice) == 0 {
		return 0, fmt.Errorf("слайс не может быть пустым")
	}
	if k <= 0 || k > len(slice) {
		return 0, fmt.Errorf("значение k должно лежать в диапазоне 0 < k <= %d (длина вашего массива)", len(slice))
	}
	// сообщение об ошибке "значение k <= 0 или > длины слайса" не дает инфы о нужных значениях, изменил на более информативное

	maxSum := math.MinInt32
	// я забыл, что переменные начинающиеся с заглавной буквы становятся глобальными, лучше писать с маленькой (изменил MaxSum на maxSum, currentSum->currentSum)
	currentSum := 0

	for i := 0; i < k; i++ {
		currentSum += slice[i]
	}
	maxSum = currentSum

	for i := k; i < len(slice); i++ {
		currentSum += slice[i] - slice[i-k]
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}
	return maxSum, nil
}

/*
package main

import "fmt"

func main() {
	slice := []int{0, 3, 4, 5, 6, 6, 9, 9, 9, 0, 8, 5, 3, 0}
	k := 3
	result := MaxSubarraySum(slice, k)
	fmt.Println(result) // 27
}

func MaxSubarraySum(slice []int, k int) int {
	if k > len(slice) { // граничное условие проработано, но не полностью, не рассмотрен вариант с k < 0
		return 0 // лучше вернуть ошибку (пока не умею этого делать), потому что значение 0 может быть не очевидно
	}

	// лучше назвать переменную отражающим смысл названием, например MaxSum
	res := 0 // этой переменной лучше присвоить значение math.MinInt32, то есть самое маленькое возможное int32 число (-2147483648)
	// то же самое, лучше назвать переменную, к примеру CurrentSum
	sum := 0

	for i := 0; i < k; i++ {
		sum += slice[i]
	}
	res = sum
	fmt.Printf("sum = %d, res = %d\n", sum, res)

	// скользящее окно удобнее писать начиная с индекса k и двигаясь до полной длины слайса
	// моя логика нагромождённая (я сам сомневался заработает ли и подгонял один из индексов компиляциями)
	// лучше переделать
	for i := 1; i <= len(slice)-k; i++ {
		sum -= slice[i-1]
		sum += slice[k+i-1] // это можно делать в одном действии, что будет более читаемо
		if sum > res {
			res = sum
		}
		fmt.Printf("sum = %d, res = %d\n", sum, res)
	}
	return res
}

*/
