package main

import (
	"fmt"
)

func main() {
	s := Memoize(square)
	fmt.Println(s(10))
	fmt.Println(s(10))
	fmt.Println(s(20))
	fmt.Println(s(20))
	d := Memoize(double)
	fmt.Println(d(10))
	fmt.Println(d(10))
	fmt.Println(d(20))
	fmt.Println(d(20))
}

func Memoize(fn func(int) int) func(int) int {
	// создал захваченную мапу для кэша
	cacheMap := make(map[int]int)

	// замыкание т.е. анонимная функция принимающая значение value,
	// которое фактически используется для расчета функции fn(value)
	return func(value int) int {
		// если такое значение value уже поступало, то результат берем из кэша
		if result, exists := cacheMap[value]; exists {
			fmt.Printf("Берем значение из кэша: %d → %d\n", value, result)
			return result
		}
		// если нет, то выполняем расчет и сохраняем значение в кэш
		result := fn(value) // весь вычислительный блок здесь, сюда можно положить любую функцию принимающую на вход одно значение int и выдающую тоже одно значение int. вся заворушка с тем, чтобы Memoize принимала на вход функцию связана как раз с желанием сделать эту функцию кэширования универсальной. То есть подставляй любую математику, всё закешируется конкретно для того вычислительного блока, что ты выбрал.
		cacheMap[value] = result
		fmt.Printf("Выполняем расчёт и сохраняем в кэш: %d → %d\n", value, result)
		return result
	}
}

func square(a int) int {
	return a * a
}

func double(a int) int {
	return 2 * a
}
