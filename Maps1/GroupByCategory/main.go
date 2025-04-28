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
	}
	resultMap := GroupByCategory(prod)
	fmt.Println(resultMap)
}

func GroupByCategory(products []struct {
	Name     string
	Category string
}) map[string][]string {
	mapOfProds := make(map[string][]string)
	for i := range products {
		mapOfProds[products[i].Category] = append(mapOfProds[products[i].Category], products[i].Name)
	}
	return mapOfProds
}
