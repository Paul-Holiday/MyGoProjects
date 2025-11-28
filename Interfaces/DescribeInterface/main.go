package main

import "fmt"

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

func DescribeNew(i interface{}) string {
	return fmt.Sprintf("%T", i)
}

func main() {
	fmt.Println("==Вывод типа значения с помощью пустого интерфейса и TypeSwitch (сложно и плохо)==")
	a := 45
	b := "Grass"
	c := 't'
	d := complex(13, -42)
	e := true
	f := map[string]int{}
	g := &a

	values := []interface{}{a, b, c, d, e, f, g}

	for _, v := range values {
		fmt.Println(Describe(v))
	}
	fmt.Println("==Вывод типа значения с помощью спецификатора %Т (легко и просто)==")

	for _, v := range values {
		fmt.Println(DescribeNew(v))
	}
}
