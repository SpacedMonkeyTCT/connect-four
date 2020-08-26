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
	win             *pixelgl.Window
	board           *board
	player          int
	currentChip     *chip
	redChipFactory  chipFactory
	blueChipFactory chipFactory
	holdY           float64
}

func New(game connectFour) *GUI {
	columns := game.Columns()
	rows := game.Rows()
	win := newWindow(columns, rows, tileSize)
	winHeight := win.Bounds().H()
	holdY := winHeight - float64(3*tileSize/4)

	t := NewTileSet("tiles.png", tileSize)
	b := NewBoard(win, columns, rows, t.GetBoardTile())
	rcf := NewChipFactory(win, t.GetRedChip())
	bcf := NewChipFactory(win, t.GetBlueChip())

	g := GUI{
		game:            game,
		win:             win,
		board:           b,
		redChipFactory:  rcf,
		blueChipFactory: bcf,
		currentChip:     rcf.New(),
		holdY:           holdY,
	}

	return &g
}

func (g *GUI) ProcessInput() {
	if column := g.board.CheckForMove(); column > 0 {

		if row := g.game.MakeMove(column); row > 0 {
			g.addCurrentChipToBoard(column, row)
			g.player = g.game.CurrentPlayer()

			if g.player == redPlayer {
				g.newRedChip()
			} else {
				g.newBlueChip()
			}
		}
	}
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

func (g GUI) Closed() bool {
	return g.win.Closed()
}

func newWindow(columns, rows, tileSize int) *pixelgl.Window {
	cfg := pixelgl.WindowConfig{
		Title:  "Connect Four",
		Bounds: pixel.R(0, 0, float64(tileSize*(columns+1)), float64(tileSize*(rows+2))),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	return win
}

func (g *GUI) newRedChip() {
	g.currentChip = g.redChipFactory.New()
}

func (g *GUI) newBlueChip() {
	g.currentChip = g.blueChipFactory.New()
}

func (g *GUI) addCurrentChipToBoard(column, row int) {
	g.board.AddChip(g.currentChip, column, row)
}
