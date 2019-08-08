package asteroids

import (
	"golang.org/x/image/colornames"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"

	"github.com/jenska/learnpixel/pkg/core"
)

const shipSize = 10 // size of space ship in pixels

var delta = pixel.V(0, 0.05) // delta for velocity

type Spaceship struct {
	center   pixel.Vec
	velocity pixel.Vec
	vertices *[]pixel.Vec
	angle    float64
}

func (s *Spaceship) Update(d float64, bounds *pixel.Rect) {
	core.Clip(&s.center, shipSize, bounds)
	s.center = s.center.Add(s.velocity.Scaled(1.0 + d))
}

func (s *Spaceship) Draw(imd *imdraw.IMDraw) {
	imd.Color = colornames.White
	imd.SetMatrix(pixel.IM.Rotated(pixel.ZV, s.angle).Moved(s.center))
	for _, v := range *s.vertices {
		imd.Push(v)
	}
	imd.Polygon(1)
}

func (s *Spaceship) RotateLeft() {
	s.angle += core.Rot2
}

func (s *Spaceship) RotateRight() {
	s.angle -= core.Rot2
}

func (s *Spaceship) Forward() {
	s.velocity = s.velocity.Add(delta.Rotated(s.angle))
}

func (s *Spaceship) Backward() {
	s.velocity = s.velocity.Sub(delta.Rotated(s.angle))
}

func NewSpaceship(bounds *pixel.Rect) Spaceship {
	return Spaceship{
		center:   bounds.Center(),
		vertices: &[]pixel.Vec{{-shipSize, -shipSize}, {0, shipSize}, {shipSize, -shipSize}, {0, 0}},
	}
}
