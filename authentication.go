package main

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	State    int    `json:"-"`
	Token    string `json:"-"`
}

type jwt struct {
	Token string
}