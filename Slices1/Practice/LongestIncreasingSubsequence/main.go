package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 0, 4, 5, 6, 5, 5, 5, 4, 5, 6, 7, 8, 10, 11, 12, 14, 1, -1, -3, -5, -2, -1, 0, 3, 10, 11}
	lis := LongestIncreasingSubsequence(slice)
	fmt.Println(lis)
}

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
