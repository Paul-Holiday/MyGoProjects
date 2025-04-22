// кажется, я схватился за слишком сложную задачу, пока что не хватает навыков чтобы решить, вернусь к задачке позже

package main

import (
	"fmt"
)

// strings.HasPrefix("hello", "he") // true
// TrimLeft(s, cutset string) string – удаляет символы cutset только с начала строки.

// M C X I могут повторяться 1-3 раза
// V L D не могут повторяться
func romanToInt(s string) int {
	s = Reverse(s)
	RomanDigits := "IVXLCDM"
	runeRef := []rune(RomanDigits)
	runeStr := []rune(s)
	var count int = 0
	for i := 0; i < len(RomanDigits); i++ {
		if runeStr[i] == runeRef[i] {
			if len(runeStr) > 1 {
				if runeStr[i+1] == runeRef[i] {
					if len(runeStr) > 2 {
						if runeStr[i+2] == runeRef[i] {
							count += 3
							runeStr = runeStr[3:]
							continue
						}
					}
				}
				count += 2
				continue
			}
			count++
			continue
		}
	}
	fmt.Println(count)
	return count
}

// разворачиваю строку так как пришла в голову идея читать строку с конца, понимаю что это однозначно можно сделать проще, но ради практики решил сделать, как итог узнал как работают строки в Go и как можно манипулировать ими посимвольно через руны
func Reverse(s string) string {
	runes := []rune(s)
	//fmt.Println(runes)
	for i := 0; i < len(s)/2; i++ {
		runes[i], runes[len(s)-1-i] = runes[len(s)-1-i], runes[i]
	}
	// fmt.Println(runes)
	// fmt.Println(string(runes))
	return string(runes)
}

func main() {
	fmt.Println(romanToInt("III"))
}

// MCMXCIV
