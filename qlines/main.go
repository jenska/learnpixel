package main

// update pixel.Line

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

const (
	qlines  = 100
	vel     = 20
	wheight = 768
	wwidth  = 1024
)

type (
	mline struct {
		line [2]pixel.Vec
		vel  [2]pixel.Vec
	}
)

func randomV(b pixel.Rect) pixel.Vec {
	return pixel.V(
		rand.Float64()*(b.Max.X-b.Min.X)+b.Min.X,
		rand.Float64()*(b.Max.Y-b.Min.Y)+b.Min.Y)
}

func moveV(p, v *pixel.Vec) {
	x, y := p.X, p.Y
	vx, vy := v.X, v.Y

	if x+vx < 0 || x+vx > wwidth {
		vx = -vx
	}
	if y+vy < 0 || y+vy > wheight {
		vy = -vy
	}

	p.X, p.Y = x+vx, y+vy
	v.X, v.Y = vx, vy
}

func run() {
	bounds := pixel.R(0, 0, wwidth, wheight)
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

	var mlines [qlines]mline
	top := &mlines[0]
	top.line[0] = randomV(bounds)
	top.line[1] = randomV(bounds)
	top.vel[0] = randomV(pixel.R(-vel/2, -vel/2, vel/2, vel/2))
	top.vel[1] = randomV(pixel.R(-vel/2, -vel/2, vel/2, vel/2))
	for !win.Closed() {
		win.SetClosed(win.JustPressed(pixelgl.KeyEscape))

		imd.Clear()

		copy(mlines[1:], mlines[0:qlines-2])
		moveV(&top.line[0], &top.vel[0])
		moveV(&top.line[1], &top.vel[1])
		for _, m := range mlines {
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
