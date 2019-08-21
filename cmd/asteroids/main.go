package main

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/jenska/learnpixel/pkg/asteroids"
	"github.com/jenska/learnpixel/pkg/core"
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
		Title:     "Asteroids",
		Bounds:    bounds,
		Resizable: true,
		VSync:     true,
	})
	if err != nil {
		panic(err)
	}
	objects := core.NewSet()

	for i := 0; i < initialAsteroids; i++ {
		objects.Put(asteroids.NewAsteroid(&bounds))
	}

	score := asteroids.NewScore(&bounds)
	objects.Put(&score)

	ship := asteroids.NewSpaceship(&bounds,
		func() bool { return win.Pressed(pixelgl.KeyLeft) || win.Pressed(pixelgl.KeyA) },
		func() bool { return win.Pressed(pixelgl.KeyRight) || win.Pressed(pixelgl.KeyD) },
		func() bool { return win.Pressed(pixelgl.KeyUp) || win.Pressed(pixelgl.KeyW) },
		func() bool { return win.Pressed(pixelgl.KeyDown) || win.Pressed(pixelgl.KeyS) },
		func() bool { return win.JustPressed(pixelgl.KeySpace) },
	)
	objects.Put(&ship)

	imd := imdraw.New(nil)
	imd.Precision = 7
	paused := false
	last := time.Now()
	for !win.Closed() {
		win.SetClosed(win.JustPressed(pixelgl.KeyEscape) || win.JustPressed(pixelgl.KeyQ))
		if win.JustPressed(pixelgl.KeyP) {
			paused = !paused
		}
		win.Clear(color.Black)
		if !paused {
			d := time.Since(last).Seconds()
			last = time.Now()
			imd.Clear()

			objects.Do(
				func(object interface{}) {
					if element, ok := object.(asteroids.Drawable); ok {
						element.Update(d)
						element.Draw(imd)
					}
				})

			imd.Draw(win)
		}
		win.Update()
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	pixelgl.Run(run)
}
