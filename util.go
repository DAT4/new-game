package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/png"
	"log"
)

func getImg(path string) *ebiten.Image {
	file, err := ebitenutil.OpenFile(path)
	if err != nil {
		log.Fatal(err)
	}
	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	out := ebiten.NewImageFromImage(img)
	return out
}
