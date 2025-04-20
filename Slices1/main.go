package main

import "fmt"

func Names(s []string) {
	s[2] = "Franklin"
}

func main() {
	arr := [7]int{0, 1, 2, 3, 4, 5, 6} // так инициализируется массив !!!равно перед списком значений не нужно!!!

	var slice []int = arr[:7] // так создается слайс

	fmt.Println(slice[:]) // так выводится слайс

	str := [3]string{"Michael", "Trevor", "CJ"} // снова инициализация массива с именами

	sli := str[:] // слайс можно объявлять даже без указания типа данных массива и границ слайса (!!!двоеточие обязательно!!! будет взят массив полностью)

	Names(sli) // функция получает на вход слайс

	fmt.Println(str) // но изменения в слайсе влияют на исходный массив, потому что слайс содержит указатели на "лежащий под ним" массив

	logic := []bool{true, true, false, false} // создать слайс можно и без предварительного создания массива, он создастся неявно
	logic[3] = true
	fmt.Println(logic)

	cars := []struct { //создаю слайс структуру по примеру из go.dev
		model      string
		maxSpeed   int
		isElectric bool
	}{
		{"Mustang", 250, false},
		{"Cybertruck", 200, true},
		{"Viper", 280, false},
		{"Civic", 180, false},
		{"Leaf", 140, true},
	}

	fmt.Println(cars)

	fmt.Println("Car in slot 2 is", cars[1].model) // вывожу только одно поле из слайса с машинами, по сути таким образом можно сделать структуру в структуре с помощью слайсов и удобно обращаться к данным, надо выяснить удобно ли это, или есть способы сделать это лучшим образом
	i := 0

	for i < len(cars) { // просто вывод через цикл
		fmt.Println("#", i+1, "car in garage is", cars[i].model)
		fmt.Println(cars[i].model, "max Speed is", cars[i].maxSpeed)
		if cars[i].isElectric {
			fmt.Println(cars[i].model, "is an Elecric car!")
		} else {
			fmt.Println(cars[i].model, "is not an Elecric car!\n")
		}

		i++
	}

	c := cap(cars)
	l := len(cars)
	fmt.Println("cap of slice cars -", c, "length of slice cars - ", l) // вместимость (ёмкость) и длина слайса cars
}
