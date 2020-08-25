package main

import (
	"time"

	"github.com/SpacedMonkeyTCT/connect-four/game"
	"github.com/SpacedMonkeyTCT/connect-four/gui"
	"github.com/faiface/pixel/pixelgl"
)

const (
	redPlayer  = 1
	bluePlayer = 2
	width      = 7
	height     = 6
)

func main() {
	pixelgl.Run(connectFour)
}

func connectFour() {
	cf := game.NewConnectFour(width, height, redPlayer, bluePlayer)
	g := gui.New(width, height)
	player := redPlayer

	for last := time.Now(); !g.Closed(); {
		_ = time.Since(last).Seconds()
		last = time.Now()

		if column := g.CheckForMove(); column > 0 {

			if row := cf.MakeMove(column); row > 0 {
				g.AddCurrentChipToBoard(row, column)

				if player == redPlayer {
					g.NewBlueChip()
				} else {
					g.NewRedChip()
				}
				player = cf.CurrentPlayer()
			}
		}
		g.Update()
	}
}
