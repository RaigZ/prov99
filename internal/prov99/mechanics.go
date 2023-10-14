package prov99

import (
	"fmt"
	"strconv"
	"unicode"

	tl "github.com/JoelOtter/termloop"
)

type Selection struct {
	x, y int
	*tl.Entity
}

func NewSelection(x, y, w, h int) *Selection {
	s := &Selection{
		x:      x,
		y:      y,
		Entity: tl.NewEntity(x, y, w, h),
	}
	s.SetCell(0, 0, &tl.Cell{Fg: tl.ColorBlack, Bg: tl.ColorRed, Ch: ' '})
	for i := 1; i < w; i++ {
		s.SetCell(i, 0, &tl.Cell{Fg: tl.ColorBlack, Bg: tl.ColorRed, Ch: ' '})
	}
	for i := 1; i < h; i++ {
		s.SetCell(0, i, &tl.Cell{Fg: tl.ColorBlack, Bg: tl.ColorRed, Ch: ' '})
	}
	return s
}

func (s *Selection) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		switch event.Key {
		case tl.KeyArrowUp:
			if s.y > 1 {
				s.y--
			}
		case tl.KeyArrowDown:
			if s.y < 3 {
				s.y++
			}
		case tl.KeyEnter:
			if game.att <= 2 && !game.hasWon() {
				switch s.y {
				case 1:
					game.att += 1
					// fmt.Printf("\033[1K\r" + textInput.GetContent())
				case 2:
					textInput.content = ""
				case 3:
					panic("EXIT")
				}
			} else if game.hasWon() {
				panic("WON")
			} else {
				panic("LOSE")
			}
		}
		s.SetPosition(s.x, s.y)
	}
}

type TextInput struct {
	x, y     int
	content  string
	maxWidth int
}

func (t *TextInput) Draw(screen *tl.Screen) {
	for i, char := range t.content {
		screen.RenderCell(t.x+i, t.y, &tl.Cell{Ch: char})
	}
}

func (t *TextInput) GetContent() string {
	return t.content
}

func (t *TextInput) DeleteCharacterAt(position int) {
	if position >= 0 && position < len(t.content) {
		t.content = t.content[:position] + t.content[position+1:]
	}
}

func (t *TextInput) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		switch event.Key {
		case tl.KeyEnter:
			game.pv, _ = strconv.Atoi(t.GetContent())
			t.content = ""
			if game.pv != game.rv {
				fmt.Printf("\033[1K\r"+"att: %d, pv: %d, rv: %d", game.att, game.pv, game.rv)
				if game.att == 3 {
					RV := tl.NewText(0, 5, strconv.Itoa(game.rv), tl.ColorGreen, tl.ColorBlack)
					game.Engine.Screen().AddEntity(RV)
				}
			} else if game.pv == game.rv {
				game.state = true
				fmt.Printf("\033[1K\r"+"CORRECT: att: %d, pv: %d, rv: %d", game.att, game.pv, game.rv)
				RV := tl.NewText(0, 5, strconv.Itoa(game.rv), tl.ColorGreen, tl.ColorBlack)
				game.Engine.Screen().AddEntity(RV)
			}
		case tl.KeyArrowUp:
		case tl.KeyArrowDown:
		default:
			if len(t.content) < t.maxWidth {
				if !unicode.IsSpace(event.Ch) {
					t.content += string(event.Ch)
				}
			}
		}
	}
}
