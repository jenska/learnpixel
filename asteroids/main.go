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
	width   = 1024
	height  = 768
)

type (
	asteroid struct {
		center   pixel.Vec
		radius   float64
		velocity pixel.Vec

		vertices *[]pixel.Vec
		angle    int
		color    color.Color
	}
)

var (
	bounds    pixel.Rect
	asteroids []*asteroid
	vertices  [][]pixel.Vec
)

func (a *asteroid) update(d float64) {
	// clip
	switch {
	case a.center.X+a.radius < 0:
		a.center.X = width - a.radius
	case a.center.X-a.radius > width:
		a.center.X = 0
	case a.center.Y+a.radius < 0:
		a.center.Y = height - a.radius
	case a.center.Y-a.radius > height:
		a.center.Y = 0
	}
	// move
	a.center = a.center.Add(a.velocity.Scaled(1.0 + d))
}

func (a *asteroid) draw(imd *imdraw.IMDraw) {
	imd.Color = a.color
	imd.SetMatrix(pixel.IM.Rotated(pixel.ZV, float64(a.angle)*math.Pi/180.0).Moved(a.center))
	for _, v := range *a.vertices {
		imd.Push(v.Scaled(a.radius))
		a.angle = (a.angle + 1) % 360
	}
	imd.Polygon(1)
}

func randomVec(b pixel.Rect) pixel.Vec {
	return pixel.V(
		rand.Float64()*(b.Max.X-b.Min.X)+b.Min.X,
		rand.Float64()*(b.Max.Y-b.Min.Y)+b.Min.Y)
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

}

func main() {
	pixelgl.Run(run)
}
