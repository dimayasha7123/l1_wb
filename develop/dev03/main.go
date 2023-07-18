package main

import "fmt"

/*
Задание:
	Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
	квадратов(2^2+3^2+4^2...) с использованием конкурентных вычислений.
*/

func main() {
	nums := getNums(10)

	ansCh := make(chan int)
	go calcNumsSqrsSum(nums, ansCh)

	ans := <-ansCh
	fmt.Println(ans)
}

func calcNumsSqrsSum(nums []int, ansCh chan<- int) {
	sqrs := make(chan int, len(nums)/2)
	go makeSqrs(nums, sqrs)

	ans := 0
	for sqr := range sqrs {
		ans += sqr
	}
	ansCh <- ans
}

func makeSqrs(nums []int, sqrs chan<- int) {
	for _, num := range nums {
		sqrs <- num * num
	}
	close(sqrs)
}

func getNums(count int) []int {
	ret := make([]int, 0, count)
	for i := 0; i < count; i++ {
		ret = append(ret, 2+i*2)
	}
	return ret
}
