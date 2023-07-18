package main

import (
	"fmt"
	"math"
	"sort"
)

/*
Задание:
	Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0,
	15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10
	градусов. Последовательность в подмножноствах не важна.
	Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/

func main() {
	wide := 10.0
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 30}

	tempSets := make(map[int][]float64)
	for _, t := range temps {
		ts := int(math.Round(math.Floor(t/wide) * wide))
		tempSets[ts] = append(tempSets[ts], t)
	}

	fmt.Printf("Next grouping for n multiples of %f - n: [n; n+10)\n", wide)
	keys := make([]int, 0, len(tempSets))
	for k := range tempSets {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Printf("%d: %v\n", k, tempSets[k])
	}
}
