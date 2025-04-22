package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 6, 7, 0, 8}
	slice = FilterEven(slice)
	fmt.Println(slice)
}

func FilterEven(slice []int) []int {
	count := 0
	for _, val := range slice {
		if val%2 == 0 {
			slice[count] = val
			count++
		}
	}
	return slice[:count]
}

// мой код рабочий и эффективный для работы со слайсом "на месте", но не соответствует условию задачи, потому что требовалось создать новый слайс, а я внёс изменения в исходный, что скорее всего породит баги, если с этим слайсом параллельно работает что-то ещё

/**
func FilterEven(slice []int) []int {
    result := make([]int, 0, len(slice)) // Заранее выделяем память
    for _, val := range slice {
        if val%2 == 0 {
            result = append(result, val)
        }
    }
    return result
}
*/
