package main

import (
	"fmt"
)

/*
Задание:
	Дана переменная int64. Разработать программу, которая устанавливает i-й бит в 1 или 0.
*/

func main() {
	x := int64(14)
	printX(x)

	setIthBitTo(&x, 0, true)
	printX(x)

	setIthBitTo(&x, 1, false)
	printX(x)
}

func printX(x int64) {
	fmt.Printf("x - %d, binary - %b\n", x, x)
}

func setIthBitTo(num *int64, i int, value bool) {
	if value {
		*num |= 1 << i
	} else {
		*num &^= 1 << i // or *num &= math.MaxInt64 - 1<<i
	}
}
