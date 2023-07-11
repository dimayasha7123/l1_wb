package main

import (
	"errors"
	"fmt"
)

/*
Задание:
	Удалить i-ый элемент из слайса.
*/

func main() {
	sl := []int{1, 2, 3, 4, 5}
	fmt.Println("sl (before):", sl)

	i := 0
	del, err := deleteIth(sl, i)
	//del, err := deleteIthCopy(sl, i)

	fmt.Println("sl (after):", sl)
	fmt.Printf("with %d-th deleted: %v\n", i, del)
	fmt.Println("deleting error:", err)
}

var OutOfRangeErr = errors.New("index out of range")

func deleteIth(sl []int, i int) ([]int, error) {
	if i < 0 || i >= len(sl) {
		return sl, OutOfRangeErr
	}
	return append(sl[:i], sl[i+1:]...), nil
}

func deleteIthCopy(sl []int, i int) ([]int, error) {
	if i < 0 || i >= len(sl) {
		return sl, OutOfRangeErr
	}
	ret := make([]int, len(sl)-1)
	copy(ret, sl[:i])
	copy(ret[i:], sl[i+1:])
	return ret, nil
}
