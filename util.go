package main

import (
	"encoding/json"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func getImg(path string) *ebiten.Image {
	//file, err := ebitenutil.OpenFile(path)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//img, err := png.Decode(file)
	//if err != nil {
	//	log.Fatal(err)
	//}
	img, err := ebitenutil.NewImageFromURL(path)
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

func readMapHttp() (_map [][]int, err error) {
	r, err := http.Get("https://mama.sh/map.json")
	if err != nil {
		return
	}
	err = json.NewDecoder(r.Body).Decode(&_map)
	if err != nil {
		return
	}
	return
}

func getMap() (_map [][]int, err error) {
	return [][]int{
		{243, 243, 218, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 218, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 218, 243, 243, 243, 243, 218, 218, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 218, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 218, 243, 243, 243, 218, 243, 243, 243, 243, 243, 243, 243, 243, 243, 218, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 218, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 218, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 218, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 218, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 218, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 218, 218, 218, 243, 243, 218, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 218, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 218, 243, 243, 218, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 218, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 218, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 218, 243, 243, 243, 243, 243, 243, 243, 218, 243, 243, 243, 243, 243, 218, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 218, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 218, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 218, 218, 243, 243, 243, 243, 243, 243, 243, 243, 243, 218, 243, 243, 218, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 218, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 218, 243, 243, 218, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 218, 218, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 218, 243, 243, 243, 243, 218, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 218, 243, 243, 243, 243, 243, 243, 243, 218, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 220, 247, 242, 0, 0, 0, 0, 0, 0, 0, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 197, 0, 0, 0, 0, 0, 0, 0, 220, 247, 242, 0, 0, 0, 0, 0, 0, 0, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 217, 0, 0, 0, 0, 0, 0, 0, 220, 247, 242, 0, 0, 0, 0, 0, 0, 0, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 195, 247, 217, 0, 0, 0, 0, 0, 0, 0, 220, 247, 242, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 220, 247, 217, 0, 0, 0, 0, 0, 0, 0, 220, 247, 242, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 220, 247, 217, 0, 0, 0, 0, 0, 0, 0, 220, 247, 242, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 220, 247, 217, 0, 0, 0, 0, 0, 0, 0, 245, 247, 242, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 220, 247, 217, 0, 0, 0, 196, 269, 269, 269, 270, 247, 242, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 220, 217, 0, 0, 220, 247, 217, 0, 0, 0, 220, 247, 247, 247, 247, 247, 242, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 196, 270, 267, 269, 269, 270, 247, 217, 0, 0, 0, 220, 247, 192, 194, 194, 194, 222, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 220, 247, 247, 247, 247, 247, 247, 217, 0, 0, 0, 220, 247, 242, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 220, 247, 192, 194, 194, 194, 194, 222, 0, 0, 0, 220, 247, 242, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 220, 247, 242, 0, 0, 0, 0, 0, 0, 0, 0, 220, 247, 242, 0, 0, 0, 0, 0, 0, 196, 269, 269, 269, 269, 0, 0, 0, 0, 0, 220, 247, 242, 0, 0, 0, 0, 0, 0, 0, 0, 220, 247, 242, 0, 0, 0, 0, 0, 0, 220, 247, 247, 247, 247, 0, 0, 0, 0, 0, 220, 247, 242, 0, 0, 0, 0, 0, 0, 0, 0, 220, 247, 267, 269, 269, 269, 197, 0, 0, 220, 247, 192, 193, 193, 0, 0, 0, 0, 0, 220, 247, 242, 0, 0, 0, 0, 0, 0, 0, 0, 220, 247, 247, 247, 247, 247, 242, 0, 0, 220, 247, 242, 0, 0, 0, 0, 0, 0, 0, 220, 247, 242, 0, 0, 0, 0, 0, 0, 0, 0, 221, 194, 194, 194, 195, 247, 267, 269, 269, 270, 247, 242, 0, 0, 0, 0, 0, 0, 0, 220, 247, 242, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 220, 247, 247, 247, 247, 247, 247, 242, 0, 0, 0, 0, 0, 0, 0, 220, 247, 267, 269, 269, 269, 269, 269, 269, 197, 0, 0, 0, 0, 0, 220, 247, 192, 193, 193, 193, 193, 222, 0, 0, 0, 0, 0, 0, 0, 220, 247, 247, 247, 247, 247, 247, 247, 247, 242, 0, 0, 0, 0, 0, 220, 247, 242, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 221, 194, 194, 194, 194, 194, 194, 195, 247, 242, 0, 0, 245, 242, 0, 220, 247, 242, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 220, 247, 267, 269, 269, 270, 267, 269, 270, 247, 242, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 220, 247, 247, 247, 247, 247, 247, 247, 247, 247, 242, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 221, 194, 194, 194, 194, 194, 194, 194, 194, 194, 222, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 287, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 45, 46, 47, 48, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 70, 71, 72, 73, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 45, 46, 47, 48, 26, 27, 28, 29, 30, 31, 0, 0, 0, 0, 0, 95, 96, 97, 98, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 70, 71, 72, 73, 51, 52, 53, 54, 55, 56, 0, 0, 0, 0, 0, 120, 121, 122, 123, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 95, 96, 97, 98, 76, 77, 78, 79, 80, 81, 0, 0, 0, 0, 0, 145, 146, 147, 148, 0, 0, 0, 0, 0, 0, 45, 46, 47, 48, 0, 120, 121, 122, 123, 101, 102, 103, 104, 105, 106, 0, 0, 0, 0, 0, 0, 0, 0, 0, 287, 0, 0, 0, 0, 0, 70, 71, 72, 73, 0, 145, 146, 147, 148, 126, 127, 128, 129, 130, 131, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 95, 96, 97, 98, 0, 0, 0, 0, 0, 151, 303, 153, 154, 303, 156, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 120, 121, 122, 123, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 45, 46, 47, 48, 0, 0, 145, 146, 147, 148, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 70, 71, 72, 73, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 95, 96, 97, 98, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 188, 189, 189, 189, 189, 189, 189, 190, 0, 0, 0, 120, 121, 122, 123, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 213, 0, 0, 0, 0, 0, 301, 213, 0, 0, 0, 287, 146, 147, 148, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 213, 0, 301, 0, 0, 0, 0, 213, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 213, 0, 0, 0, 301, 0, 0, 213, 0, 0, 0, 0, 0, 0, 0, 287, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 213, 0, 0, 0, 0, 0, 0, 213, 58, 59, 60, 61, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 238, 239, 239, 239, 239, 239, 239, 240, 83, 84, 85, 86, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 108, 109, 110, 111, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 186, 133, 134, 135, 136, 0, 0, 0, 0, 0, 45, 46, 47, 48, 0, 0, 0, 0, 0, 0, 0, 0, 0, 45, 46, 47, 48, 0, 0, 0, 211, 304, 159, 160, 304, 0, 0, 0, 0, 0, 70, 71, 72, 73, 0, 0, 0, 0, 0, 0, 0, 0, 0, 70, 71, 72, 73, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 95, 96, 97, 98, 0, 0, 0, 0, 0, 0, 0, 0, 0, 95, 96, 97, 98, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 120, 121, 122, 123, 0, 0, 0, 0, 0, 0, 0, 0, 0, 120, 121, 122, 123, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 145, 146, 147, 148, 0, 0, 0, 0, 0, 0, 0, 0, 0, 145, 146, 147, 148, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}, nil
}
