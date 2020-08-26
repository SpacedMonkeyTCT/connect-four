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
	holdY           float64
}

func New(width, height int) *GUI {
	cfg := pixelgl.WindowConfig{
		Title:  "Connect Four",
		Bounds: pixel.R(0, 0, float64(tileSize*(width+1)), float64(tileSize*(height+2))),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	winHeight := win.Bounds().H()
	holdY := winHeight - float64(3*tileSize/4)

	t := NewTileSet("tiles.png", tileSize)
	b := NewBoard(win, width, height, t.Get(1, 0))
	rcf := NewChipFactory(win, t.Get(0, 1))
	rc := rcf.New()
	bcf := NewChipFactory(win, t.Get(1, 1))

	g := GUI{
		win:             win,
		board:           b,
		redChipFactory:  rcf,
		blueChipFactory: bcf,
		currentChip:     rc,
		holdY:           holdY,
	}

	return &g
}

func (g GUI) Closed() bool {
	return g.win.Closed()
}

func (g GUI) CheckForMove() int {
	return g.board.CheckForMove()
}

func (g *GUI) Update() {
	x := g.win.MousePosition().X
	pos := pixel.V(x, g.holdY)
	g.currentChip.SetPos(pos)
}

func (g GUI) Draw() {
	g.win.Clear(colornames.Skyblue)
	g.board.Draw()
	g.currentChip.Draw()
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
