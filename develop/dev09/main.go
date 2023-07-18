package main

import (
	"fmt"
	"sync"
)

/*
Задание:
	Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из
	массива, во второй — результат операции x*2, после чего данные из второго
	канала должны выводиться в stdout.
*/

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	first := make(chan int)
	second := make(chan int)

	wg := &sync.WaitGroup{}
	wg.Add(3)

	go send(a, first, wg)
	go doWork(first, second, func(i int) int { return 2 * i }, wg)
	go show(second, wg)

	wg.Wait()
}

func send(a []int, ch chan<- int, wg *sync.WaitGroup) {
	for _, num := range a {
		fmt.Printf("Sending %d next\n", num)
		ch <- num
	}

	close(ch)
	fmt.Println("Quit from send")
	wg.Done()
}

func doWork(from <-chan int, to chan<- int, work func(int) int, wg *sync.WaitGroup) {
	for num := range from {
		done := work(num)
		fmt.Printf("Done work on %d, sending %d next\n", num, done)
		to <- done
	}

	close(to)
	fmt.Println("Quit from doWork")
	wg.Done()
}

func show(ch <-chan int, wg *sync.WaitGroup) {
	for num := range ch {
		fmt.Printf("Get %d\n", num)
	}

	fmt.Println("Quit from show")
	wg.Done()
}
