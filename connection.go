package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"nhooyr.io/websocket"
	"os"
	"os/signal"
)

func (g *Game) getToken() {
	//link := "http://localhost/login"
	link := "https://api.backend.mama.sh/login"
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

	u := "wss://api.backend.mama.sh/join"
	//u := "ws://localhost/join"
	log.Printf("connecting to %s", u)

	c, _, err = websocket.Dial(context.Background(), u, nil)
	if err != nil {
		return nil, err
	}
	token = fmt.Sprintf("Bearer %v", token)
	fmt.Println(token)
	message := append([]byte{0, 0, 0, 0}, token...)
	err = c.Write(context.Background(), websocket.MessageBinary, message)
	return
}
