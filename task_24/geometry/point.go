package geometry

import "math"

type point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) point {
	return point{x, y}
}

func PointsDistance(a, b point) float64 {
	return math.Sqrt((a.x-b.x)*(a.x-b.x) + (a.y-b.y)*(a.y-b.y))
}

func (p point) DistanceTo(d point) float64 {
	return PointsDistance(p, d)
}
