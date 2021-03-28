package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	left  *ebiten.Image
	right *ebiten.Image
	up    *ebiten.Image
	down  *ebiten.Image
}

type Position struct {
	x int
	y int
}

type Player struct {
	id byte
	Sprite
	*Position
	face *ebiten.Image
}

func (p *Player) assign() []byte {
	return []byte{p.id, byte(p.x), byte(p.y), ASSIGN}
}

func (p *Player) sendMove(direction byte) []byte {
	switch direction {
	case LEFT:
		return []byte{p.id, byte(p.x - 1), byte(p.y), MOVE, LEFT}
	case RIGHT:
		return []byte{p.id, byte(p.x + 1), byte(p.y), MOVE, RIGHT}
	case UP:
		return []byte{p.id, byte(p.x), byte(p.y - 1), MOVE, UP}
	case DOWN:
		return []byte{p.id, byte(p.x), byte(p.y + 1), MOVE, DOWN}
	default:
		return []byte{p.id, byte(10), byte(10), MOVE, DOWN}
	}
}

func (p *Player) move(direction, x, y byte) {
	switch direction {
	case LEFT:
		fmt.Println("Player", p.id, "> Pos(", p.x, ",", p.y, ")")
		p.face = p.left
	case RIGHT:
		fmt.Println("Player", p.id, "> Pos(", p.x, ",", p.y, ")")
		p.face = p.right
	case UP:
		fmt.Println("Player", p.id, "> Pos(", p.x, ",", p.y, ")")
		p.face = p.up
	case DOWN:
		fmt.Println("Player", p.id, "> Pos(", p.x, ",", p.y, ")")
		p.face = p.down
	}
	p.x = int(x)
	p.y = int(y)
}

func (p *Player) setupPlayerSprite() {
	p.Sprite = Sprite{
		left:  getImg("images/p1l.png"),
		right: getImg("images/p1r.png"),
		up:    getImg("images/p1u.png"),
		down:  getImg("images/p1d.png"),
	}
	p.face = p.left
}

func createPlayer(id byte, pos *Position) *Player {
	p := &Player{
		id:       id,
		Position: pos,
	}
	p.setupPlayerSprite()
	return p
}
