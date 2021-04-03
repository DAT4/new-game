package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func setupConnection(token string) (c *websocket.Conn, err error) {
	//TODO Will be used to close the connection later
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := "wss://api.backend.mama.sh/join"
	//u := "ws://127.0.0.1:8056/join"
	log.Printf("connecting to %s", u)

	header := http.Header{}
	header.Add("Authorization", "bearer "+token)
	c, _, err = websocket.DefaultDialer.Dial(u, header)
	return
}
