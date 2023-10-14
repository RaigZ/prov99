package prov99

import (
	"math/rand"
	"time"

	tl "github.com/JoelOtter/termloop"
)

var textInput TextInput
var game Game

type Game struct {
	Engine *tl.Game
	level  *tl.BaseLevel
	title  *tl.Text
	rv     int // Random value generated from 1-99
	pv     int // Current Player value
	att    int
	state  bool
}

func NewGame() *Game {
	game.Engine = tl.NewGame()
	game.Engine.Screen().SetFps(30)
	game.level = tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
	})
	selection := NewSelection(2, 1, 20, 1)
	textInput = TextInput{
		x:        4,
		y:        4,
		content:  "",
		maxWidth: 20,
	}
	enter := tl.NewText(3, 1, "ENTER", tl.ColorWhite, tl.ColorRed)
	clear := tl.NewText(3, 2, "CLEAR", tl.ColorWhite, tl.ColorRed)
	exit := tl.NewText(3, 3, "EXIT", tl.ColorWhite, tl.ColorRed)
	game.Engine.Screen().AddEntity(selection)
	game.Engine.Screen().AddEntity(&textInput)

	game.Engine.Screen().AddEntity(enter)
	game.Engine.Screen().AddEntity(clear)
	game.Engine.Screen().AddEntity(exit)

	game.Engine.Screen().SetLevel(game.level)
	return &game
}

func (g *Game) Start() {
	g.Context(&g.rv) // Display in-game prompts, set game.rv value to chosen rv
	g.Engine.Start()
}

func (g *Game) Context(rv *int) {
	g.title = tl.NewText(0, 0, "PROV99", tl.ColorGreen, tl.ColorBlack)
	*rv = chosenRV()
	// RV := tl.NewText(0, 5, strconv.Itoa(g.rv), tl.ColorGreen, tl.ColorBlack)

	g.Engine.Screen().AddEntity(g.title)
	// g.Engine.Screen().AddEntity(RV)
}

func (g *Game) Tick(e tl.Event) {}

func (g *Game) Draw(s *tl.Screen) {}

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

// State in which the player wins, which only occurs when game.pv == game.rv
func (g *Game) hasWon() bool {
	return g.state
}

func GamePrompt() string {
	return ("Choose a number between (1-99): ")
}
