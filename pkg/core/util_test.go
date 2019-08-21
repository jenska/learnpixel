package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntZip(t *testing.T) {
	a := []int{1, 3, 5, 7, 9}
	b := []int{2, 4, 6, 8, 10}

	c, err := IntZip(a, b)
	assert.Nil(t, err)
	assert.Equal(t, 10, len(c))
	for i := 1; i <= 10; i++ {
		assert.Equal(t, i, c[i-1])
	}
}

func TestZip(t *testing.T) {
	a := []int{1, 3, 5, 7, 9}
	b := []int{2, 4, 6, 8, 10}

	a1 := make([]interface{}, len(a))
	for i, v := range a {
		a1[i] = v
	}

	b1 := make([]interface{}, len(b))
	for i, v := range b {
		b1[i] = v
	}

	c, err := Zip(a1, b1)
	assert.Nil(t, err)
	assert.Equal(t, 10, len(c))
	for i := 1; i <= 10; i++ {
		assert.Equal(t, i, c[i-1])
	}
}
