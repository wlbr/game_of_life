package main

import (
	"log" // Added log import
	"os"
	"runtime"
	"time"
)

func main() {
	defer os.Exit(0)
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
	//b.AddSuicide(30, 15)
	//b.AddFPentomino(50, 20)

	// New error handling for AddGlider
	if err := b.AddSuicide(25, 15); err != nil {
		log.Fatalf("Error adding glider pattern: %v", err)
	}

	b.RunGameOfLife(54, 100)
	time.Sleep(3 * time.Second)

	runtime.Goexit()
}
