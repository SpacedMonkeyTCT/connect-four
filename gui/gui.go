package gui

import (
	"fmt"

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
	CheckForWin() int
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
	ceiling         float64
	currentChip     *chip
	column          int
	row             int
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
		ceiling:         window.Bounds().H() - float64(3*tileSize/4),
	}
	g.currentChip = g.newChip()

	return g
}

func (g *GUI) ProcessInput() {
	if g.currentChip.Dropping() {
		return
	}

	if column := g.board.CheckForMove(); column > 0 {

		if row := g.game.MakeMove(column); row > 0 {
			// align chip to column
			x := g.board.Xpos(column)
			g.currentChip.SetXPos(x)

			// drop chip to floor
			y := g.board.Ypos(row)
			g.currentChip.Drop(y)

			// store for adding to the board after drop completes
			g.column = column
			g.row = row
		}
	}
}

func (g *GUI) Update() {
	if winner := g.game.CheckForWin(); winner > 0 {
		fmt.Println("Winner was player:", winner)
	}
	if dropped := g.currentChip.Update(); dropped {
		g.board.AddChip(g.currentChip, g.column, g.row)
		g.currentChip = g.newChip()
	}
}

func (g GUI) Draw() {
	g.window.Clear(colornames.Skyblue)
	g.currentChip.Draw()
	g.board.Draw()
	g.window.Update()
}

func (g GUI) Closed() bool {
	return g.window.Closed()
}

func (g GUI) newChip() *chip {
	pos := pixel.V(g.window.MousePosition().X, g.ceiling)
	if g.game.CurrentPlayer() == redPlayer {
		return g.redChipFactory.New(pos)
	}
	return g.blueChipFactory.New(pos)
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
