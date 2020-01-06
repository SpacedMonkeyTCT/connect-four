package sprites

import (
	"image"
	"os"

	"github.com/faiface/pixel"
)

type Sprites struct {
	tileSize int
	pic      pixel.Picture
}

func New(filePath string, tileSize int) Sprites {
	pic, err := loadPicture(filePath)
	if err != nil {
		panic(err)
	}

	return Sprites{
		tileSize: tileSize,
		pic:      pic,
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

func (s Sprites) Get(x, y int) *pixel.Sprite {
	u := float64(x * s.tileSize)
	v := float64(y * s.tileSize)
	rect := pixel.R(u, v, u+float64(s.tileSize), v+float64(s.tileSize))
	return pixel.NewSprite(s.pic, rect)
}
