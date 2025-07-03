package main

import "fmt"

func main() {
	fmt.Println("HW!")
	var m = map[string]string{
		"Animal 1": "Tiger",
		"Animal 2": "Lion",
	}
	m["Animal 3"] = "Giraffe"

	fmt.Println(m)
}
