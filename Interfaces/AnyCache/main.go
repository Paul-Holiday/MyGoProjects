package main

import "fmt"

type Cache struct {
	storage map[int]interface{}
}

func NewCache() *Cache {
	return &Cache{
		storage: make(map[int]interface{}),
	}
}

// СДЕЛАТЬ МЕТОДЫ:
// Set - добавление значения
func (c *Cache) Set(key int, value interface{}) {
	c.storage[key] = value
}

// Get - получение значения
func (c *Cache) Get(key int) (interface{}, bool) {
	value, ok := c.storage[key] // флаг, есть ли значение вообще
	return value, ok
}

// Delete - удаление значения
func (c *Cache) Delete(key int) {
	delete(c.storage, key)
}

// Size - количество элементов
func (c *Cache) Size() int {
	return len(c.storage)
}

// Clear - очистка кэша
func (c *Cache) Clear() {
	clear(c.storage)
}

// ТОЧКА ВХОДА
func main() {
	fmt.Printf("")
	a := 54
	b := "str"
	c := 'C'
	d := 12.34
	f := []uint{0, 88, 32, 45, 67}

	cache := NewCache()
	cache.Set(1, a)
	cache.Set(2, b)
	cache.Set(3, c)
	cache.Set(4, d)
	cache.Set(5, f)

	fmt.Printf("Количество элементов в кэше: %d\n\n", cache.Size())

	fmt.Println("Добавляем два элемента:")
	cache.Set(6, "Six")
	cache.Set(7, "Seven")

	fmt.Printf("Количество элементов в кэше: %d\n\n", cache.Size())

	key := 6
	fmt.Printf("Вывод элемента %d:\n", key)
	value, ok := cache.Get(key)
	if ok {
		fmt.Printf("%v\n\n", value)
	} else {
		fmt.Printf("Элемента по ключу %d нет в кэше\n\n", key)
	}

	key = 8
	fmt.Printf("Вывод элемента %d:\n", key)
	value, ok = cache.Get(key)
	if ok {
		fmt.Printf("%v\n\n", value)
	} else {
		fmt.Printf("Элемента по ключу %d нет в кэше\n\n", key)
	}

	fmt.Println("Вывод всех элементов:")
	for _, item := range cache.storage {
		fmt.Println(item)
	}
	fmt.Printf("\nУдаляем один элемент\n")
	cache.Delete(7)

	fmt.Printf("Количество элементов в кэше после удаления: %d; Вывод всех элементов:\n", cache.Size())
	for _, item := range cache.storage {
		fmt.Println(item)
	}
	fmt.Printf("\nОчищаем кэш\n")
	cache.Clear()

	fmt.Printf("Количество элементов в кэше после очистки: %d; Вывод всех элементов:\n", cache.Size())
	for _, item := range cache.storage {
		fmt.Println(item)
	}
}
