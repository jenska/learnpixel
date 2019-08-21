package asteroids

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCircle(t *testing.T) {
	circle := makeCircle(20, 20)
	assert.Equal(t, 20, len(circle))
	for _, point := range circle {
		assert.Equal(t, 20.0*20.0, math.Round(point.X*point.X+point.Y*point.Y))
	}

	circle = makeCircle(30, 70)
	assert.Equal(t, 70, len(circle))
	for _, point := range circle {
		assert.Equal(t, 30.0*30.0, math.Round(point.X*point.X+point.Y*point.Y))
	}
}

func TestMutateCircle(t *testing.T) {
	circle := makeCircle(30, 30)

	for _, point := range circle {
		assert.True(t, 35.0*35.0 > point.X*point.X+point.Y*point.Y)
		assert.True(t, 25.0*25.0 < point.X*point.X+point.Y*point.Y)
	}

	shape := mutateShape(circle, 5.0)
	for _, point := range shape {
		assert.True(t, 35.0*35.0 > point.X*point.X+point.Y*point.Y)
		assert.True(t, 25.0*25.0 < point.X*point.X+point.Y*point.Y)
	}

}

func TestCalcRadius(t *testing.T) {
	circle := makeCircle(30, 30)
	radius := calcRadius(circle)
	assert.Equal(t, 30.0, radius)
}
