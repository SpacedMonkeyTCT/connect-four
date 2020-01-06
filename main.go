package main

import (
	_ "image/png"

	"github.com/SpacedMonkeyTCT/connect-four/board"
	"github.com/SpacedMonkeyTCT/connect-four/sprites"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	width    = 7
	height   = 6
	tileSize = 32
)

func main() {
	pixelgl.Run(connectFour)
}

func connectFour() {
	cfg := pixelgl.WindowConfig{
		Title:  "Connect Four",
		Bounds: pixel.R(0, 0, tileSize*(width+3), tileSize*(height+3)),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.Clear(colornames.Skyblue)

	tiles := sprites.New("tiles.png", tileSize)

	redToken := tiles.Get(0, 1)
	redToken.Draw(win, pixel.IM.Moved(pixel.V(tileSize, win.Bounds().H()-tileSize)))

	blueToken := tiles.Get(1, 1)
	blueToken.Draw(win, pixel.IM.Moved(pixel.V(win.Bounds().W()-tileSize, win.Bounds().H()-tileSize)))

	boardTile := tiles.Get(1, 0)
	b := board.New(width, height, tileSize, boardTile)
	b.Draw(win)

	for !win.Closed() {
		win.Update()
	}
}
