package main

import (
	"fmt"

	"l1_wb/develop/dev12/string_set"
)

/*
Задание:
	Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее
	собственное множество.
*/

func main() {
	init := []string{"cat", "cat", "dog", "cat", "tree"}
	ss := string_set.New(init...)
	ss.Add("cat")
	ss.Add("mouse")
	fmt.Println(ss.Slice())
}
