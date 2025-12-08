package main

import "fmt"

type SimplePrinter struct {
}

func (p *SimplePrinter) Init(int, int) {
}

func (p *SimplePrinter) Quit() {
}

func (p *SimplePrinter) Update(b *Board) {
	fmt.Println()
	fmt.Println(b)
}
