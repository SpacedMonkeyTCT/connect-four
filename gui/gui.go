package gui

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	tileSize = 32
)

type GUI struct {
	win             *pixelgl.Window
	board           *board
	redChipFactory  chipFactory
	blueChipFactory chipFactory
	currentChip     *chip
}

func New(width, height int) *GUI {
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
	rc := rcf.New()
	bcf := newChipFactory(win, t.get(1, 1))

	g := GUI{
		win:             win,
		board:           b,
		redChipFactory:  rcf,
		blueChipFactory: bcf,
		currentChip:     rc,
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
	g.board.Draw()
	g.win.Update()
}

func (g *GUI) NewRedChip() {
	g.currentChip = g.redChipFactory.New()
}

func (g *GUI) NewBlueChip() {
	g.currentChip = g.blueChipFactory.New()
}

func (g *GUI) AddCurrentChipToBoard(row, column int) {
	g.board.AddChip(g.currentChip, row, column)
}
