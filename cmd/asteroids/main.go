package main

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"

	"github.com/jenska/learnpixel/pkg/asteroids"
)

const (
	initialAsteroids = 10 // initial number of asteroids
	explodeIntoParts = 3  // exploded asteroid splits into parts

	winWidth  = 1024
	winHeight = 768
)

func run() {
	bounds := pixel.R(0, 0, winWidth, winHeight)
	win, err := pixelgl.NewWindow(pixelgl.WindowConfig{
		Title:  "Asteroids",
		Bounds: bounds,
		VSync:  true,
	})
	if err != nil {
		panic(err)
	}

	aList := make([]*asteroids.Asteroid, initialAsteroids)
	for i := range aList {
		aList[i] = asteroids.NewAsteroid(rand.Intn(5)+5, &bounds)
	}

	ship := asteroids.NewSpaceship(&bounds)
	imd := imdraw.New(nil)
	imd.Precision = 7

	last := time.Now()
	for !win.Closed() {
		win.SetClosed(win.JustPressed(pixelgl.KeyEscape) || win.JustPressed(pixelgl.KeyQ))
		win.Clear(color.Black)

		d := time.Since(last).Seconds()
		last = time.Now()

		switch {
		case win.Pressed(pixelgl.KeyLeft):
			ship.RotateLeft()
		case win.Pressed(pixelgl.KeyRight):
			ship.RotateRight()
		case win.Pressed(pixelgl.KeyUp):
			ship.Forward()
		case win.Pressed(pixelgl.KeyDown):
			ship.Backward()
		case win.JustPressed(pixelgl.KeySpace):
			// shoot
		case win.JustPressed(pixelgl.KeyP):
			// pause
		}

		imd.Clear()

		for _, a := range aList {
			a.Update(d, &bounds)
			a.Draw(imd)
		}

		ship.Update(d, &bounds)
		ship.Draw(imd)

		imd.Draw(win)
		win.Update()
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	pixelgl.Run(run)
}
