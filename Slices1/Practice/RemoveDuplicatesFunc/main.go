package main

import "fmt"

func main() {
	slice := []int{0, 0, 1, 5, 1, 4, 4, 3, 2, 1, 6, 7, 0, 5}
	result := RemoveDuplicates(slice)
	fmt.Println(result) // [1 5 4 6 0 3 2 7]
}

func RemoveDuplicates(slice []int) []int {
	result := make([]int, 0, len(slice)) // я выделяю память на весь слайс, что не очень разумно, лучше написать len(slice)/2 или 0 и Go если надо сам оптимизирует
	resmap := make(map[int]struct{}, len(slice))

	for _, v := range slice {
		if _, ok := resmap[v]; !ok { // вместо переменной ок лучше писать exists чтобы улучшить читаемость
			resmap[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}

// мой код рабочий, но не сохраняет порядок элементов - это плохо. На собесах ожидают сохранение порядка элементов. upd. перенес аппенд элементов из мапы в тот же цикл и порядок сразу сохранился, оказалось проще, чем я думал. Ещё некоторые замечания от AI оставил в комментах
