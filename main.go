package main

import (
	"log"
	"os"

	"github.com/robmadden/mars/simio"
	"github.com/robmadden/mars/simulation"
)

// Always expect instructions in a file named "instructions.txt"
const instructions = "instructions.txt"

func main() {
	file, err := os.Open(instructions)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a new Reader to do custom line parsing
	simio := simio.New(file)
	simulation := simulation.New(simio)
	simulation.Run()
}
