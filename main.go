package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	game := &Game{
		player: createPlayer(1),
	}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Backend Game")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
