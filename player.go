package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"strconv"
)

type Sprite struct {
	left  *ebiten.Image
	right *ebiten.Image
	up    *ebiten.Image
	down  *ebiten.Image
}

type Position struct {
	x float64
	y float64
}

type Player struct {
	*User
	id int
	Sprite
	*Position
	face    *ebiten.Image
}

func (p *Player) move(direction int) {
	switch direction {
	case LEFT:
		fmt.Println("Player",p.id,"> Pos(", p.x,",",p.y, ")")
		p.x--
	case RIGHT:
		fmt.Println("Player",p.id,"> Pos(", p.x,",",p.y, ")")
		p.x++
	case UP:
		fmt.Println("Player",p.id,"> Pos(", p.x,",",p.y, ")")
		p.y--
	case DOWN:
		fmt.Println("Player",p.id,"> Pos(", p.x,",",p.y, ")")
		p.y++
	}
}

func (p *Player) setupPlayerSprite(playerId int) {
	id := strconv.Itoa(playerId)
	p.Sprite = Sprite{
		left:  getImg("images/player" + id + "_l.png"),
		right: getImg("images/player" + id + "_r.png"),
		up:    getImg("images/player" + id + "_u.png"),
		down:  getImg("images/player" + id + "_d.png"),
	}
	p.face = p.left
}

func createPlayer(id int) *Player {
	p := &Player{
		id:     id,
		Position: &Position{
			x: 10,
			y: 10,
		},
		User: &User{},
	}
	p.setupPlayerSprite(id)
	return p
}

