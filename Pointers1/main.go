package main

import "fmt"

func main() {
	i := 10
	var p *int
	p = &i
	fmt.Println(*p)
	*p = 20
	fmt.Println(i)
}
