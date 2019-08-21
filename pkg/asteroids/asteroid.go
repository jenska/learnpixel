package asteroids

import (
	"fmt"
	"math"
	"math/rand"

	"golang.org/x/image/colornames"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"

	"github.com/jenska/learnpixel/pkg/core"
)

const (
	MinRadius    = 20 // minimal size of an asteroid
	MaxRadius    = 60 // maximum size
	MaxMutation  = 5  // mutates a perfect circle
	NumVertices  = 20 // maximum vertices
	Split2Chunks = 3  // split into chunks
)

type Asteroid struct {
	center   pixel.Vec
	velocity pixel.Vec
	rotation float64
	radius   float64
	spin     float64
	shape    []pixel.Vec
	bounds   *pixel.Rect
}

func (a *Asteroid) Update(d float64) {
	core.Clip(&a.center, a.radius, a.bounds)
	a.center = a.center.Add(a.velocity.Scaled(1.0 + d))
	a.rotation += a.spin
}

func (a *Asteroid) Draw(imd *imdraw.IMDraw) {
	imd.Color = colornames.Mediumseagreen
	imd.SetMatrix(pixel.IM.Rotated(pixel.ZV, a.rotation).Moved(a.center))
	for _, v := range a.shape {
		imd.Push(v)
	}
	imd.Polygon(1)
}

func (a *Asteroid) IsSplittable() bool {
	return len(a.shape) >= 10
}

func (a *Asteroid) Split(p pixel.Vec) []Asteroid {
	// get the index of the nearest point
	index := 0
	min := math.MaxFloat64
	for i, s := range a.shape {
		if dist := core.Distance(s, p); min < dist {
			min = dist
			index = i
		}
	}

	chunks := make([]Asteroid, Split2Chunks)
	chunkSize := len(a.shape) / Split2Chunks
	for i := range chunks {
		chunks[i].shape = make([]pixel.Vec, chunkSize)
		copy(chunks[i].shape, a.shape[i*chunkSize:])
		chunks[i].shape = append(chunks[i].shape, a.shape[index])
	}
	return chunks
}

func (a *Asteroid) normalizeRotation() {
	for index := range a.shape {
		a.shape[index] = a.shape[index].Rotated(a.rotation)
	}
	a.rotation = 0
}

func NewAsteroid(bounds *pixel.Rect) *Asteroid {
	center := bounds.Center()
	angle := 2 * math.Pi * rand.Float64()
	spawnRadius := math.Max(bounds.W(), bounds.H()) + MaxRadius
	target := core.RandomVec(pixel.R(MaxRadius, MaxRadius, bounds.W()-MaxRadius, bounds.H()-MaxRadius))

	a := new(Asteroid)
	a.bounds = bounds
	a.radius = MinRadius + rand.Float64()*(MaxRadius-MinRadius)
	a.shape = mutateShape(makeCircle(a.radius, NumVertices), MaxMutation)

	a.center = pixel.Unit(angle).Scaled(spawnRadius).Add(center)
	a.velocity = pixel.Unit(core.Angle(a.center, target)).Scaled(0.5 + rand.Float64()*0.7)
	a.spin = (rand.Float64() - 0.5) * core.Rot2
	a.rotation = 0.0
	return a
}

func (a *Asteroid) String() string {
	return fmt.Sprintf("Asteroid(center: %v, velocity: %v, radius: %v)",
		a.center, a.velocity, a.radius)
}

func makeCircle(radius float64, segments int) []pixel.Vec {
	angular := 2.0 * math.Pi / float64(segments)
	circle := make([]pixel.Vec, segments)

	for index := range circle {
		circle[index] = pixel.Unit(float64(index) * angular).Scaled(radius)
	}
	return circle
}

func mutateShape(shape []pixel.Vec, max float64) []pixel.Vec {
	bounds := pixel.R(0, 0, max, max)
	average := pixel.V(0, 0)
	for i, point := range shape {
		shape[i] = point.Add(core.RandomVec(bounds))
		average = average.Add(point)
	}
	average = average.Scaled(1.0 / float64(len(shape)))
	for i, point := range shape {
		shape[i] = point.Sub(average)
	}
	return shape
}
