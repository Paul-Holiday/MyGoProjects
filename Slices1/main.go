package main

import "fmt"

func main() {
	arr := [7]int{1, 2, 3, 4, 5, 6, 7}
	var slice []int = arr[:7]
	fmt.Println(slice)
}
