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

func (u *User) getToken() {
	fmt.Println(u.Username)
	fmt.Println(u.Password)
	link := "http://localhost:8056/login"
	//link := "https://tmp.mama.sh/api/login"
	jsonStr, err := json.Marshal(u)
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

	u.token = token.Token
	u.state = LOGGEDIN
}
