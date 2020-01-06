package main

import (
	_ "image/png"

	"github.com/SpacedMonkeyTCT/connect-four/sprites"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	tileSize = 32
)

func main() {
	pixelgl.Run(connectFour)
}

func connectFour() {
	cfg := pixelgl.WindowConfig{
		Title:  "Connect Four",
		Bounds: pixel.R(0, 0, 320, 288),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	tiles := sprites.New("tiles.png", tileSize)

	redToken := tiles.Get(0, 1)
	blueToken := tiles.Get(1, 1)
	boardTile := tiles.Get(1, 0)

	win.Clear(colornames.Skyblue)
	redToken.Draw(win, pixel.IM.Moved(pixel.V(tileSize, win.Bounds().H()-tileSize)))
	blueToken.Draw(win, pixel.IM.Moved(pixel.V(win.Bounds().W()-tileSize, win.Bounds().H()-tileSize)))

	boardPos := pixel.V(64, win.Bounds().H()-tileSize*2.25)

	for v := 0; v < 6; v++ {
		for u := 0; u < 7; u++ {
			xOff := tileSize * u
			yOff := tileSize * -v
			pos := pixel.V(float64(xOff), float64(yOff))
			boardTile.Draw(win, pixel.IM.Moved(boardPos).Moved(pos))
		}
	}

	for !win.Closed() {
		win.Update()
	}
}
