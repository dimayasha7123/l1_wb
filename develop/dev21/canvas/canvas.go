package canvas

import (
	"math"

	"l1_wb/develop/dev21/interfaces"
)

// имеет начальную точку в левом верхнем углу
type Canvas struct {
	W, H float64
}

func (c Canvas) Distance(shape interfaces.Shape) float64 {
	x, y := shape.Centre()
	return math.Sqrt(x*x + y*y)
}
