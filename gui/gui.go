package gui

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type GUI struct {
	win      *pixelgl.Window
	board    board
	redChip  chipFactory
	blueChip chipFactory
	chips    []chip
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
	t := newTileSet("tiles.png", tileSize)
	b := newBoard(win, width, height, t.get(1, 0))
	rcf := newChipFactory(win, t.get(0, 1))
	bcf := newChipFactory(win, t.get(1, 1))
	g := GUI{
		win:      win,
		board:    b,
		redChip:  rcf,
		blueChip: bcf,
	}

	return &g
}

func (g *GUI) Closed() bool {
	return g.win.Closed()
}

func (g *GUI) CheckForMove() int {
	return g.board.CheckForMove()
}

func (g *GUI) Update() {
	for _, c := range g.chips {
		c.Update()
	}
	g.board.Update()
	g.win.Update()
}

func (g *GUI) AddRedChip(row, column int) {
	g.addChip(g.redChip, row, column)
}

func (g *GUI) AddBlueChip(row, column int) {
	g.addChip(g.blueChip, row, column)
}

func (g *GUI) addChip(cf chipFactory, row, column int) {
	pos := g.board.Pos(row, column)
	c := cf.New(pos)
	g.chips = append(g.chips, c)
}
