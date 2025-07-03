package main

import "fmt"

func main() {
	words := []string{"apple", "cherry", "orange", "pear", "pear", "apple", "pear", "cherry", "pear", "apple", "orange"}
	wordsFreq := CountWords(words)
	fmt.Println(wordsFreq)

	fruit := "pineapple"
	fmt.Println(CharFreq(fruit))
}

func CountWords(words []string) map[string]int {
	wordsFreq := make(map[string]int)

	for _, word := range words {
		wordsFreq[word] += 1
	}

	return wordsFreq
}

func CharFreq(word string) map[string]int {
	charFreq := make(map[string]int)

	for _, char := range word {
		charFreq[string(char)]++
	}

	return charFreq
}

// // частота каждого из слов
// func CountWords(words []string) map[string]int {
// 	wordsMap := make(map[string]int)
// 	for _, v := range words {
// 		wordsMap[v]++
// 	}
// 	return wordsMap
// }

// // частота символов в строке
// func CharFreq(str string) map[string]int {
// 	chFreq := make(map[string]int)

// 	for _, r := range str {
// 		chFreq[string(r)]++
// 	}
// 	return chFreq
// }
