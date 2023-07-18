package interfaces

type CanvasInterface interface {
	Distance(Shape) float64
}

type Shape interface {
	Centre() (float64, float64)
}
