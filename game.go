package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
	"log"
)

const (
	usernamelbl = `username: `
	passwordlbl = `password: `
	loading     = `loading...`
)

var (
	mplusNormalFont font.Face
)

func init() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}
	const dpi = 72
	mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	player  *Player
	players map[byte]*Player
	state   int
}

func (g *Game) MoveActualPlayer() {
	if ebiten.IsKeyPressed(ebiten.KeyH) {
		g.player.move(LEFT)
		g.player.face = g.player.left
	} else if ebiten.IsKeyPressed(ebiten.KeyJ) {
		g.player.move(DOWN)
		g.player.face = g.player.down
	} else if ebiten.IsKeyPressed(ebiten.KeyK) {
		g.player.move(UP)
		g.player.face = g.player.up
	} else if ebiten.IsKeyPressed(ebiten.KeyL) {
		g.player.move(RIGHT)
		g.player.face = g.player.right
	}
}

const (
	USERNAMETYPING = iota
	PASSWORDTYPING
	WAITING
	GAMEPLAY
)

func (g *Game) Update() error {
	switch g.state {
	case GAMEPLAY:
		g.MoveActualPlayer()
	case USERNAMETYPING:
		g.player.Username += string(ebiten.InputChars())
		if repeatingKeyPressed(ebiten.KeyEnter) || repeatingKeyPressed(ebiten.KeyKPEnter) {
			g.state = PASSWORDTYPING
		}
		if repeatingKeyPressed(ebiten.KeyBackspace) {
			if len(g.player.Username) >= 1 {
				g.player.Username = g.player.Username[:len(g.player.Username)-1]
			}
		}
	case PASSWORDTYPING:
		g.player.Password += string(ebiten.InputChars())
		if repeatingKeyPressed(ebiten.KeyEnter) || repeatingKeyPressed(ebiten.KeyKPEnter) {
			go g.player.getToken()
			g.state = WAITING
		}
		if repeatingKeyPressed(ebiten.KeyBackspace) {
			if len(g.player.Password) >= 1 {
				g.player.Password = g.player.Password[:len(g.player.Password)-1]
			}
		}
	case WAITING:
		if g.player.state == LOGGEDIN {
			g.state = GAMEPLAY
		}
		return nil

	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Green)
	switch g.state {
	case GAMEPLAY:
		text.Draw(screen, g.player.token, mplusNormalFont, 160, 80, color.White)
		for _, player := range g.players {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(player.x, player.y)
			screen.DrawImage(player.face, op)

		}
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(g.player.x, g.player.y)
		screen.DrawImage(g.player.face, op)
	case USERNAMETYPING:
		text.Draw(screen, usernamelbl, mplusNormalFont, 20, 80, color.White)
		text.Draw(screen, g.player.Username, mplusNormalFont, 160, 80, color.White)
	case PASSWORDTYPING:
		text.Draw(screen, usernamelbl, mplusNormalFont, 20, 80, color.White)
		text.Draw(screen, g.player.Username, mplusNormalFont, 160, 80, color.White)
		text.Draw(screen, passwordlbl, mplusNormalFont, 20, 120, color.White)
		text.Draw(screen, g.player.Password, mplusNormalFont, 160, 120, color.White)
	case WAITING:
		text.Draw(screen, loading, mplusNormalFont, 20, 120, color.White)
	}
}
