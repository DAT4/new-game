package main

import (
	"github.com/hajimehoshi/ebiten"
	"golang.org/x/image/font"
)

var (
	tilesImage      *ebiten.Image
	mplusNormalFont font.Face
)

//Global settings

//Screensize
const (
	screenWidth  = 240
	screenHeight = 240
)

//Tiles
const (
	tileSize = 16
	tileXNum = 25
)

//Strings
const (
	usernamelbl = `username: `
	passwordlbl = `password: `
	loading     = `loading...`
)

//STATES (ENUMS)

type state int
type states struct {
	globalState state
	loginState  state
}

//Global
const (
	LOGIN state = iota
	GAMEPLAY
)

//Login
const (
	USERNAMETYPING state = iota
	PASSWORDTYPING
	WAITING
)

//Directions
const (
	LEFT = iota
	RIGHT
	UP
	DOWN
)
