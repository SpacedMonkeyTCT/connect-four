package board

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Board struct {
	width     int
	height    int
	tileSize  int
	boardTile *pixel.Sprite
}

func New(width, height, tileSize int, boardTile *pixel.Sprite) Board {
	return Board{
		width:     width,
		height:    height,
		tileSize:  tileSize,
		boardTile: boardTile,
	}
}

func (b Board) Draw(win *pixelgl.Window) {
	boardPos := pixel.V(float64(b.tileSize*2), win.Bounds().H()-float64(b.tileSize)*2.25)

	for v := 0; v < b.height; v++ {
		for u := 0; u < b.width; u++ {
			xOff := b.tileSize * u
			yOff := b.tileSize * -v
			pos := pixel.V(float64(xOff), float64(yOff))
			b.boardTile.Draw(win, pixel.IM.Moved(boardPos).Moved(pos))
		}
	}
}
