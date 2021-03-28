package main

import (
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
		p.face = p.left
	case RIGHT:
		p.face = p.right
	case UP:
		p.face = p.up
	case DOWN:
		p.face = p.down
	}
	p.x = int(x)
	p.y = int(y)
}

func (p *Player) setupPlayerSprite() {
	p.Sprite = Sprite{
		left:  getImg("https://mama.sh/p1l.png"),
		right: getImg("https://mama.sh/p1r.png"),
		up:    getImg("https://mama.sh/p1u.png"),
		down:  getImg("https://mama.sh/p1d.png"),
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
