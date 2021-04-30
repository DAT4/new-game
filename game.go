package main

import (
	"context"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/colornames"
	"image"
	"image/color"
	"log"
	"nhooyr.io/websocket"
	"sync"
	"time"
)

type Game struct {
	sync.Mutex
	canMove bool
	user    *User
	conn    *websocket.Conn
	you     byte
	layers  [][]int
	players map[byte]*Player
	states  states
}

func (g *Game) moveActualPlayer() {
	if g.canMove {
		g.Lock()
		g.canMove = false
		g.Unlock()
		switch {
		case ebiten.IsKeyPressed(ebiten.KeyH):
			if g.players[g.you].x > 0 {
				g.conn.Write(context.Background(), websocket.MessageBinary, g.players[g.you].sendMove(LEFT))
			}
		case ebiten.IsKeyPressed(ebiten.KeyJ):
			if g.players[g.you].y < 29 {
				g.conn.Write(context.Background(), websocket.MessageBinary, g.players[g.you].sendMove(DOWN))
			}
		case ebiten.IsKeyPressed(ebiten.KeyK):
			if g.players[g.you].y > 0 {
				g.conn.Write(context.Background(), websocket.MessageBinary, g.players[g.you].sendMove(UP))
			}
		case ebiten.IsKeyPressed(ebiten.KeyL):
			if g.players[g.you].x < 29 {
				g.conn.Write(context.Background(), websocket.MessageBinary, g.players[g.you].sendMove(RIGHT))
			}
		default:
			g.Lock()
			g.canMove = true
			g.Unlock()
		}
	}
}

func (g *Game) movementTimer() {
	for {
		time.Sleep(200 * time.Millisecond)
		g.Lock()
		g.canMove = true
		g.Unlock()
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
		op.GeoM.Translate(float64(player.x*tileSize), float64(player.y*tileSize))
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

func (g *Game) connect() {
	var err error
	g.conn, err = setupConnection(g.user.Token)
	if err != nil {
		log.Fatal(err)
	}
	for {
		_, message, err := g.conn.Read(context.Background())
		fmt.Println(message)
		if err != nil {
			log.Fatal(err)
		}
		if message[3] == CREATE {
			g.you = message[0]
			fmt.Println(g.you)
			g.players[g.you] = createPlayer(g.you, &Position{
				x: int(message[1]),
				y: int(message[2]),
			})
			g.conn.Write(context.Background(), websocket.MessageBinary, g.players[g.you].assign())
			break
		}
	}
	g.states.globalState = GAMEPLAY
	for {
		_, message, err := g.conn.Read(context.Background())
		fmt.Println(message)
		if err != nil {
			log.Fatal(err)
		}
		player := message[0]
		switch message[3] {
		case MOVE:
			g.players[message[0]].move(message[4], message[1], message[2])
		case ASSIGN:
			g.players[player] = createPlayer(player, &Position{
				x: int(message[1]),
				y: int(message[2]),
			})
		}
	}
}
