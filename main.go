package main

import (
	_ "image/png"

	"github.com/SpacedMonkeyTCT/connect-four/board"
	"github.com/SpacedMonkeyTCT/connect-four/gui"
	"github.com/SpacedMonkeyTCT/connect-four/sprites"
	"github.com/faiface/pixel/pixelgl"
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
	tiles := sprites.New("tiles.png", tileSize)
	b := board.New(width, height)
	g := gui.New(tiles, b)

	for !g.Closed() {
		g.DrawBoard()
		g.Update()
	}
}
