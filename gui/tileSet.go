package gui

import (
	"image"
	_ "image/png"
	"os"

	"github.com/faiface/pixel"
)

type tileSet struct {
	size float64
	pic  pixel.Picture
}

func newTileSet(filePath string, size int) tileSet {
	pic, err := loadPicture(filePath)
	if err != nil {
		panic(err)
	}

	return tileSet{
		size: float64(size),
		pic:  pic,
	}
}

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func (t tileSet) get(x, y int) *pixel.Sprite {
	u := float64(x) * t.size
	v := float64(y) * t.size
	rect := pixel.R(u, v, u+t.size, v+t.size)
	return pixel.NewSprite(t.pic, rect)
}

func (t tileSet) tileSize() int {
	return int(t.size)
}
