package main

import (
	"image/color"
	"math"
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

const (
	start   = 10 // initial number of asteroids
	explode = 3  // exploded asteroid splits into parts

	width  = 1024
	height = 768

	rot1 = 1.0 * math.Pi / 180
	rot2 = 2.0 * math.Pi / 180
)

type (
	asteroid struct {
		center   pixel.Vec
		radius   float64
		velocity pixel.Vec

		vertices *[]pixel.Vec
		angle    float64
		color    color.Color
	}

	spaceship struct {
		center   pixel.Vec
		velocity pixel.Vec
		vertices []pixel.Vec

		angle2 float64
		color  color.Color
	}
)

var (
	bounds pixel.Rect

	asteroids []*asteroid   // asteroid field
	vertices  [][]pixel.Vec // pre-calculated vertices for asteroids

	ship spaceship
)

func (a *asteroid) update(d float64) {
	clip(&a.center, a.radius)
	a.center = a.center.Add(a.velocity.Scaled(1.0 + d))
}

func (a *asteroid) draw(imd *imdraw.IMDraw) {
	imd.Color = a.color
	imd.SetMatrix(pixel.IM.Rotated(pixel.ZV, a.angle).Moved(a.center))
	for _, v := range *a.vertices {
		imd.Push(v.Scaled(a.radius))
		a.angle += rot1
	}
	imd.Polygon(1)
}

func newAsteroid(numVertices int) *asteroid {
	var result asteroid
	result.center = randomVec(bounds)
	result.radius = float64(numVertices * 2)
	result.velocity = randomVec(pixel.R(-3, -3, 3, 3))

	result.vertices = &vertices[numVertices-5]
	result.color = pixel.RGB(0, 1, 0)

	return &result
}

func (s *spaceship) update(d float64) {
	clip(&s.center, 20)
	s.center = s.center.Add(s.velocity.Scaled(1.0 + d))
}

func (s *spaceship) draw(imd *imdraw.IMDraw) {
	imd.Color = s.color
	imd.SetMatrix(pixel.IM.Rotated(pixel.ZV, s.angle2).Moved(s.center))
	for _, v := range s.vertices {
		imd.Push(v)
	}
	imd.Polygon(1)
}

func randomVec(b pixel.Rect) pixel.Vec {
	return pixel.V(
		rand.Float64()*(b.Max.X-b.Min.X)+b.Min.X,
		rand.Float64()*(b.Max.Y-b.Min.Y)+b.Min.Y)
}

func clip(v *pixel.Vec, d float64) {
	switch {
	case v.X+d < 0:
		v.X = width - d
	case v.X-d > width:
		v.X = 0
	case v.Y+d < 0:
		v.Y = height - d
	case v.Y-d > height:
		v.Y = 0
	}
}

func run() {
	win, err := pixelgl.NewWindow(pixelgl.WindowConfig{
		Title:  "Asteroids",
		Bounds: bounds,
		VSync:  true,
	})
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)
	imd.Precision = 7

	last := time.Now()
	for !win.Closed() {
		win.SetClosed(win.JustPressed(pixelgl.KeyEscape) || win.JustPressed(pixelgl.KeyQ))

		d := time.Since(last).Seconds()
		last = time.Now()

		imd.Clear()
		for _, a := range asteroids {
			a.update(d)
			a.draw(imd)
		}
		switch {
		case win.Pressed(pixelgl.KeyLeft):
			ship.angle2 += rot2
		case win.Pressed(pixelgl.KeyRight):
			ship.angle2 -= rot2
		case win.Pressed(pixelgl.KeyUp):
			ship.velocity = ship.velocity.Add(pixel.V(0, 0.05).Rotated(ship.angle2))
		case win.Pressed(pixelgl.KeyDown):
			ship.velocity = ship.velocity.Sub(pixel.V(0, 0.05).Rotated(ship.angle2))
		}
		ship.update(d)
		ship.draw(imd)

		win.Clear(color.Black)
		imd.Draw(win)
		win.Update()
	}
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	bounds = pixel.R(0, 0, width, height)

	vertices = make([][]pixel.Vec, 5)
	vertex := pixel.V(1, 0)
	for i := range vertices {
		alpha := 2.0 * math.Pi / float64(i+5)
		vertices[i] = make([]pixel.Vec, i+5)
		for j := range vertices[i] {
			vertices[i][j] = vertex.Rotated(alpha * float64(j))
		}
	}

	asteroids = make([]*asteroid, start)
	for i := range asteroids {
		asteroids[i] = newAsteroid(rand.Intn(5) + 5)
	}
	ship.color = pixel.RGB(1, 1, 1)
	ship.center = bounds.Center()
	ship.vertices = []pixel.Vec{{-10, -10}, {0, 10}, {10, -10}}
}

func main() {
	pixelgl.Run(run)
}
