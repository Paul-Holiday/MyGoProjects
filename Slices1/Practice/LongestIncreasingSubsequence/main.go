package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 0, 4, 5, 6, 5, 5, 5, 4, 5, 6, 7, 8, 10, 11, 12, 14, 1, -1, -3, -5, -2, -1, 0, 3, 10, 11}
	lis := LongestIncreasingSubsequence(slice)
	fmt.Println(lis)
}

// решение данной задачи с помощью жадного алгоритма (с бинарным поиском) оказалось намного более простым и эффективным решением,
// чем решение с помощью динамического программирования. Сложность O(log(n)), а логика намного проще
func LongestIncreasingSubsequence(slice []int) int {
	if len(slice) == 0 {
		return 0
	}
	tails := []int{}

	for _, v := range slice {
		i := BinarySearch(tails, v)
		if i == len(tails) {
			tails = append(tails, v)
		} else {
			tails[i] = v
		}
	}
	return len(tails)
}

// мой первый бинарный поиск, написал "руками"
// на самом деле оказалось проще, чем  казалось
func BinarySearch(nums []int, x int) int {

	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if x == nums[mid] {
			return mid
		} else if x > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

/*
func LongestIncreasingSubsequence(slice []int) int {
	if len(slice) == 0 {
		return 0
	}

	count := 1
	MaxLen := 1

	for i := 0; i < len(slice)-1; i++ {
		if slice[i+1] > slice[i] {
			count++
			if count > MaxLen {
				MaxLen = count
			}
		} else {
			count = 1
		}

	}

	return MaxLen
}
*/

/*
func LongestIncreasingSubsequence(slice []int) (int, []int) {
	if len(slice) == 0 { // проверка на пустой слайс
		return 0, nil
	}

	dp := make([]int, len(slice)) // создаю слайс dp чтобы хранить в нём LIS для slice[i]
	maxLen := 1                   // минимальная длина LIS в любом случае будет как минимум 1,
	// так как слайс не пустой

	prev := make([]int, len(slice)) // создаю слайс для хранения предыдущих элементов,
	// участвовавших в формировании последовательности
	LastIndex := 0 // индекс последнего элемента в LIS будет равен 0, если слайс состоит из одного элемента или, если в нём не будет последовательности длиннее чем LIS=1, в любом другом случае это значение перезапишется
	for i := range dp {
		dp[i] = 1    // наполняю dp единицами, потому что каждый элемент из slice по сути будет образовывать собой последовательность длиной 1
		prev[i] = -1 // так обозначим индекс последнего элемента (так как в конце будем считывать с конца, восстанавливая последовательность)
	}

	for i := 1; i < len(slice); i++ {
		// прохожу по всем элементам slice начиная со второго (1),
		// так как первый элемент (0) уже обработан (уже dp[0] = 1)
		for j := 0; j < i; j++ {
			// прохожу по всем предыдущим элементам стоящим до slice[i] и сравниваю, чтобы
			if slice[i] > slice[j] {
				// если текущий элемент больше того, что проверяется на данной итерации j,
				// (из предыдущих), то обновляем наше текущее значение dp[i],
				// чтобы оно содержало значение длины предыдущей цепочки (dp[j]) + наш новый найденный элемент (+1)
				// (эта операция выполняется чуть позже, после выполнения следующего условия)
				if dp[j]+1 > dp[i] {
					// обязательно нужно выбрать между новым увеличенным значением dp[j]+1 и dp[i]
					// сюда в начале любом случае запишется любое значение больше единицы
					// (конечно, если выполняется предыдущее условие), так как в начале наш массив
					// был заполнен единицами, а затем с каждой итерацией подберётся наибольшее значение
					// (выберется цепочка с наибольшим количеством элементов) согласно предыдущему условию
					dp[i] = dp[j] + 1
					prev[i] = j // записываем индекс (как бы адрес) предыдущего элемента, который сохранится в текущей ячейке prev[i] (по сути в ячейке окажется индекс того элемента, через значение которого и строится текущая самая длинная последовательность)
				}
			}
		}
		if dp[i] > maxLen {
			maxLen = dp[i]
			LastIndex = i
			// если у нас появилась новая самая длинная последовательность,
			// то сохраним её в переменную для вывода
			// а так же сохраняем индекс её последнего элемента, он не может храниться вместе с предыдущими
		}
		fmt.Printf("Slice:    %v\ndp:       %v\nprevious: %v\nMaxLen = %d, LastIndex = %d\n_________________\n", slice, dp, prev, maxLen, LastIndex)
	}

	sequence := make([]int, 0, maxLen)
	for i := LastIndex; i != -1; i = prev[i] {
		sequence = append(sequence, slice[i])
	}
	fmt.Printf("Sequence: %v\n", sequence)

	for i, j := 0, len(sequence)-1; i < j; i, j = i+1, j-1 {
		sequence[i], sequence[j] = sequence[j], sequence[i]
	}

	return maxLen, sequence
}
*/
