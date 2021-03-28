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
	x float64
	y float64
}

type Player struct {
	id int
	Sprite
	*Position
	face *ebiten.Image
}

func (p *Player) sendMove(direction byte) []byte {
	switch direction {
	case LEFT:
		return []byte{byte(p.id), byte(p.x - 1), byte(p.y), MOVE, LEFT}
	case RIGHT:
		return []byte{byte(p.id), byte(p.x + 1), byte(p.y), MOVE, RIGHT}
	case UP:
		return []byte{byte(p.id), byte(p.x), byte(p.y - 1), MOVE, UP}
	case DOWN:
		return []byte{byte(p.id), byte(p.x), byte(p.y + 1), MOVE, DOWN}
	default:
		return []byte{byte(p.id), byte(p.x), byte(p.y), MOVE, DOWN}
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
	p.x = float64(x)
	p.y = float64(y)
}

func (p *Player) setupPlayerSprite(playerId int) {
	fmt.Println(playerId, "Find color")
	p.Sprite = Sprite{
		left:  getImg("images/p1l.png"),
		right: getImg("images/p1r.png"),
		up:    getImg("images/p1u.png"),
		down:  getImg("images/p1d.png"),
	}
	p.face = p.left
}

func createPlayer(id int, pos *Position) *Player {
	p := &Player{
		id:       id,
		Position: pos,
	}
	p.setupPlayerSprite(id)
	return p
}
