package core

import (
	"testing"

	"github.com/faiface/pixel"
	"github.com/stretchr/testify/assert"
)

// func Clip(v *pixel.Vec, d float64, bounds *pixel.Rect)
func TestClip(t *testing.T) {
	rect := pixel.R(10, 10, 100, 100)
	in := pixel.V(11, 11)
	out := pixel.V(91, 91)
	farOut := pixel.V(95, 95)

	Clip(&in, 1, &rect)
	assert.True(t, in.X == 11.0 && in.Y == 11.0, "point inside rect, in: "+in.String())
	Clip(&out, 1, &rect)
	assert.True(t, out.X == 91.0 && out.Y == 91.0, "point outside rect, but not far enough, out: "+out.String())
	Clip(&farOut, 1, &rect)
	assert.True(t, farOut.X == 0 && farOut.Y == 91.0, "point outside rect, farOut: "+farOut.String())

}
