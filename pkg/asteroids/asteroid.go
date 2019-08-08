package asteroids

import (
	"math"

	"golang.org/x/image/colornames"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"

	"github.com/jenska/learnpixel/pkg/core"
)

var (
	vertices [][]pixel.Vec // pre-calculated vertices for asteroids
	velocityRange = pixel.R(-3, -3, 3, 3)
)

type Asteroid struct {
	center   pixel.Vec
	velocity pixel.Vec
	radius   float64
	vertices *[]pixel.Vec
	angle    float64
}

func (a *Asteroid) Update(d float64, bounds *pixel.Rect) {
	core.Clip(&a.center, a.radius, bounds)
	a.center = a.center.Add(a.velocity.Scaled(1.0 + d))
}

func (a *Asteroid) Draw(imd *imdraw.IMDraw) {
	imd.Color = colornames.Mediumseagreen
	imd.SetMatrix(pixel.IM.Rotated(pixel.ZV, a.angle).Moved(a.center))
	for _, v := range *a.vertices {
		imd.Push(v.Scaled(a.radius))
		a.angle += core.Rot1
	}
	imd.Polygon(1)
}

func NewAsteroid(numVertices int, bounds *pixel.Rect) *Asteroid {
	return &Asteroid{
		center:   core.RandomVec(bounds),
		radius:   float64(numVertices * 2),
		velocity: core.RandomVec(&velocityRange),
		vertices: &vertices[numVertices-5],
	}
}

func init() {
	vertices = make([][]pixel.Vec, 5)
	vertex := pixel.V(1, 0)
	for i := range vertices {
		alpha := 2.0 * math.Pi / float64(i+5)
		vertices[i] = make([]pixel.Vec, i+5)
		for j := range vertices[i] {
			vertices[i][j] = vertex.Rotated(alpha * float64(j))
			
		}
	}
}