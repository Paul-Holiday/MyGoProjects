package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	groceries := make([]string, 0, 5)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\nВведите продукт (или 'готово' для завершения): ")
		product, _ := reader.ReadString('\n')
		product = strings.TrimSpace(product)
		if product == "готово" {
			break
		} else {
			groceries = append(groceries, product)
		}
	}

	fmt.Println(groceries)
}
