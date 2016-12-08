package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/robmadden/mars/plateau"
	"github.com/robmadden/mars/rover"
)

// Always expect instructions in a file named "instructions.txt"
const instructions = "instructions.txt"
const debug = false

const delimiter = " "

// Error Messages
const GENERIC_PARSING_FAILED_MSG = "Failed to parse instructions properly, please check instructions.txt file."
const PLATEAU_CREATION_FAILED_MSG = "Unable to parse Plateau creation, please fix your instructions.txt file."
const ROVER_CREATION_FAILED_MSG = "Unable to parse Rover creation, please fix your instructions.txt file."

var p *plateau.Plateau

// Helper method to do basic error handling
// Mars is not graceful and will stop execution if input
// Is not in expected format
func checkForFatalError(msg string, err error) {
	if err != nil && err != io.EOF {
		log.Fatalf("%s\nREASON: %s", msg, err)
	}
}

// Helper method to read a line from the file and trim the trailing newline
func readLine(reader *bufio.Reader) (string, error) {
	line, err := reader.ReadString('\n')
	line = strings.Trim(line, "\n")
	return line, err
}

// Given a known plateau input line, create a plateau object
func createPlateau(text string) {
	dimensions := strings.Split(text, delimiter)

	x, err := strconv.Atoi(dimensions[0])
	checkForFatalError(PLATEAU_CREATION_FAILED_MSG, err)

	y, err := strconv.Atoi(dimensions[1])
	checkForFatalError(PLATEAU_CREATION_FAILED_MSG, err)

	p = plateau.New(x, y)

	if debug {
		fmt.Printf("PLATEAU: (%x, %x)\n", p.Width, p.Height)
	}
}

// Parse Rover staring point and create a Rover object
func createRover(text string) *rover.Rover {
	position := strings.Split(text, delimiter)

	x, err := strconv.Atoi(position[0])
	checkForFatalError(ROVER_CREATION_FAILED_MSG, err)

	y, err := strconv.Atoi(position[1])
	checkForFatalError(ROVER_CREATION_FAILED_MSG, err)

	cardinal := position[2]
	cardinal = strings.Trim(cardinal, "\n")
	checkForFatalError(ROVER_CREATION_FAILED_MSG, err)

	r := rover.New(p.Width, p.Height, x, y, cardinal)

	// Add the Rover to the Plateau
	added := p.Add(r)
	if debug {
		if added {
			// Rover added
			fmt.Printf("ROVER ADDED (%x, %x, %s)\n", r.Position.X, r.Position.Y, r.Position.Cardinal)
		} else {
			// Rover not added
			fmt.Printf("ROVER SKIPPED (%x, %x, %s)\n", r.Position.X, r.Position.Y, r.Position.Cardinal)
		}
	}

	return r
}

// Parse Rover instructions and feed them to Rover object
func instructRover(text string, r *rover.Rover) {
	if debug {
		fmt.Printf("INSTRUCTIONS: %s\n", text)
	}

	for _, instruction := range strings.Split(text, "") {
		if debug {
			fmt.Printf("MOVE: %s\n", instruction)
		}
		r.Move(instruction)
	}
}

// Iterate over the instructions file
func parseInstructions(reader *bufio.Reader) {
	line, err := readLine(reader)
	for ok := true; ok; ok = (err != io.EOF) {
		// A Rover's starting point and Instructions
		r := createRover(line)

		// Now parse the associated instructions for the Rover
		line, err = readLine(reader)
		checkForFatalError(GENERIC_PARSING_FAILED_MSG, err)
		instructRover(line, r)

		line, err = readLine(reader)
	}
}

// Iterate over the Plateau's Rovers and print their last known whereabouts
func printOutcome() {
	for rover := range p.Grid {
		r := p.Grid[rover]
		fmt.Printf("%x %x %s\n", r.Position.X, r.Position.Y, r.Position.Cardinal)
	}
}

func main() {
	file, err := os.Open(instructions)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a new Reader to do custom line parsing
	reader := bufio.NewReader(file)

	// Cherry pick the first line of the input and create a Plateau object
	line, err := readLine(reader)
	checkForFatalError(GENERIC_PARSING_FAILED_MSG, err)
	createPlateau(line)

	// Parse the rest of the instructions
	parseInstructions(reader)

	// Print the last known location of each Rover after the instructions
	// were consumed
	printOutcome()
}
