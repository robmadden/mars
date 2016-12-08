package simulation

import (
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/robmadden/mars/plateau"
	"github.com/robmadden/mars/rover"
	"github.com/robmadden/mars/simio"
)

const debug = false
const delimiter = " "

// Error Messages
const GENERIC_PARSING_FAILED_MSG = "Failed to parse instructions properly, please check instructions.txt file."
const PLATEAU_CREATION_FAILED_MSG = "Unable to parse Plateau creation, please fix your instructions.txt file."
const ROVER_CREATION_FAILED_MSG = "Unable to parse Rover creation, please fix your instructions.txt file."

var p *plateau.Plateau

type Simulation struct {
	io *simio.SimIO
}

func New(io *simio.SimIO) (simulation *Simulation) {
	return &Simulation{
		io: io,
	}
}

func (s *Simulation) Run() {
	// Cherry pick the first line of the input and create a Plateau object
	p = s.CreatePlateau()

	// Parse the rest of the instructions
	s.ParseInstructions()

	// Print the last known location of each Rover
	s.io.PrintOutcome(p)
}

// Helper method to do basic error handling
// Mars is not graceful and will stop execution if input
// Is not in expected format
func (s *Simulation) CheckForFatalError(msg string, err error) {
	if err != nil && err != io.EOF {
		log.Fatalf("%s\nREASON: %s", msg, err)
	}
}

// Given a known plateau input line, create a plateau object
func (s *Simulation) CreatePlateau() *plateau.Plateau {
	line, err := s.io.ReadLine()
	s.CheckForFatalError(GENERIC_PARSING_FAILED_MSG, err)

	dimensions := strings.Split(line, delimiter)

	x, err := strconv.Atoi(dimensions[0])
	s.CheckForFatalError(PLATEAU_CREATION_FAILED_MSG, err)

	y, err := strconv.Atoi(dimensions[1])
	s.CheckForFatalError(PLATEAU_CREATION_FAILED_MSG, err)

	p = plateau.New(x, y)

	if debug {
		fmt.Printf("PLATEAU: (%x, %x)\n", p.Width, p.Height)
	}

	return p
}

// Parse Rover staring point and create a Rover object
func (s *Simulation) CreateRover(text string) *rover.Rover {
	position := strings.Split(text, delimiter)

	x, err := strconv.Atoi(position[0])
	s.CheckForFatalError(ROVER_CREATION_FAILED_MSG, err)

	y, err := strconv.Atoi(position[1])
	s.CheckForFatalError(ROVER_CREATION_FAILED_MSG, err)

	cardinal := position[2]
	cardinal = strings.Trim(cardinal, "\n")
	s.CheckForFatalError(ROVER_CREATION_FAILED_MSG, err)

	r := rover.New(x, y, cardinal)

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
func (s *Simulation) InstructRover(text string, r *rover.Rover) {
	if debug {
		fmt.Printf("INSTRUCTIONS: %s\n", text)
	}

	for _, instruction := range strings.Split(text, "") {
		if debug {
			fmt.Printf("MOVE: %s\n", instruction)
		}

		p.Move(r, instruction)
	}
}

// Iterate over the instructions file
func (s *Simulation) ParseInstructions() {
	line, err := s.io.ReadLine()
	for ok := true; ok; ok = (err != io.EOF) {
		// A Rover's starting point and Instructions
		r := s.CreateRover(line)

		// Now parse the associated instructions for the Rover
		line, err = s.io.ReadLine()
		s.CheckForFatalError(GENERIC_PARSING_FAILED_MSG, err)
		s.InstructRover(line, r)

		line, err = s.io.ReadLine()
	}
}
