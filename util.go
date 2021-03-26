package main

import (
	"encoding/json"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/png"
	"io/ioutil"
	"log"
	"os"
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

func readMap() (_map [][]int, err error) {
	file, err := os.Open("map.json")
	if err != nil {
		return
	}
	defer file.Close()
	layers, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}
	err = json.Unmarshal(layers, &_map)
	return
}
