package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/colornames"
	"image"
	"image/color"
	"log"
	"net/http"
	"sync"
)

type Game struct {
	sync.Mutex
	user    *User
	conn    *websocket.Conn
	you     byte
	layers  [][]int
	players map[byte]*Player
	states  states
}

func (g *Game) moveActualPlayer() {
	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeyH):
		g.conn.WriteMessage(websocket.BinaryMessage, g.players[g.you].sendMove(LEFT))
	case inpututil.IsKeyJustPressed(ebiten.KeyJ):
		g.conn.WriteMessage(websocket.BinaryMessage, g.players[g.you].sendMove(LEFT))
	case inpututil.IsKeyJustPressed(ebiten.KeyK):
		g.conn.WriteMessage(websocket.BinaryMessage, g.players[g.you].sendMove(LEFT))
	case inpututil.IsKeyJustPressed(ebiten.KeyL):
		g.conn.WriteMessage(websocket.BinaryMessage, g.players[g.you].sendMove(LEFT))
	}
}

func (g *Game) updateLoginState() {
	switch g.states.loginState {
	case PASSWORDTYPING:
		g.user.Password += string(ebiten.InputChars())
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			go g.getToken()
			g.states.loginState = WAITING
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) {
			if len(g.user.Password) >= 1 {
				g.user.Password = g.user.Password[:len(g.user.Password)-1]
			}
		}
	case USERNAMETYPING:
		g.user.Username += string(ebiten.InputChars())
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			g.states.loginState = PASSWORDTYPING
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) {
			if len(g.user.Username) >= 1 {
				g.user.Username = g.user.Username[:len(g.user.Username)-1]
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
	//text.Draw(screen, g.player.Token, mplusNormalFont, 160, 80, color.White)
	for _, player := range g.players {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(player.x, player.y)
		screen.DrawImage(player.face, op)

	}
}

func (g *Game) drawLoginScreen(screen *ebiten.Image) {
	switch g.states.loginState {
	case USERNAMETYPING:
		text.Draw(screen, usernamelbl, mplusNormalFont, 20, 80, color.White)
		text.Draw(screen, g.user.Username, mplusNormalFont, 100, 80, color.White)
	case PASSWORDTYPING:
		text.Draw(screen, usernamelbl, mplusNormalFont, 20, 80, color.White)
		text.Draw(screen, g.user.Username, mplusNormalFont, 100, 80, color.White)
		text.Draw(screen, passwordlbl, mplusNormalFont, 20, 120, color.White)
		text.Draw(screen, g.user.Password, mplusNormalFont, 100, 120, color.White)
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
	//link := "https://api.backend.mama.sh/login"
	jsonStr, err := json.Marshal(g.user)
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
	g.user.Token = tkn.Token
	g.states.globalState = GAMEPLAY
	g.Unlock()
	go g.connect()

}

func (g *Game) connect() {
	var err error
	g.conn, err = setupConnection(g.user.Token)
	if err != nil {
		log.Fatal(err)
	}
	for {
		_, message, err := g.conn.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}
		if message[3] == ASSIGN {
			g.you = message[0]
			g.players[g.you] = createPlayer(int(g.you), &Position{
				x: float64(message[1]),
				y: float64(message[2]),
			})
			break
		}
	}
	for {
		_, message, err := g.conn.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}
		if message[3] == MOVE {
			g.players[message[0]].move(message[4], message[1], message[2])
			g.players[message[0]].x = float64(message[2])
			continue
		}
	}
}
