package main

import (
	"fmt"
	"strings"
)

/*
Задание:
	Разработать программу, которая переворачивает слова в строке.
	Пример: «snow dog sun — sun dog snow».
*/

func main() {
	s := "snow dog sun"
	r := reverseWordsInString(s)
	fmt.Printf("reverse words in %q is %q\n", s, r)
}

func reverseWordsInString(s string) string {
	sep := " "
	words := strings.Split(s, sep)
	sb := strings.Builder{}
	sb.Grow(len(s))
	for i := len(words) - 1; i >= 0; i-- {
		if i != len(words)-1 {
			sb.WriteString(sep)
		}
		sb.WriteString(words[i])
	}
	return sb.String()
}
