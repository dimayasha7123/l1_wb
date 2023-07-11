package main

import (
	"fmt"
	"unicode"
)

/*
Задание:
	Разработать программу, которая проверяет, что все символы в строке
	уникальные (true — если уникальные, false etc). Функция проверки должна быть
	регистронезависимой.

	Например:
	abcd — true
	abCdefAaf — false
	aabcd — false
*/

func main() {
	s := "abcd"
	fmt.Printf("%q consists of uniq symbols (case insensitive): %v\n", s, uniqSymbolsCaseInsensitive(s))
}

func uniqSymbolsCaseInsensitive(s string) bool {
	uniq := make(map[rune]struct{})
	for _, r := range s {
		lower := unicode.ToLower(r)
		_, ok := uniq[lower]
		if ok {
			return false
		}
		uniq[lower] = struct{}{}
	}
	return true
}
