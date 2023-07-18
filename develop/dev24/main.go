package main

import (
	"fmt"

	"l1_wb/develop/dev24/geometry"
)

/*
Задание:
	Разработать программу нахождения расстояния между двумя точками, которые
	представлены в виде структуры Point с инкапсулированными параметрами x,y и
	конструктором.
*/

func main() {
	p1 := geometry.NewPoint(3, 4)
	p2 := geometry.NewPoint(6, 0)

	fmt.Println("Distance func:", geometry.PointsDistance(p1, p2))
	fmt.Println("Distance method:", p1.DistanceTo(p2))
}
