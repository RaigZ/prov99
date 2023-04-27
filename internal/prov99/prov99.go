package prov99

import (
	"math/rand"
	"strconv"
	"time"

	tl "github.com/JoelOtter/termloop"
)

// Want the game to represent a top level termloop app
type Game struct {
	Engine *tl.Game
	level  *tl.BaseLevel
	prompt *tl.Text
	rv     int // Random value generated from 1-99
	pv     int // Current Player value
	att    int // Each game, we are given 3 attempts, decrement by 1 each attempt has been made

	// Attempts can only be made if a player inputs a value
	//matrix *Matrix
	// Win and loss states require to create a new game when true

	//wonStreak  int // count total winnings
	//lossStreak int // count total losses
}

func NewGame() *Game {
	var game Game
	game.Engine = tl.NewGame()
	game.Engine.Screen().SetFps(30)
	game.level = tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
	})
	game.Engine.Screen().SetLevel(game.level)
	return &game
}

func (g *Game) Start() {
	g.Context(&g.rv) // Display in-game prompts, set game.rv value to chosen rv
	g.Engine.Start()
}

func (g *Game) Context(rv *int) {
	g.prompt = tl.NewText(0, 0, "PROV99", tl.ColorGreen, tl.ColorBlack)
	*rv = chosenRV()
	RV := tl.NewText(0, 5, strconv.Itoa(g.rv), tl.ColorGreen, tl.ColorBlack)

	g.Engine.Screen().AddEntity(g.prompt)
	g.Engine.Screen().AddEntity(RV)
}

func (g *Game) Tick(e tl.Event) {}

func (g *Game) Draw(s *tl.Screen) {

}

// Randomizes a value between 1-99
func randomizedValue(x *int) int {
	rand.Seed(time.Now().UnixNano())
	*x = rand.Intn(99) + 1
	return *x
}

func chosenRV() int {
	var rv int
	randomizedValue(&rv)
	return rv
}

// State in which the player wins, which only occurs when pv == rv
func (g *Game) hasWon() bool {
	return true
}

// State in which the player loses, which only occurs when att == 3
func (g *Game) hasLost() bool {
	return false
}

func (g *Game) attemptsLeft() int {
	return g.att
}

/*
Separated prompts for my simplicity of possibly
enabling certain events to occur depending on prompts
*/
func GamePrompt() string {
	return ("Choose a number between (1-99): ")
}

/*
func (g *Game) AdjustGameStreaks(state bool) {
	switch state {
	case g.hasWon():
		g.wonStreak = g.wonStreak + 1
	case g.hasLost():
		g.lossStreak = g.lossStreak + 1
	}
}
*/

/*"Attempts() + Choose a number between (1-99)" + "*/
/*
---
	IF (pv == rv) THEN {
		hasWon(); generate newGame session
		wonStreak = wonStreak + 1
	}
	ELSE IF (att == 3) THEN {
		hasLost(); generate newGame session
		lossStreak = lossStreak + 1
	}
	ELSE IF (pv != rv) THEN
		att = att + 1
*/
