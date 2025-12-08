package main

type Printer interface {
	Init(int, int)
	Quit()
	Update(b *Board)
}

type Board struct {
	dimx, dimy int
	elements   [][]bool
	torus      bool
	printer    Printer
}

func NewBoard(dimx, dimy int, printer Printer) *Board {
	b := &Board{dimx: dimx, dimy: dimy, torus: false}
	b.elements = make([][]bool, dimy)
	for y := range b.elements {
		b.elements[y] = make([]bool, dimx)
	}
	b.printer = printer
	b.printer.Init(b.dimx, b.dimy)
	return b
}

func NewTorus(dimx, dimy int, printer Printer) *Board {
	b := NewBoard(dimx, dimy, printer)
	b.torus = true
	return b
}

func NewBoardFromTemplate(b *Board) *Board {
	n := NewBoard(b.dimx, b.dimy, b.printer)
	n.torus = b.torus
	n.printer = b.printer
	return n
}

func (b *Board) String() string {
	result := ""
	for _, l := range b.elements {
		for _, e := range l {
			if e {
				result = result + "x "
			} else {
				result = result + "  "
			}
		}
		result = result + "\n"
	}
	return result
}

func (b *Board) Get(x, y int) bool {
	return b.elements[y][x]
}

func (b *Board) Set(x, y int, v bool) {
	b.elements[y][x] = v
}

func (b *Board) GetNeighboorCoords(x, y int) (n [][]int) {
	candidates := [][]int{{x - 1, y - 1}, {x, y - 1}, {x + 1, y - 1}, {x - 1, y}, {x, y}, {x + 1, y}, {x - 1, y + 1}, {x, y + 1}, {x + 1, y + 1}}
	for _, c := range candidates {
		if !(c[0] == x && c[1] == y) {
			if !b.torus {
				if c[0] >= 0 && c[0] < b.dimx && c[1] >= 0 && c[1] < b.dimy {
					n = append(n, c)
				}
			} else {
				if c[0] < 0 {
					c[0] = b.dimx - 1
				} else {
					if c[0] == b.dimx {
						c[0] = 0
					}
				}
				if c[1] < 0 {
					c[1] = b.dimy - 1
				} else {
					if c[1] == b.dimy {
						c[1] = 0
					}
				}
				n = append(n, c)
			}
		}
	}
	return n
}
