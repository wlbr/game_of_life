package main

import (
	"log"

	"github.com/gdamore/tcell"
)

type NCursesPrinter struct {
	screen tcell.Screen
}

func (p *NCursesPrinter) Init(xdim, ydim int) {
	var err error
	p.screen, err = tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := p.screen.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	// Set default text style
	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorLightGray)
	p.screen.SetStyle(defStyle)

	// Clear screen
	p.screen.Clear()

	// Draw borders
	x1 := 0
	y1 := 0
	x2 := x1 + xdim + 1
	y2 := y1 + ydim + 1
	style := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)
	for col := x1; col <= x2; col++ {
		p.screen.SetContent(col, y1, tcell.RuneHLine, nil, style)
		p.screen.SetContent(col, y2, tcell.RuneHLine, nil, style)
	}
	for row := y1 + 1; row < y2; row++ {
		p.screen.SetContent(x1, row, tcell.RuneVLine, nil, style)
		p.screen.SetContent(x2, row, tcell.RuneVLine, nil, style)
	}
	p.screen.SetContent(x1, y1, tcell.RuneULCorner, nil, style)
	p.screen.SetContent(x2, y1, tcell.RuneURCorner, nil, style)
	p.screen.SetContent(x1, y2, tcell.RuneLLCorner, nil, style)
	p.screen.SetContent(x2, y2, tcell.RuneLRCorner, nil, style)

	p.screen.Show()
}

func (p *NCursesPrinter) Quit() {
	p.screen.Fini()
}

func (p *NCursesPrinter) Update(b *Board) {
	for y, l := range b.elements {
		for x, e := range l {
			if e {
				p.screen.SetContent(x+1, y+1, ' ', nil, tcell.StyleDefault.Background(tcell.ColorLightGray).Foreground(tcell.ColorBlack))
			} else {
				p.screen.SetContent(x+1, y+1, ' ', nil, tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorBlack))
			}

		}
	}
	p.screen.Show()
}
