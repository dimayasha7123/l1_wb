package main

import "strings"

/*
Задание:
	К каким негативным последствиям может привести данный фрагмент кода, и как
	это исправить? Приведите корректный пример реализации.
	var justString string
	func someFunc() {
		v := createHugeString(1 << 10)
		justString = v[:100]
	}
	func main() {
		someFunc()
	}
*/

/*
Во-первых, необходимо аккуратно делать срез, если размер строки, которую мы срезаем будет меньше,
чем правая граница среза, то получим panic: runtime error: slice bounds out of range [:100]

Во-вторых, у нас утекает память, так как внутри строка устроена почти также как слайс, то есть
имеется указатель на область в памяти, где лежат байтики самой строки, их мы выделили много,
а пользуемся малым количеством после среза, но вся исходная строка в памяти остается,
решается это, например, копированием:
*/

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	bytes := make([]byte, 100)
	copy(bytes, v)
	justString = string(bytes)
}

func main() {
	someFunc()
}

func createHugeString(size int) string {
	return strings.Repeat("A", size)
}
