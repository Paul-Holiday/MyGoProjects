package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice = ReverseSlice(slice)
	fmt.Println(slice)
}

func ReverseSlice(slice []int) []int {
	length := len(slice)
	for ind, _ := range slice {
		if ind == length/2 {
			break
		}
		slice[ind], slice[length-1-ind] = slice[length-1-ind], slice[ind]
	}
	return slice
}

/*код получился хорошим, но есть замечания. В данной задаче лучше использовать полный for,
потому что условие выхода из цикла выглядит громоздко и не лаконично. Лучше было написать так:
func ReverseSlice(slice []int) {
    length := len(slice)
    for i := 0; i < length/2; i++ {
        slice[i], slice[length-1-i] = slice[length-1-i], slice[i]
    }
}
*/
