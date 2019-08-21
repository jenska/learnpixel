package main

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/jenska/learnpixel/pkg/core"
)

const (
	nlines    = 100
	velocity  = 15
	winHeight = 768
	winWidth  = 1024
)

type qline struct {
	line [2]pixel.Vec
	vel  [2]pixel.Vec
}

func moveVec(p, v *pixel.Vec, b pixel.Rect) {
	x, y := p.X, p.Y
	vx, vy := v.X, v.Y

	if x+vx < 0 || x+vx > b.W() {
		vx = -vx
	}
	if y+vy < 0 || y+vy > b.H() {
		vy = -vy
	}

	p.X, p.Y = x+vx, y+vy
	v.X, v.Y = vx, vy
	core.Clamp(p, b)
}

func run() {
	bounds := pixel.R(0, 0, winWidth, winHeight)
	vel := pixel.R(-velocity, -velocity, velocity, velocity)
	cfg := pixelgl.WindowConfig{
		Title:     "QLines Demo",
		Bounds:    bounds,
		Resizable: true,
		VSync:     true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)
	imd.Color = pixel.RGB(1, 0, 0)

	var qlines [nlines]qline
	top := &qlines[0]
	top.line[0] = core.RandomVec(bounds)
	top.line[1] = core.RandomVec(bounds)
	top.vel[0] = core.RandomVec(vel)
	top.vel[1] = core.RandomVec(vel)
	for !win.Closed() {
		win.SetClosed(win.JustPressed(pixelgl.KeyEscape))

		imd.Clear()

		copy(qlines[1:], qlines[0:nlines-2])
		moveVec(&top.line[0], &top.vel[0], bounds)
		moveVec(&top.line[1], &top.vel[1], bounds)
		for _, m := range qlines {
			imd.Push(m.line[0])
			imd.Push(m.line[1])
			imd.Line(1)
		}
		win.Clear(color.Black)
		imd.Draw(win)
		win.Update()
		bounds = win.Bounds()
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	pixelgl.Run(run)
}
