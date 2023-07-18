package main

import (
	"fmt"
)

/*
Задание:
	Написать программу, которая конкурентно рассчитает значение квадратов чисел
	взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

func main() {
	nums := getNums(5)

	done := make(chan struct{})

	for _, num := range nums {
		go printSqr(num, done)
	}

	for i := 0; i < len(nums); i++ {
		<-done
	}
}

func printSqr(num int, done chan<- struct{}) {
	fmt.Println(num * num)
	done <- struct{}{}
}

func getNums(count int) []int {
	ret := make([]int, 0, count)
	for i := 0; i < count; i++ {
		ret = append(ret, 2+i*2)
	}
	return ret
}
