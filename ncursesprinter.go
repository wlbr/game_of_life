package main

import (
	"log"
	"time"

	"github.com/gdamore/tcell/v3"
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
		p.screen.Put(col, y1, string(tcell.RuneHLine), style)
		p.screen.Put(col, y2, string(tcell.RuneHLine), style)
	}
	for row := y1 + 1; row < y2; row++ {
		p.screen.Put(x1, row, string(tcell.RuneVLine), style)
		p.screen.Put(x2, row, string(tcell.RuneVLine), style)
	}
	p.screen.Put(x1, y1, string(tcell.RuneULCorner), style)
	p.screen.Put(x2, y1, string(tcell.RuneURCorner), style)
	p.screen.Put(x1, y2, string(tcell.RuneLLCorner), style)
	p.screen.Put(x2, y2, string(tcell.RuneLRCorner), style)

	p.screen.Show()
}

func (p *NCursesPrinter) Quit() {
	p.screen.Put(1, 1, "X", tcell.StyleDefault.Background(tcell.ColorLightGray).Foreground(tcell.ColorBlack))
	p.screen.Show()
	time.Sleep(1 * time.Second)
	p.screen.Fini()
}

func (p *NCursesPrinter) Update(b *Board) {
	for y, l := range b.elements {
		for x, e := range l {
			if e {
				p.screen.Put(x+1, y+1, " ", tcell.StyleDefault.Background(tcell.ColorLightGray).Foreground(tcell.ColorBlack))
			} else {
				p.screen.Put(x+1, y+1, " ", tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorBlack))
			}

		}
	}
	p.screen.Show()
}
