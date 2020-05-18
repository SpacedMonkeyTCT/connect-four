package main

import (
	"fmt"
	"time"

	"github.com/SpacedMonkeyTCT/connect-four/gui"
	"github.com/faiface/pixel/pixelgl"
)

const (
	redPlayer  = 1
	bluePlayer = 2
)

func main() {
	pixelgl.Run(connectFour)
}

func connectFour() {
	g := gui.New()
	player := redPlayer

	for last := time.Now(); !g.Closed(); {
		_ = time.Since(last).Seconds()
		last = time.Now()

		if move := g.CheckForMove(); move > 0 {
			fmt.Println("Move:", move)
			// check move is valid with game
			row := height(move)
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

func height(c int) int {
	columns[c]++
	return columns[c]
}
