package main

import (
	"fmt"
)

func main() {
	str := "рраббота"
	index := FirstNonRepeatingChar(str)
	if index == nil {
		fmt.Println(index)
	} else {
		fmt.Println(string(*index))
	}
}

// функция получает на вход строку и возвращает указатель на первый не повторяющийся символ,
// если не находит, либо строка пуста, то возвращает пустой указатель
func FirstNonRepeatingChar(s string) *rune {

	if s == "" {
		return nil
	}

	charMap := make(map[rune]int)

	for _, ch := range s {
		charMap[ch]++
	}

	for _, ch := range s {
		if charMap[ch] == 1 {
			result := ch
			return &result
		}
	}

	return nil
}
