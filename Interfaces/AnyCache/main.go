package main

import "fmt"

type Cache struct {
	Storage map[int]interface{}
}

func NewCache() *Cache {
	return &Cache{
		Storage: make(map[int]interface{}),
	}
}

func main() {
	fmt.Printf("")
	a := 54
	b := "str"
	c := 'C'
	d := 12.34
	f := []uint{0, 88, 32, 45, 67}

	cache := NewCache()
	cache.Storage[1] = a
	cache.Storage[2] = b
	cache.Storage[3] = c
	cache.Storage[4] = d
	cache.Storage[5] = f

	for _, item := range cache.Storage {
		fmt.Println(item)
	}
}
