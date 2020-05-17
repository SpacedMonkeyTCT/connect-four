package gui

import (
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type board struct {
	win      *pixelgl.Window
	width    int
	height   int
	tile     *pixel.Sprite
	tileSize float64
	rect     pixel.Rect
}

func newBoard(win *pixelgl.Window, width, height int, tile *pixel.Sprite) board {

	wb := win.Bounds()
	windowWidth := wb.W()
	windowHeight := wb.H()

	size := tile.Frame().W()
	boardWidth := float64(width) * size
	boardHeight := float64(height) * size

	xoff := math.Max(0, (windowWidth-boardWidth)/2.0)
	yoff := math.Max(0, (windowHeight-boardHeight)/2.0)
	rect := pixel.R(xoff, yoff, xoff+boardWidth, yoff+boardHeight)

	return board{
		win:      win,
		width:    width,
		height:   height,
		tile:     tile,
		tileSize: size,
		rect:     rect,
	}
}

func (b board) CheckForMove() int {
	if b.win.JustPressed(pixelgl.MouseButtonLeft) {

		if b.rect.Contains(b.win.MousePosition()) {
			f := math.Floor((b.win.MousePosition().X - b.rect.Min.X) / b.tileSize)
			i := int(f) + 1
			if i > b.width {
				i = b.width
			}
			return i
		}
	}
	return 0
}

func (b board) Update() {
	for x := b.rect.Min.X; x < b.rect.Max.X; x += b.tileSize {
		for y := b.rect.Min.Y; y < b.rect.Max.Y; y += b.tileSize {
			pos := pixel.V(x+tileSize/2, y+tileSize/2)
			b.tile.Draw(b.win, pixel.IM.Moved(pos))
		}
	}
}
