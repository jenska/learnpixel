package core

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/faiface/pixel"
)

func TestClamp(t *testing.T) {
	v := pixel.V(0, 1000)
	b := pixel.R(0, 0, 100, 100)
	Clamp(&v, b)
	assert.Equal(t, 100.0, v.Y)
	assert.Equal(t, 0.0, v.X)
}
