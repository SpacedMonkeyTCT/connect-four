package gui

import (
	"github.com/SpacedMonkeyTCT/connect-four/board"
	"github.com/SpacedMonkeyTCT/connect-four/sprites"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type GUI struct {
	win       *pixelgl.Window
	tileSize  float64
	redToken  *pixel.Sprite
	blueToken *pixel.Sprite
	boardTile *pixel.Sprite
	board     board.Board
}

func New(s sprites.Sprites, b board.Board) GUI {
	g := GUI{
		tileSize:  float64(s.TileSize()),
		redToken:  s.Get(0, 1),
		blueToken: s.Get(1, 1),
		boardTile: s.Get(1, 0),
		board:     b,
	}

	cfg := pixelgl.WindowConfig{
		Title:  "Connect Four",
		Bounds: pixel.R(0, 0, g.tileSize*float64(b.Width()+3), g.tileSize*float64(b.Height()+3)),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.Clear(colornames.Skyblue)
	g.win = win

	return g
}

func (g GUI) DrawBoard() {
	g.redToken.Draw(g.win, pixel.IM.Moved(pixel.V(g.tileSize, g.win.Bounds().H()-g.tileSize)))
	g.blueToken.Draw(g.win, pixel.IM.Moved(pixel.V(g.win.Bounds().W()-g.tileSize, g.win.Bounds().H()-g.tileSize)))
	boardPos := pixel.V(float64(g.tileSize*2), g.win.Bounds().H()-float64(g.tileSize)*2.25)

	for v := 0; v < g.board.Height(); v++ {
		for u := 0; u < g.board.Width(); u++ {
			xOff := g.tileSize * float64(u)
			yOff := g.tileSize * float64(-v)
			pos := pixel.V(xOff, yOff)
			g.boardTile.Draw(g.win, pixel.IM.Moved(boardPos).Moved(pos))
		}
	}
}

func (g GUI) Closed() bool {
	return g.win.Closed()
}

func (g GUI) Update() {
	g.win.Update()
}
