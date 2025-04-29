package main

import "fmt"

type Products struct {
	Name     string
	Category string
}

func main() {
	prod := []Products{
		{"Мука", "Бакалея"},
		{"Сахар", "Бакалея"},
		{"Диван", "Мебель"},
		{"Стол", "Мебель"},
		{"Чашка", "Посуда"},
		{"Тарелка", "Посуда"},
		{"Чайник", "Посуда"},
	}

	resultMap := GroupByCategory(prod)
	fmt.Println(resultMap)
}

func GroupByCategory(products []Products) map[string][]string {
	if len(products) == 0 {
		return map[string][]string{}
	}

	mapOfProds := make(map[string][]string)
	for _, p := range products {
		mapOfProds[p.Category] = append(mapOfProds[p.Category], p.Name)
	}
	return mapOfProds
}
