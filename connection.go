package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
	"os/signal"
)

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
	g.user.Token = tkn.AuthToken
	g.Unlock()
	go g.connect()

}
func setupConnection(token string) (c *websocket.Conn, err error) {
	//TODO Will be used to close the connection later
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	//u := "wss://api.backend.mama.sh/join"
	u := "ws://127.0.0.1:8056/join"
	log.Printf("connecting to %s", u)

	c, _, err = websocket.DefaultDialer.Dial(u, nil)
	token = fmt.Sprintf("Bearer %v", token)
	fmt.Println(token)
	message := append([]byte{0,0,0,0},token...)
	c.WriteMessage(websocket.BinaryMessage,message)
	return
}
