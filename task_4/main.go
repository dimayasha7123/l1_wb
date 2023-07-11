package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

/*
Задание:
	Реализовать постоянную запись данных в канал (главный поток). Реализовать
	набор из N воркеров, которые читают произвольные данные из канала и
	выводят в stdout. Необходима возможность выбора количества воркеров при
	старте.
	Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать
	способ завершения работы всех воркеров.
*/

func main() {
	n := 10
	flag.IntVar(&n, "n", n, "count of workers")
	flag.Parse()

	data := make(chan string, n)
	prodDone := make(chan struct{})
	wg := &sync.WaitGroup{}

	go producer(data, prodDone)

	wg.Add(n)
	for i := 0; i < n; i++ {
		go consumer(i, data, wg)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	fmt.Println()

	prodDone <- struct{}{}
	close(data)

	wg.Wait()
}

func consumer(number int, data <-chan string, wg *sync.WaitGroup) {
	for val := range data {
		fmt.Printf("%d: %s\n", number, val)
	}
	// maybe we must do something else
	fmt.Printf("Consumer %d stop working\n", number)
	wg.Done()
}

func producer(data chan<- string, done <-chan struct{}) {
	strs := []string{"Tiger", "Puss", "Smokey", "Misty", "Tigger", "Kitty", "Oscar", "Missy", "Max", "Ginger"}
	for {
		select {
		case <-done:
			fmt.Println("Producer stop sending")
			return
		default:
			data <- strs[rand.Intn(len(strs))]
			time.Sleep(time.Millisecond * 100)
		}
	}
}
