package gui

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type chipFactory struct {
	win    *pixelgl.Window
	sprite *pixel.Sprite
}

func NewChipFactory(win *pixelgl.Window, sprite *pixel.Sprite) chipFactory {
	return chipFactory{
		win:    win,
		sprite: sprite,
	}
}

func (cf chipFactory) New() *chip {
	return &chip{
		win:    cf.win,
		sprite: cf.sprite,
		pos:    pixel.V(0, 0),
		vel:    pixel.V(0, 0),
		acc:    pixel.V(0, 0),
	}
}

type chip struct {
	win    *pixelgl.Window
	sprite *pixel.Sprite
	pos    pixel.Vec
	vel    pixel.Vec
	acc    pixel.Vec
}

func (c chip) Draw() {
	c.sprite.Draw(c.win, pixel.IM.Moved(c.pos))
}

func (c *chip) SetPos(pos pixel.Vec) {
	c.pos = pos
}
