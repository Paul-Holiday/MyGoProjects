package main

import "fmt"

type value interface{}

func Describe(i interface{}) string {
	switch i.(type) {
	case int, *int:
		return "int"
	case uint, *uint:
		return "uint"
	case string, *string:
		return "string"
	case float64, *float64:
		return "float64"
	case bool, *bool:
		return "bool"
	case rune, *rune:
		return "rune"
	case complex128, *complex128:
		return "complex128"
	default:
		return "unknown"
	}
}

func main() {
	fmt.Println("==Вывод типа значения с помощью пустого интерфейса и TypeSwitch==")
	a := 45
	b := "Grass"
	c := 't'
	d := complex(13, -42)
	e := true
	f := map[string]int{}

	values := []value{a, b, c, d, e, f}

	for _, v := range values {
		fmt.Println(Describe(v))
	}
}
