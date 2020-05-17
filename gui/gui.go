package gui

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type GUI struct {
	win   *pixelgl.Window
	board board
}

const (
	width    = 7
	height   = 6
	tileSize = 32
)

func New() *GUI {
	cfg := pixelgl.WindowConfig{
		Title:  "Connect Four",
		Bounds: pixel.R(0, 0, float64(tileSize*(width+1)), float64(tileSize*(height+1))),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.Clear(colornames.Skyblue)
	t := newTiles("tiles.png", tileSize)
	b := newBoard(win, width, height, t.get(1, 0))
	g := GUI{
		win:   win,
		board: b,
	}

	return &g
}

func (g *GUI) Closed() bool {
	return g.win.Closed()
}

func (g *GUI) Update() {
	g.board.Update()
	g.win.Update()
}
