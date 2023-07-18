package dot

type Dot struct {
	X, Y float64
}

func (d Dot) Centre() (float64, float64) {
	return d.X, d.Y
}
