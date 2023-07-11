package main

import (
	"fmt"
	"sort"
)

/*
Задание:
	Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/

func main() {
	a := data()
	fmt.Println("Initial:", a)

	a = data()
	sort.Ints(a)
	fmt.Println("sort.Ints:", a)

	a = data()
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	fmt.Println("sort.Slice:", a)
	// SliceStable works as well

	a = data()
	sort.Sort(Desc(a))
	fmt.Println("sort.Sort (descending):", a)
	// sort.Stable works as well
}

func data() []int {
	return []int{6, 8, 7, 4, 3, 5, 1, 2, 9, 0}
}

// may use sort.Reverse
type Desc []int

func (d Desc) Len() int           { return len(d) }
func (d Desc) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }
func (d Desc) Less(i, j int) bool { return d[i] > d[j] }
