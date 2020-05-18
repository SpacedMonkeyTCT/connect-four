package gui

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type chipFactory struct {
	win    *pixelgl.Window
	sprite *pixel.Sprite
}

func newChipFactory(win *pixelgl.Window, sprite *pixel.Sprite) chipFactory {
	return chipFactory{
		win:    win,
		sprite: sprite,
	}
}

func (cf chipFactory) New(pos pixel.Vec) chip {
	return chip{
		win:    cf.win,
		sprite: cf.sprite,
		pos:    pos,
	}
}

type chip struct {
	win    *pixelgl.Window
	sprite *pixel.Sprite
	pos    pixel.Vec
}

func (c chip) Update() {
	c.sprite.Draw(c.win, pixel.IM.Moved(c.pos))
}
