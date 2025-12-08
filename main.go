package main

func main() {

	b := NewBoard(60, 35, &NCursesPrinter{}) // &SimplePrinter
	defer b.printer.Quit()
	//b.torus = true

	//b.AddRandom(0, 0)
	// b.AddBeehive(70, 2)
	// b.AddBlinker(40, 10)
	//b.AddGlider(0, 0)
	// b.AddGliderGun(0, 0)
	// b.AddGliderGun2(0, 0)
	//b.AddReflector(2, 2)
	b.AddSuicide(30, 15)
	//b.AddFPentomino(50, 20)

	b.RunGameOfLife(60, 100)
}
