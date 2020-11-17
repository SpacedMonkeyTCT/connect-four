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

func (cf chipFactory) New(pos pixel.Vec) *chip {
	return &chip{
		win:    cf.win,
		sprite: cf.sprite,
		pos:    pos,
		vel:    pixel.V(0, 0),
		acc:    pixel.V(0, 0),
		floor:  0.0,
	}
}

type chip struct {
	win    *pixelgl.Window
	sprite *pixel.Sprite
	pos    pixel.Vec
	vel    pixel.Vec
	acc    pixel.Vec
	floor  float64
}

func (c *chip) Update() bool {
	if c.floor > 0.0 {
		return c.fall()
	}
	c.slide()
	return false
}

func (c chip) Draw() {
	c.sprite.Draw(c.win, pixel.IM.Moved(c.pos))
}

func (c *chip) SetPos(pos pixel.Vec) {
	c.pos = pos
}

func (c *chip) SetXPos(x float64) {
	c.pos.X = x
}

func (c *chip) slide() {
	c.pos.X = c.win.MousePosition().X
}

func (c *chip) Drop(floor float64) {
	c.floor = floor
	c.acc = pixel.V(0, -1)
}

func (c *chip) fall() bool {
	c.vel = c.vel.Add(c.acc)
	c.pos = c.pos.Add(c.vel)

	if c.pos.Y < c.floor {
		c.acc = pixel.V(0, 0)
		c.vel = pixel.V(0, 0)
		c.pos.Y = c.floor
		c.floor = 0.0
		return true
	}
	return false
}
