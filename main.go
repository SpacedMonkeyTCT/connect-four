package main

import (
	"time"

	"github.com/SpacedMonkeyTCT/connect-four/game"
	"github.com/SpacedMonkeyTCT/connect-four/gui"
	"github.com/faiface/pixel/pixelgl"
)

const (
	width  = 7
	height = 6
)

func main() {
	pixelgl.Run(connectFour)
}

func connectFour() {
	cf := game.NewConnectFour(width, height)
	g := gui.New(cf)

	for last := time.Now(); !g.Closed(); {
		_ = time.Since(last).Seconds()
		last = time.Now()

		g.ProcessInput()
		g.Update()
		g.Draw()
	}
}
