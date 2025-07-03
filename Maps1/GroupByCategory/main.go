package main

import "fmt"

type Products struct {
	Name     string
	Category string
}

func main() {
	goods := []Products{
		{"Мука", "Бакалея"},
		{"Чайник", "Посуда"},
		{"Сахар", "Бакалея"},
		{"Диван", "Мебель"},
		{"Стол", "Мебель"},
		{"Чашка", "Посуда"},
		{"Тарелка", "Посуда"},
	}

	resultMap := GroupByCategory(goods)
	fmt.Println(resultMap)
}

// группирую по ключу
func GroupByCategory(products []Products) map[string][]string {
	if len(products) == 0 {
		return map[string][]string{}
	}

	groupedProducts := make(map[string][]string)
	for _, p := range products {
		groupedProducts[p.Category] = append(groupedProducts[p.Category], p.Name)
	}
	return groupedProducts
}
