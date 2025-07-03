package main

import "fmt"

func main() {
	testSlice := make([]int, 5, 10) // make избыточен, так как сразу после инициализации нулями идёт перезапись этих нулей, проще сразу сделать инициализацию через литерал
	testSlice = []int{1, 2, 3, 4, 5, 6}
	sum := SumSliceFunc(testSlice)
	fmt.Println(sum)
}

func SumSliceFunc(slice []int) int { // название функции в задании было SumSlice, а я сделал SumSliceFunc, что не ошибка, но отклонение от запрошенного стиля
	if len(slice) == 0 { // избыточная проверка, так как цикл ниже не запустится при нулевой длине слайса
		return 0
	}
	sum := 0
	for i := range slice { // лучше использовать range из пары индекс-значение, так код читаемее
		sum += slice[i]
	}
	return sum
}

/**
package main

import "fmt"

func main() {
	testSlice := []int{1, 2, 3, 4, 5, 6}
	sum := SumSlice(testSlice)
	fmt.Println(sum) // 21
}

func SumSlice(slice []int) int {
	sum := 0
	for _, num := range slice {
		sum += num
	}
	return sum
}
**/
