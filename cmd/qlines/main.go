package main

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

const (
	nlines    = 100
	velocity  = 15
	winHeight = 768
	winWidth  = 1024
)

type (
	qline struct {
		line [2]pixel.Vec
		vel  [2]pixel.Vec
	}
)

func randomVec(b pixel.Rect) pixel.Vec {
	return pixel.V(
		rand.Float64()*(b.Max.X-b.Min.X)+b.Min.X,
		rand.Float64()*(b.Max.Y-b.Min.Y)+b.Min.Y)
}

func moveVec(p, v *pixel.Vec) {
	x, y := p.X, p.Y
	vx, vy := v.X, v.Y

	if x+vx < 0 || x+vx > winWidth {
		vx = -vx
	}
	if y+vy < 0 || y+vy > winHeight {
		vy = -vy
	}

	p.X, p.Y = x+vx, y+vy
	v.X, v.Y = vx, vy
}

func run() {
	bounds := pixel.R(0, 0, winWidth, winHeight)
	vel := pixel.R(-velocity, -velocity, velocity, velocity)
	cfg := pixelgl.WindowConfig{
		Title:  "QLines Demo",
		Bounds: bounds,
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)
	imd.Color = pixel.RGB(1, 0, 0)

	var qlines [nlines]qline
	top := &qlines[0]
	top.line[0] = randomVec(bounds)
	top.line[1] = randomVec(bounds)
	top.vel[0] = randomVec(vel)
	top.vel[1] = randomVec(vel)
	for !win.Closed() {
		win.SetClosed(win.JustPressed(pixelgl.KeyEscape))

		imd.Clear()

		copy(qlines[1:], qlines[0:nlines-2])
		moveVec(&top.line[0], &top.vel[0])
		moveVec(&top.line[1], &top.vel[1])
		for _, m := range qlines {
			imd.Push(m.line[0])
			imd.Push(m.line[1])
			imd.Line(1)
		}
		win.Clear(color.Black)
		imd.Draw(win)
		win.Update()
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	pixelgl.Run(run)
}
