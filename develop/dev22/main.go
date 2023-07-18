package main

import (
	"fmt"

	"l1_wb/develop/dev22/long"
)

/*
Задание:
	Разработать программу, которая перемножает, делит, складывает, вычитает две
	числовых переменных a,b, значение которых > 2^20.
*/

func main() {
	a := long.NewFromInt(2)
	b := long.NewFromInt(144)
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	fmt.Println("Sum:", long.Sum(a, b))
	fmt.Println("Sub:", long.Sub(b, a))
	fmt.Println("Mul:", long.Mul(a, b))
	q, d := long.Div(b, a)
	fmt.Println("Div:", q, d)

	fmt.Println()

	a, _ = long.NewFromStr("111222333444555666")
	b, _ = long.NewFromStr("999888777666555444")
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	fmt.Println("Sum:", long.Sum(a, b))
	fmt.Println("Sub:", long.Sub(b, a))
	fmt.Println("Mul:", long.Mul(a, b))
	q, d = long.Div(b, a)
	fmt.Println("Div:", q, d)
}
