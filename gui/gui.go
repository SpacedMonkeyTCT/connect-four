package gui

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	redPlayer = 1
	// bluePlayer = 2
	tileSize = 32
)

type connectFour interface {
	MakeMove(int) int
	Columns() int
	Rows() int
	CurrentPlayer() int
}

type GUI struct {
	game            connectFour
	window          *pixelgl.Window
	board           *board
	redChipFactory  chipFactory
	blueChipFactory chipFactory
	holdY           float64
	currentChip     *chip
}

func New(game connectFour) *GUI {
	columns := game.Columns()
	rows := game.Rows()
	window := newWindow(columns, rows, tileSize)
	t := NewTileSet("tiles.png", tileSize)

	g := &GUI{
		game:            game,
		window:          window,
		board:           NewBoard(window, columns, rows, t.GetBoardTile()),
		redChipFactory:  NewChipFactory(window, t.GetRedChip()),
		blueChipFactory: NewChipFactory(window, t.GetBlueChip()),
		holdY:           window.Bounds().H() - float64(3*tileSize/4),
	}
	g.currentChip = g.newChip(game.CurrentPlayer())

	return g
}

func (g *GUI) ProcessInput() {
	if column := g.board.CheckForMove(); column > 0 {

		if row := g.game.MakeMove(column); row > 0 {
			g.board.AddChip(g.currentChip, column, row)
			g.currentChip = g.newChip(g.game.CurrentPlayer())
		}
	}
}

func (g *GUI) Update() {
	x := g.window.MousePosition().X
	pos := pixel.V(x, g.holdY)
	g.currentChip.SetPos(pos)
}

func (g GUI) Draw() {
	g.window.Clear(colornames.Skyblue)
	g.board.Draw()
	g.currentChip.Draw()
	g.window.Update()
}

func (g GUI) Closed() bool {
	return g.window.Closed()
}

func (g GUI) newChip(player int) *chip {
	if player == redPlayer {
		return g.redChipFactory.New()
	}
	return g.blueChipFactory.New()
}

func newWindow(columns, rows, tileSize int) *pixelgl.Window {
	cfg := pixelgl.WindowConfig{
		Title:  "Connect Four",
		Bounds: pixel.R(0, 0, float64(tileSize*(columns+1)), float64(tileSize*(rows+2))),
		VSync:  true,
	}
	window, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	return window
}
