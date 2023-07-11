package main

import (
	"fmt"

	"l1_wb/task_11/generic_set"
)

/*
Задание:
	Реализовать пересечение двух неупорядоченных множеств.
*/

func main() {
	//also we can use bitset for int sets

	s1 := generic_set.New(1, 2, 3, 4, 5)
	s2 := generic_set.New(3, 4, 5, 6, 7, 8)

	intersect := generic_set.Intersect(s1, s2)

	fmt.Println(intersect.Slice())

	s3 := generic_set.New("cat", "dog", "mouse")
	s4 := generic_set.New("mouse", "house", "shmouse")

	intersect2 := generic_set.Intersect(s3, s4)

	fmt.Println(intersect2.Slice())
}
