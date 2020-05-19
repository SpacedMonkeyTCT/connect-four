package main

import (
	"time"

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
	g := gui.New(width, height)
	player := redPlayer

	for last := time.Now(); !g.Closed(); {
		_ = time.Since(last).Seconds()
		last = time.Now()

		if move := g.CheckForMove(); move > 0 {
			// check move is valid with game
			row := columnHeight(move)
			if player == redPlayer {
				g.AddRedChip(row, move)
			} else {
				g.AddBlueChip(row, move)
			}
			player = swapPlayer(player)
		}
		g.Update()
	}
}

func swapPlayer(player int) int {
	return player ^ (redPlayer | bluePlayer)
}

// for testing purposes
var columns [8]int

func columnHeight(c int) int {
	columns[c]++
	return columns[c]
}
