package geometry

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPointsDistance(t *testing.T) {
	tests := []struct {
		name   string
		p1, p2 point
		want   float64
	}{
		{
			name: "egypt",
			p1: point{
				x: 0,
				y: 3,
			},
			p2: point{
				x: 4,
				y: 0,
			},
			want: 5,
		},
		{
			name: "zero",
			p1: point{
				x: 1,
				y: 1,
			},
			p2: point{
				x: 1,
				y: 1,
			},
			want: 0,
		},
		{
			name: "on line",
			p1: point{
				x: 1,
				y: 1,
			},
			p2: point{
				x: 10,
				y: 1,
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PointsDistance(tt.p1, tt.p2)

			assert.InDelta(t, tt.want, got, math.Pow(10, -4))
		})
	}
}
