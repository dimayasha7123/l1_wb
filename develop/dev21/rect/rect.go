package rect

import (
	"math"
)

// здесь же начальная точка располагается в центре
type Rect struct {
	W, H float64
}

func (r Rect) DistanceTo(x, y float64) float64 {
	X := x - r.W/2
	Y := y - r.H/2
	return math.Sqrt(X*X + Y*Y)
}
