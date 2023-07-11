package main

import (
	"fmt"
	"sync"

	"l1_wb/task_7/sync_map"
)

/*
Задание:
	Реализовать конкурентную запись данных в map.
*/

func main() {
	routines := 100
	incsPerRoutine := 1000

	fmt.Printf("Using %d goroutines\n", routines)
	fmt.Printf("Every goroutine will increment one map value %d times\n", incsPerRoutine)

	wg := &sync.WaitGroup{}
	wg.Add(routines)

	sm := sync_map.New[int, int]()
	key := 42

	for i := 0; i < routines; i++ {
		go func() {
			for i := 0; i < incsPerRoutine; i++ {
				val, _ := sm.Get(key)
				sm.Set(key, val+1)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	val, _ := sm.Get(key)
	fmt.Printf("Value is %d\n", val)
	fmt.Println("It is ok, because we get value and set it in two steps...")
	fmt.Println("And is is ok, because we don't get panics like this:")
	fmt.Println("fatal error: concurrent map writes")
	fmt.Println("fatal error: concurrent map read and map write")
}
