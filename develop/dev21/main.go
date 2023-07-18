package main

import (
	"fmt"

	"l1_wb/develop/dev21/canvas"
	"l1_wb/develop/dev21/dot"
	"l1_wb/develop/dev21/interfaces"
	"l1_wb/develop/dev21/rect"
)

/*
Задание:
	Реализовать паттерн «адаптер» на любом примере.
*/

/*
	Допустим, имеются объекты двух типов: Canvas и Rect, а также имеется место в коде,
	где мы захотели использовать Rect вместо Canvas. Однако они немного отличаются с точки зрения интерфейсов и
	логики работы: отличаются сигнатуры функций Distance и местоположение центра на этих объектах (центр у Canvas
	в левом верхнем углу, а у Rect по центру фигуры.
*/

func main() {
	var canva interfaces.CanvasInterface

	canva = canvas.Canvas{W: 10, H: 4}

	canva = rect.CanvasAdapter{
		Rect: rect.Rect{W: 10, H: 4},
	}

	var shp interfaces.Shape = dot.Dot{X: 4, Y: 3}

	fmt.Println(canva.Distance(shp))
}
