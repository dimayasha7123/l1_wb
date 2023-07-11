package main

import (
	"fmt"
	"os/exec"
	"time"
)

/*
Задание:
	Реализовать собственную функцию sleep.
*/

// cat test with assert.WithinDuration
func main() {
	funcs := map[string]sleepFunc{
		"embedded":               time.Sleep,
		"using channel":          sleepChannel,
		"using other sleep func": sleepOther,
		"active waiting":         sleepActive,
	}

	dur := time.Second * 1
	tryings := 10
	fmt.Println("Test duration:", dur)
	fmt.Println("Test", tryings, "times")

	for name, f := range funcs {
		fmt.Println()
		fmt.Println("Testing sleep function:", name)
		avg := time.Duration(0)

		for t := 0; t < tryings; t++ {
			start := time.Now()

			f(dur)

			end := time.Since(start)
			fmt.Println(t+1, ":", end-dur)
			avg += end
		}
		avg /= time.Duration(tryings)

		fmt.Println("Actually average sleep for", avg)
		fmt.Println("Average time error (with measuring):", avg-dur)
	}
}

type sleepFunc func(time.Duration)

func sleepChannel(d time.Duration) {
	<-time.After(d)
}

func sleepOther(d time.Duration) {
	exec.Command("sleep", fmt.Sprint(d)).Run()
}

func sleepActive(d time.Duration) {
	now := time.Now()
	for time.Since(now) < d {
	}
}
