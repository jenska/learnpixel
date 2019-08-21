package asteroids

import (
	"golang.org/x/image/colornames"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"

	"github.com/jenska/learnpixel/pkg/core"
)

const shipSize = 10 // size of space ship in pixels

type (
	Spaceship struct {
		center   pixel.Vec
		velocity pixel.Vec
		angle    float64
		event    events
		bounds   *pixel.Rect
	}

	events struct {
		rotateLeft  func() bool
		rotateRight func() bool
		forward     func() bool
		backward    func() bool
		shoot       func() bool
	}
)

var (
	delta    = pixel.V(0, 0.05) // delta for velocity
	vertices = []pixel.Vec{
		{-shipSize, -shipSize},
		{0, shipSize},
		{shipSize, -shipSize},
		{0, 0}}
)

func (s *Spaceship) Update(d float64) {
	switch {
	case s.event.shoot():
		// shoot
		fallthrough
	case s.event.rotateLeft():
		s.angle += core.Rot2
	case s.event.rotateRight():
		s.angle -= core.Rot2
	case s.event.forward():
		s.velocity = s.velocity.Add(delta.Rotated(s.angle))
	case s.event.backward():
		s.velocity = s.velocity.Sub(delta.Rotated(s.angle))
	}
	core.Clip(&s.center, shipSize, s.bounds)
	s.center = s.center.Add(s.velocity.Scaled(1.0 + d))
}

func (s *Spaceship) Draw(imd *imdraw.IMDraw) {
	imd.Color = colornames.White
	imd.SetMatrix(pixel.IM.Rotated(pixel.ZV, s.angle).Moved(s.center))
	for _, v := range vertices {
		imd.Push(v)
	}
	imd.Polygon(1)
}

func NewSpaceship(b *pixel.Rect,
	left, right, forward, backward, shoot func() bool) Spaceship {
	e := events{
		rotateLeft:  left,
		rotateRight: right,
		forward:     forward,
		backward:    backward,
		shoot:       shoot,
	}

	return Spaceship{
		bounds: b,
		center: b.Center(),
		event:  e,
	}
}
