package core

import (
	"math"
	"math/rand"

	"github.com/faiface/pixel"
)

const (
	Rot1 = 1.0 * math.Pi / 180
	Rot2 = 2.0 * math.Pi / 180
)

func RandomVec(b *pixel.Rect) pixel.Vec {
	return pixel.V(
		rand.Float64()*(b.Max.X-b.Min.X)+b.Min.X,
		rand.Float64()*(b.Max.Y-b.Min.Y)+b.Min.Y)
}


func Clip(v *pixel.Vec, d float64, bounds *pixel.Rect) {
	switch {
	case v.X+d < 0:
		v.X = bounds.W() - d
	case v.X-d > bounds.W():
		v.X = 0
	case v.Y+d < 0:
		v.Y = bounds.H() - d
	case v.Y-d > bounds.H():
		v.Y = 0
	}
}
