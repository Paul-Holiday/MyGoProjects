package main

import "fmt"

func main() {
	s1 := []int{0, 0, 1, 3, 3, 5, 6, 6, 8, 9, 10, 11, 11}
	s2 := []int{3, 4, 5, 7, 7, 7, 9, 13, 13}
	result := MergeSortedSlices(s1, s2)
	fmt.Println(result) // [0 1 3 4 5 6 7 8 9 10 11 13]
}

// третья версия функции, которая не дублирует элементы AI говорит я отлично справился, но предлагает вынести
// проверку дубликатов в отдельную функцию и сделать тесты, но я уже наелся с этой задачей, пойду к следующей
func MergeSortedSlices(s1, s2 []int) []int {
	result := make([]int, 0, len(s1)+len(s2))
	i := 0
	j := 0
	for i < len(s1) && j < len(s2) {
		if s1[i] <= s2[j] {
			if len(result) == 0 || s1[i] != result[len(result)-1] {
				result = append(result, s1[i])
			}
			i++
		} else {
			if len(result) == 0 || s2[j] != result[len(result)-1] {
				result = append(result, s2[j])
			}
			j++
		}
	}

	for i < len(s1) {
		if len(result) == 0 || s1[i] != result[len(result)-1] {
			result = append(result, s1[i])
		}
		i++
	}
	for j < len(s2) {
		if len(result) == 0 || s2[j] != result[len(result)-1] {
			result = append(result, s2[j])
		}
		j++
	}
	return result
}

// это условие из моей первой версии алгоритма, моё первое г. не работало, зависало в бесконечном цикле
/*if s1[count_s1%len(s1)] > s2[count_s2%len(s2)] && s1[count_s1%len(s1)] > result[len(result)-1] {
	result = append(result, s1[count_s1])
	count_s1++
	i++
} else if s1[count_s1%len(s1)] < s2[count_s2%len(s2)] && s2[count_s2%len(s2)] > result[len(result)-1] {
	result = append(result, s2[count_s2])
	count_s2++
	i++
} else if s1[count_s1%len(s1)] == s2[count_s2%len(s2)] && s1[count_s1%len(s1)] > result[len(result)-1] {
	result = append(result, s1[count_s1], s2[count_s2])
	count_s1++
	count_s2++
	i += 2
}
*/

/*
// вторая версия моей функции слияния, она работает, хранит все элементы, даже дубликаты, что в принципе соответствует условию, но обычно в юзкейсах таких алгоритмов не дублируют элементы
// также есть некоторые вопросы по синтаксису:
func MergeSortedSlices(s1, s2 []int) []int {
	result := make([]int, 0, len(s1)+len(s2))
	count_s1 := 0 // классическое обозначение счётчиков - i, j (я осознанно их заменил на count чтобы не путаться)
	count_s2 := 0
	for count_s1 < len(s1) || count_s2 < len(s2) {
		if count_s1 < len(s1) && count_s2 < len(s2) {
			if s1[count_s1] < s2[count_s2] {
				result = append(result, s1[count_s1])
				count_s1++
			} else if s1[count_s1] > s2[count_s2] {
				result = append(result, s2[count_s2])
				count_s2++
			} else {
				result = append(result, s1[count_s1], s2[count_s2])
				count_s1++
				count_s2++
			}
		} else if count_s1 < len(s1) {
			result = append(result, s1[count_s1])
			count_s1++
		} else {
			result = append(result, s2[count_s2])
			count_s2++
		}
	}
	return result
}
*/
