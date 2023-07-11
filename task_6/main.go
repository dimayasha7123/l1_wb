package main

import (
	"context"
	"fmt"
	"time"
)

/*
Задание:
	Реализовать все возможные способы остановки выполнения горутины.
*/

func main() {
	first()
	second()
	third()
	forth()
}

func first() {
	fmt.Println("\nFirst method: check channel with \"ok\"")

	ch := make(chan string)
	go func() {
		for {
			v, ok := <-ch
			if !ok {
				fmt.Println("Channel closed, finish")
				return
			}
			fmt.Println(v)
		}
	}()

	ch <- "Hello from go"
	close(ch)

	time.Sleep(time.Second)
}

func second() {
	fmt.Println("\nSecond method: using for range on channel")

	ch := make(chan string)
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
		fmt.Println("Channel closed, finish")
	}()

	ch <- "Hello from go"
	close(ch)

	time.Sleep(time.Second)
}

func third() {
	fmt.Println("\nThird method: using done channel")

	done := make(chan struct{})
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Get done signal, finish")
				return
			default:
				fmt.Println("Do useful work...")
				time.Sleep(200 * time.Millisecond)
			}

		}
	}()

	time.Sleep(time.Second)
	done <- struct{}{}

	time.Sleep(time.Second)
}

func forth() {
	fmt.Println("\nForth method: using context.WithCancel")

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Context cancelled, finish")
				return
			default:
				fmt.Println("Do useful work...")
				time.Sleep(200 * time.Millisecond)
			}

		}
	}()

	time.Sleep(time.Second)
	cancel()

	time.Sleep(time.Second)
}

// also can use global flag or sync.Mutex, but don't
