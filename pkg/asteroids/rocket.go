package asteroids

import (
	"github.com/faiface/pixel"

	"github.com/jenska/learnpixel/pkg/core"
)

const (
	rocketSize   = 4   // size of rocket in pixels
	rocketEnergy = 100 // fuel of rocket
)

type Rocket struct {
	position pixel.Vec
	velocity pixel.Vec
	energy   int

	bounds *pixel.Rect
}

func (r *Rocket) Update(d float64) {
	core.Clip(&r.position, rocketSize, r.bounds)
	r.position = r.position.Add(r.velocity.Scaled(1.0 + d))
	r.energy--
}
