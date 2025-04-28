package main

import "fmt"

func main() {
	a := "кабан"
	b := "банка"
	fmt.Println(AreAnagrams(a, b))
}

func AreAnagrams(a, b string) bool { // решение через одну мапу, это эффективнее по памяти
	if len(a) != len(b) {
		return false
	}

	Freq := make(map[string]int)

	for _, ch := range a {
		Freq[string(ch)]++
	}

	for _, ch := range b {
		Freq[string(ch)]--
	}

	for k := range Freq {
		if Freq[k] != 0 {
			return false
		}
	}

	return true
}

/* это решение через две мапы, где я вычитаю частоты появления элементов в одной мапе из другой
func AreAnagrams(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	aFreq := make(map[string]int)
	bFreq := make(map[string]int)

	for _, i := range a {
		aFreq[string(i)]++
	}
	for _, j := range b {
		bFreq[string(j)]++
	}

	for g := range aFreq {
		_, exists := bFreq[g]
		if exists {
			aFreq[g] -= bFreq[g]
		} else {
			return false
		}
	}

	for h := range aFreq {
		if aFreq[h] != 0 {
			return false
		}
	}
	return true
}
*/
