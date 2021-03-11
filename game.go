package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/colornames"
	"image"
	"image/color"
	"net/http"
	"sync"
)

type Game struct {
	sync.Mutex
	layers  [][]int
	player  *Player
	players map[byte]*Player
	states  states
}

func (g *Game) moveActualPlayer() {
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

func (g *Game) updateLoginState() {
	switch g.states.loginState {
	case PASSWORDTYPING:
		g.player.Password += string(ebiten.InputChars())
		if repeatingKeyPressed(ebiten.KeyEnter) || repeatingKeyPressed(ebiten.KeyKPEnter) {
			go g.getToken()
			g.states.loginState = WAITING
		}
		if repeatingKeyPressed(ebiten.KeyBackspace) {
			if len(g.player.Password) >= 1 {
				g.player.Password = g.player.Password[:len(g.player.Password)-1]
			}
		}
	case USERNAMETYPING:
		g.player.Username += string(ebiten.InputChars())
		if repeatingKeyPressed(ebiten.KeyEnter) || repeatingKeyPressed(ebiten.KeyKPEnter) {
			g.states.loginState = PASSWORDTYPING
		}
		if repeatingKeyPressed(ebiten.KeyBackspace) {
			if len(g.player.Username) >= 1 {
				g.player.Username = g.player.Username[:len(g.player.Username)-1]
			}
		}
	case WAITING:
		fmt.Println("What to do while waiting?")
	}
}

func (g *Game) drawGamePlay(screen *ebiten.Image) {
	const xNum = screenWidth / tileSize
	for _, l := range g.layers {
		for i, t := range l {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64((i%xNum)*tileSize), float64((i/xNum)*tileSize))

			sx := (t % tileXNum) * tileSize
			sy := (t / tileXNum) * tileSize
			screen.DrawImage(tilesImage.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image), op)
		}
	}
	text.Draw(screen, g.player.Token, mplusNormalFont, 160, 80, color.White)
	for _, player := range g.players {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(player.x, player.y)
		screen.DrawImage(player.face, op)

	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.player.x, g.player.y)
	screen.DrawImage(g.player.face, op)
}

func (g *Game) drawLoginScreen(screen *ebiten.Image) {
	switch g.states.loginState {
	case USERNAMETYPING:
		text.Draw(screen, usernamelbl, mplusNormalFont, 20, 80, color.White)
		text.Draw(screen, g.player.Username, mplusNormalFont, 100, 80, color.White)
	case PASSWORDTYPING:
		text.Draw(screen, usernamelbl, mplusNormalFont, 20, 80, color.White)
		text.Draw(screen, g.player.Username, mplusNormalFont, 100, 80, color.White)
		text.Draw(screen, passwordlbl, mplusNormalFont, 20, 120, color.White)
		text.Draw(screen, g.player.Password, mplusNormalFont, 100, 120, color.White)
	case WAITING:
		text.Draw(screen, loading, mplusNormalFont, 20, 120, color.White)
	}
}

func (g *Game) Update() error {
	switch g.states.globalState {
	case GAMEPLAY:
		g.moveActualPlayer()
	case LOGIN:
		g.updateLoginState()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Green)
	switch g.states.globalState {
	case GAMEPLAY:
		g.drawGamePlay(screen)
	case LOGIN:
		g.drawLoginScreen(screen)
	}
}

func (g *Game) getToken() {
	link := "http://localhost:8056/login"
	//link := "https://tmp.mama.sh/api/login"
	jsonStr, err := json.Marshal(g.player)
	if err != nil {
		fmt.Println(err)
		return
	}
	req, err := http.NewRequest("POST", link, bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	var tkn jwt
	err = json.NewDecoder(resp.Body).Decode(&tkn)
	if err != nil {
		fmt.Println(err)
		return
	}

	g.Lock()
	g.player.Token = tkn.Token
	g.states.globalState = GAMEPLAY
	g.Unlock()
}
