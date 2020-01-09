package main

import (
	"fmt"
	"time"

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

	player := board.Red
	move := 0

	for last := time.Now(); !g.Closed(); {
		_ = time.Since(last).Seconds()
		last = time.Now()

		if g.TokenDropping() {
			g.MoveToken()
		} else {
			move = g.CheckForMove()
			if move > 0 {
				fmt.Println("swapped")
				player = player.Swap()
			}
		}
		g.DrawBoard()
		g.Update()
	}
}
