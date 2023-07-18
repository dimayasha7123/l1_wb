package main

import (
	"fmt"
	"sync"

	"l1_wb/develop/dev18/counter"
)

/*
Задание:
	Реализовать структуру-счетчик, которая будет инкрементироваться в
	конкурентной среде. По завершению программа должна выводить итоговое значение счетчика.
*/

func main() {
	routines := 100
	incsPerRoutine := 1000

	fmt.Printf("Using %d goroutines\n", routines)
	fmt.Printf("Every goroutine will increment count %d times\n", incsPerRoutine)

	wg := &sync.WaitGroup{}
	wg.Add(routines)

	c := counter.New(0)

	for i := 0; i < routines; i++ {
		go func() {
			for i := 0; i < incsPerRoutine; i++ {
				c.Inc()
			}
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Printf("Count %d, must be %d\n", c.Value(), routines*incsPerRoutine)
	if c.Value() == uint64(routines*incsPerRoutine) {
		fmt.Println("Seems like everything ok!")
	} else {
		fmt.Println("I think we have problems...")
	}
}
