package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	LOGGEDOUT = iota
	LOGGEDIN
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	state    int    `json:"-"`
	token    string `json:"-"`
}

type jwt struct {
	Token string
}

func (g *Game) getToken() {
	link := "http://localhost:8056/login"
	//link := "https://tmp.mama.sh/api/login"
	jsonStr, err := json.Marshal(g.player)
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

	var token jwt
	err = json.NewDecoder(resp.Body).Decode(&token)
	if err != nil {
		fmt.Println(err)
		return
	}


	g.Lock()
	g.player.token = token.Token
	g.states.globalState = GAMEPLAY
	g.Unlock()
}
