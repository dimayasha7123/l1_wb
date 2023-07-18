package main

import "fmt"

/*
Задание:
	Разработать программу, которая переворачивает подаваемую на ход строку
	(например: «главрыба — абырвалг»). Символы могут быть unicode.
*/

func main() {
	s := "😼AÅÄÖ\u20A9"
	r := reverseString(s)
	fmt.Printf("reverse string for %s is %s\n", s, r)
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
