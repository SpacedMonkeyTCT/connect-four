package main

import (
	"image"
	"os"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func main() {
	pixelgl.Run(connectFour)
}

func connectFour() {
	cfg := pixelgl.WindowConfig{
		Title:  "Connect Four",
		Bounds: pixel.R(0, 0, 320, 288),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	pic, err := loadPicture("tiles.png")
	if err != nil {
		panic(err)
	}

	redTokenRect := pixel.R(0, 32, 32, 64)
	redTokenSprite := pixel.NewSprite(pic, redTokenRect)

	blueTokenRect := pixel.R(32, 32, 64, 64)
	blueTokenSprite := pixel.NewSprite(pic, blueTokenRect)

	boardTileRect := pixel.R(32, 0, 64, 32)
	boardTileSprite := pixel.NewSprite(pic, boardTileRect)

	win.Clear(colornames.Skyblue)
	redTokenSprite.Draw(win, pixel.IM.Moved(pixel.V(32, win.Bounds().H()-24)))
	blueTokenSprite.Draw(win, pixel.IM.Moved(pixel.V(win.Bounds().W()-32, win.Bounds().H()-24)))

	boardPos := pixel.V(64, win.Bounds().H()-72)

	for v := 0; v < 6; v++ {
		for u := 0; u < 7; u++ {
			xOff := 32 * u
			yOff := 32 * -v
			pos := pixel.V(float64(xOff), float64(yOff))
			boardTileSprite.Draw(win, pixel.IM.Moved(boardPos).Moved(pos))
		}
	}

	for !win.Closed() {
		win.Update()
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
