package rect

import (
	"math"

	"l1_wb/task_21/interfaces"
)

type CanvasAdapter struct {
	Rect
}

func (ca CanvasAdapter) Distance(shape interfaces.Shape) float64 {
	// bad example...
	x, y := shape.Centre()
	return math.Sqrt(x*x + y*y)
}
