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

func RandomVec(b pixel.Rect) pixel.Vec {
	return pixel.V(
		rand.Float64()*(b.Max.X-b.Min.X)+b.Min.X,
		rand.Float64()*(b.Max.Y-b.Min.Y)+b.Min.Y)
}

func Clamp(v *pixel.Vec, bounds pixel.Rect) {
	v.X = pixel.Clamp(v.X, bounds.Min.X, bounds.Max.X)
	v.Y = pixel.Clamp(v.Y, bounds.Min.Y, bounds.Max.Y)
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

func Negate(v pixel.Vec) pixel.Vec {
	return pixel.V(-v.X, -v.Y)
}

func Magnitude(v pixel.Vec) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Distance(v, w pixel.Vec) float64 {
	return math.Sqrt((v.X-w.X)*(v.X-w.X) + (v.Y-w.Y)*(v.Y-w.Y))
}

func Angle(v, w pixel.Vec) float64 {
	diff := w.Sub(v)
	angle := math.Atan(diff.X / diff.Y)
	if diff.Y < 0.0 {
		angle += math.Pi
		if diff.Y > 0.0 {
			angle += math.Pi
		}
	} else if diff.Y < 0.0 {
		angle += math.Pi
	}
	return angle
}

func CalcRadius(shape []pixel.Vec) float64 {
	magnitude := 0.0
	for _, vertex := range shape {
		magnitude += Magnitude(vertex)
	}
	return magnitude / float64(len(shape))
}

func IndexOfFurthestPoint(shape []pixel.Vec, d pixel.Vec) int {
	maxProduct := 0.0
	index := 0
	for i, vertex := range shape {
		product := vertex.Dot(d)
		if product > maxProduct {
			maxProduct = product
			index = i
		}
	}
	return index
}

// This is to compute average center (roughly). It might be different from
// Center of Gravity, especially for bodies with nonuniform density,
// but this is ok as initial direction of simplex search in GJK.
func AveragePosition(shape []pixel.Vec) pixel.Vec {
	average := pixel.V(0, 0)
	for _, vertex := range shape {
		average = average.Add(vertex)
	}
	return average.Scaled(1.0 / float64(len(shape)))
}

// Minkowski sum support function for GJK
func minkowski(shape1, shape2 []pixel.Vec, d pixel.Vec) pixel.Vec {
	i := IndexOfFurthestPoint(shape1, d)
	j := IndexOfFurthestPoint(shape2, Negate(d))
	return shape1[i].Sub(shape2[j])
}

// Gilbert-Johnson-Keerthi (GJK) collision detection algorithm in 2D
func GJK(shape1, shape2 []pixel.Vec) bool {
	simplex := make([]pixel.Vec, 3)
	pos1 := AveragePosition(shape1)
	pos2 := AveragePosition(shape2)
	// initial direction from the center of 1st shape to the center of 2nd shape
	d := pos1.Sub(pos2)
	// if initial direction is zero – set it to any arbitrary axis (we choose X)}
	if d.X == 0 && d.Y == 0 {
		d.X = 1.0
	}
	// set the first support as initial point of the new simplex
	a := minkowski(shape1, shape2, d)
	if a.Dot(d) <= 0 {
		return false
	}

	d = Negate(a)
	simplex[0] = a
	for index := 0; ; {
		index++
		a = minkowski(shape1, shape2, d)
		if a.Dot(d) <= 0 {
			return false
		}
		simplex[index] = a

	}
}
