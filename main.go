package main

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image"
	"log"
)

//Setting the font
func init() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}
	const dpi = 72
	mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    14,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}

//Setting the tile image
func init() {
	img, _, err := image.Decode(bytes.NewReader(images.Tiles_png))
	if err != nil {
		log.Fatal(err)
	}
	tilesImage = ebiten.NewImageFromImage(img)
}

//Starting the game
func main() {
	layers, err := getMap()
	if err != nil {
		log.Fatal(err)
	}
	game := &Game{
		layers:  layers,
		canMove: true,
		user: &User{
			Username: "martin",
			Password: "T3stpass!",
		},
		players: make(map[byte]*Player),
	}
	go game.movementTimer()
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Backend Game")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
