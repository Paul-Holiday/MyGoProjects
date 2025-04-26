package main

import "fmt"

func main() {
	a := []int{0, 0, 0, 4, 5, 6, 8, 9, 5, 3, 1, 2}
	b := []int{5, 3, 12, -5, 2, 1, -7}
	fmt.Println(FindIntersection(a, b))
}

func FindIntersection(a, b []int) []int {
	intersectionMap := make(map[int]struct{})

	intersectionSlice := make([]int, 0, min(len(a), len(b)))

	for _, v := range a {
		intersectionMap[v] = struct{}{}
	}

	for _, v := range b {
		_, exists := intersectionMap[v]
		if exists {
			intersectionSlice = append(intersectionSlice, v)
			delete(intersectionMap, v)
		}
	}
	return intersectionSlice
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
