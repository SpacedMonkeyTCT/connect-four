package gui

import (
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type board struct {
	win      *pixelgl.Window
	columns  int
	rows     int
	tile     *pixel.Sprite
	tileSize float64
	rect     pixel.Rect
	chips    []*chip
}

func NewBoard(win *pixelgl.Window, columns, rows int, tile *pixel.Sprite) *board {
	wb := win.Bounds()
	windowWidth := wb.W()
	windowHeight := wb.H()

	size := tile.Frame().W()
	boardWidth := float64(columns) * size
	boardHeight := float64(rows) * size

	xoff := math.Max(0, (windowWidth-boardWidth)/2.0)
	yoff := math.Max(0, (windowHeight-boardHeight)/2.0) - size/2
	rect := pixel.R(xoff, yoff, xoff+boardWidth, yoff+boardHeight)

	return &board{
		win:      win,
		columns:  columns,
		rows:     rows,
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
			if i > b.columns {
				i = b.columns
			}
			return i
		}
	}
	return 0
}

func (b board) Draw() {
	for _, c := range b.chips {
		c.Draw()
	}

	for x := b.rect.Min.X; x < b.rect.Max.X; x += b.tileSize {
		for y := b.rect.Min.Y; y < b.rect.Max.Y; y += b.tileSize {
			pos := pixel.V(x+b.tileSize/2, y+b.tileSize/2)
			b.tile.Draw(b.win, pixel.IM.Moved(pos))
		}
	}
}

func (b *board) AddChip(c *chip, column, row int) {
	x := b.Xpos(column)
	y := b.Ypos(row)
	pos := pixel.V(x, y)
	c.SetPos(pos)
	b.chips = append(b.chips, c)
}

func (b board) Xpos(column int) float64 {
	return b.rect.Min.X + b.tileSize*float64(column) - b.tileSize/2
}

func (b board) Ypos(row int) float64 {
	return b.rect.Min.Y + b.tileSize*float64(row) - b.tileSize/2
}
