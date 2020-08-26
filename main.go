package main

import (
	"time"

	"github.com/SpacedMonkeyTCT/connect-four/game"
	"github.com/SpacedMonkeyTCT/connect-four/gui"
	"github.com/faiface/pixel/pixelgl"
)

const (
	columns = 7
	rows    = 6
)

func main() {
	pixelgl.Run(connectFour)
}

func connectFour() {
	cf := game.NewConnectFour(columns, rows)
	g := gui.New(cf)

	for last := time.Now(); !g.Closed(); {
		_ = time.Since(last).Seconds()
		last = time.Now()

		g.ProcessInput()
		g.Update()
		g.Draw()
	}
}
