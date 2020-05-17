package main

import (
	"time"

	"github.com/SpacedMonkeyTCT/connect-four/gui"
	"github.com/faiface/pixel/pixelgl"
)

func main() {
	pixelgl.Run(connectFour)
}

func connectFour() {
	g := gui.New()

	for last := time.Now(); !g.Closed(); {
		_ = time.Since(last).Seconds()
		last = time.Now()

		_ = g.CheckForMove()
		g.Update()
	}
}
