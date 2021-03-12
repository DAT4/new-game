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

func (p *Player) move(direction int, pos float64) {
	switch direction {
	case LEFT:
		fmt.Println("Player",p.id,"> Pos(", p.x,",",p.y, ")")
		p.face = p.left
		p.x = pos
	case RIGHT:
		fmt.Println("Player",p.id,"> Pos(", p.x,",",p.y, ")")
		p.face = p.right
		p.x = pos
	case UP:
		fmt.Println("Player",p.id,"> Pos(", p.x,",",p.y, ")")
		p.face = p.up
		p.y = pos
	case DOWN:
		fmt.Println("Player",p.id,"> Pos(", p.x,",",p.y, ")")
		p.face = p.down
		p.y = pos
	}
}

func (p *Player) setupPlayerSprite(playerId int) {
	id := strconv.Itoa(playerId)
	p.Sprite = Sprite{
		left:  getImg("images/p" + id + "l.png"),
		right: getImg("images/p" + id + "r.png"),
		up:    getImg("images/p" + id + "u.png"),
		down:  getImg("images/p" + id + "d.png"),
	}
	p.face = p.left
}

func createPlayer(id int) *Player {
	p := &Player{
		id:     id,
		Position: &Position{
			x: 8,
			y: 8,
		},
		User: &User{
			Username: "martin",
			Password: "T3stpass!",
			State:    0,
			Token:    "",
		},
	}
	p.setupPlayerSprite(id)
	return p
}

