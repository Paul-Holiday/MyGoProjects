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

// СДЕЛАТЬ МЕТОДЫ:
// Set - добавление значения
func (c *Cache) Set(key int, value interface{}) {
	c.Storage[key] = value
}

// Get - получение значения
func (c *Cache) Get(key int) (interface{}, bool) {
	value, ok := c.Storage[key] // флаг, есть ли значение вообще
	return value, ok
}

// Delete - удаление значения
func (c *Cache) Delete(key int) {
	delete(c.Storage, key)
}

// Size - количество элементов
func (c *Cache) Size() int {
	return len(c.Storage)
}

// Clear - очистка кэша
func (c *Cache) Clear() {
	clear(c.Storage)
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

	fmt.Println("Количество элементов в кэше:")
	len := cache.Size()
	fmt.Println(len)

	cache.Set(6, "Six")
	cache.Set(7, "Seven")

	fmt.Println("Количество элементов в кэше:")

	len = cache.Size()
	fmt.Println(len)

	key := 6
	value, ok := cache.Get(key)
	if ok {
		fmt.Println(value)
	} else {
		fmt.Printf("Элемента по ключу %d нет в кэше\n", key)
	}

	key = 8
	value, ok = cache.Get(key)
	if ok {
		fmt.Println(value)
	} else {
		fmt.Printf("Элемента по ключу %d нет в кэше\n", key)
	}

	fmt.Println("Вывод всех элементов:")
	for _, item := range cache.Storage {
		fmt.Println(item)
	}

	cache.Delete(7)
	fmt.Printf("Количество элементов в кэше после удаления: %d; Вывод всех элементов:\n", cache.Size())
	for _, item := range cache.Storage {
		fmt.Println(item)
	}

	cache.Clear()
	fmt.Printf("Количество элементов в кэше после очистки: %d; Вывод всех элементов:\n", cache.Size())
	for _, item := range cache.Storage {
		fmt.Println(item)
	}
}
