package main

import "fmt"

/*
Задание:
	Поменять местами два числа без создания временной переменной.
*/

func main() {
	a := 3
	b := 5
	fmt.Println(a, b)

	a, b = b, a
	fmt.Println(a, b)

	a ^= b
	b ^= a
	a ^= b
	fmt.Println(a, b)

	a = a + b
	b = a - b
	a = a - b
	fmt.Println(a, b)

	a = a * b
	b = a / b
	a = a / b
	fmt.Println(a, b)
}
