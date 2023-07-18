package main

import (
	"fmt"
	"sort"
)

/*
Задание:
	Реализовать бинарный поиск встроенными методами языка.
*/

func main() {
	sl := []int{0, 0, 1, 2, 2, 3, 4, 6, 7, 8, 9, 10, 20, 34}

	fmt.Println("Slice:", sl)

	x := 50
	i := sort.SearchInts(sl, x)
	fmt.Printf("Find (maybe find place for) %d at %d\n", x, i)
	// checking find - if i < len(sl) && sl[i] == x {...}

	// or

	i = sort.Search(len(sl), func(i int) bool {
		return sl[i] >= x
	})
	fmt.Printf("Find (maybe find place for) %d at %d\n", x, i)

	// or

	i, found := sort.Find(len(sl), func(i int) int {
		return x - sl[i]
	})
	if found {
		fmt.Printf("Find %d at %d\n", x, i)
	} else {
		fmt.Printf("Can't find %d\n", x)
	}
}
