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
}

func (r *Rocket) update(d float64, bounds *pixel.Rect) {
	core.Clip(&r.position, rocketSize, bounds)
	r.position = r.position.Add(r.velocity.Scaled(1.0 + d))
	r.energy--
}


