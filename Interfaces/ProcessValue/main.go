package main

import "fmt"

func main() {
	a := 12
	s := "Dear Granny!"
	fmt.Println(ProcessValue(a))
	fmt.Println(ProcessValue(s))
}

func ProcessValue(val interface{}) interface{} {
	switch v := val.(type) {
	case int:
		return v + 1
	case string:
		return v + " Thank You! "
	default:
		return v
	}
}
