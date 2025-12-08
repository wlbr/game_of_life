package main

import (
	"time"
)

func (b *Board) CalcNextGeneration(g *Board) {
	for y, l := range b.elements {
		for x := range l {
			neighboors := b.GetNeighboorCoords(x, y)
			nearPopulation := 0
			for _, n := range neighboors {
				if b.Get(n[0], n[1]) {
					nearPopulation++
				}
			}
			switch nearPopulation {
			case 2:
				g.Set(x, y, b.Get(x, y))
			case 3:
				g.Set(x, y, true)
			default:
				g.Set(x, y, false)
			}
		}
	}
}

func (b *Board) RunGameOfLife(generations int, sleep int) {
	g := NewBoardFromTemplate(b)

	for i := 1; i <= generations || generations < 1; i++ {
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		b.CalcNextGeneration(g)
		b, g = g, b
		b.printer.Update(b)
	}
}
