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

func NewMline(b pixel.Rect) mline {
	var result mline
	result.line[0] = randomV(b)
	result.line[1] = randomV(b)
	result.vel[0] = randomV(pixel.R(-vel/2, -vel/2, vel/2, vel/2))
	result.vel[1] = randomV(pixel.R(-vel/2, -vel/2, vel/2, vel/2))
	return result
}

func (m *mline) move(bounds *pixel.Rect) {
	for i := 0; i < 2; i++ {
		x, y := m.line[i].X, m.line[i].Y
		vx, vy := m.vel[i].X, m.vel[i].Y

		if x+vx < 0 || x+vx > wwidth {
			vx = -vx
		}
		if y+vy < 0 || y+vy > wheight {
			vy = -vy
		}

		m.line[i].X, m.line[i].Y = x+vx, y+vy
		m.vel[i].X, m.vel[i].Y = vx, vy
	}
}

func run() {
	bounds := pixel.R(0, 0, wwidth, wheight)
	var mlines [qlines]mline

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

	mlines[0] = NewMline(bounds)
	for !win.Closed() {
		win.SetClosed(win.JustPressed(pixelgl.KeyEscape))

		imd.Clear()
		copy(mlines[1:], mlines[0:qlines-2])
		mlines[0].move(&bounds)
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
