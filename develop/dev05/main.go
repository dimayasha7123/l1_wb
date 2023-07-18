package main

import (
	"flag"
	"fmt"
	"time"
)

/*
Задание:
	Разработать программу, которая будет последовательно отправлять значения в
	канал, а с другой стороны канала — читать. По истечению N секунд программа
	должна завершаться.
*/

func main() {
	dur := time.Second
	flag.DurationVar(&dur, "dur", dur, "time for work")
	flag.Parse()

	data := make(chan int)
	prodDone := make(chan struct{})
	consDone := make(chan struct{})

	go producer(data, prodDone)
	go consumer(data, consDone)

	time.Sleep(dur)
	prodDone <- struct{}{}
	close(data)
	<-consDone
}

func producer(data chan<- int, done <-chan struct{}) {
	c := 0
	for {
		select {
		case <-done:
			fmt.Println("Stop sending data")
			return
		default:
			fmt.Printf("Sending %d\n", c)
			data <- c
			c++
			time.Sleep(time.Millisecond * 100)
		}
	}
}

func consumer(data <-chan int, done chan<- struct{}) {
	for c := range data {
		fmt.Printf("Get %d\n", c)
	}
	fmt.Println("Stop getting data")
	done <- struct{}{}
}
