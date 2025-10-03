package main

import (
	"fmt"
	"time"
)

func main() {
	f := MemoizeLazyTTL(Pow2, 3*time.Second)
	fmt.Println(f(15))
	fmt.Println(f(15))
	time.Sleep(4 * time.Second)
	fmt.Println(f(15))

}

type CacheItem struct {
	Value     int
	ExpiresAt time.Time
}

// создаю функцию принимающую на вход другую функцию (функцию-действие)
// и возвращающую другую функцию (в которой мы будем проводить наше действие)
func MemoizeLazyTTL(fn func(key int) int, ttl time.Duration) func(int) int {
	// во внешней функции создаю мапу, чтобы она оказалась
	// захваченной для дочерней функции
	cacheMap := make(map[int]CacheItem)

	// возвращаю функкцию в которой проверю наличие готового решения
	// функции fn в кэше, если оно есть, то выведу его,
	// если его нет, то запущу fn и запишу результат в кэш
	// и выведу полученное число
	return func(key int) int {
		// проверяем наличие элемента в мапе и его TTL
		if item, exists := cacheMap[key]; exists {
			if time.Now().Before(item.ExpiresAt) {
				fmt.Println("Результат взят из кэша.")
				return item.Value
			} else {
				delete(cacheMap, key)
				fmt.Println("Результат устарел и был удален из кэша. Пересчитываю.")
			}
		}
		result := fn(key)
		// сохраняем в структуру результат и время жизни,
		// структура будет находиться в мапе
		cacheMap[key] = CacheItem{
			Value:     result,
			ExpiresAt: time.Now().Add(ttl),
		}
		return result
	}
}

func Pow2(a int) int {
	return a * a
}
