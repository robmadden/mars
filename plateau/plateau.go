package plateau

import (
	"log"
	"strconv"

	"github.com/robmadden/mars/rover"
)

type Plateau struct {
	Width    int
	Height   int
	Topology map[string]*rover.Rover
}

// Give an (x, y) coordinate, return a unique value for the map
// Example (0, 0) => "00"
// Example (1, 0) => "10"
// Example (10, 0) => "100"
func (p *Plateau) hash(x int, y int) string {
	sX := strconv.Itoa(x)
	sY := strconv.Itoa(y)
	return sX + sY
}

func New(w int, h int) *Plateau {
	if w < 0 || h < 0 {
		log.Fatal("Invalid Plateau dimensions")
	}

	return &Plateau{
		Width:    w,
		Height:   h,
		Topology: make(map[string]*rover.Rover),
	}
}

// Satiate GetWidth Topology interface function
func (p *Plateau) GetWidth() int {
	return p.Width
}

// Satiate GetHeight Topology interface function
func (p *Plateau) GetHeight() int {
	return p.Height
}

// Given an (x, y) coordinate, check if the coordinate is empty
// If a Rover doesn't exist at (x, y) spot, return true.
// If a Rover does exist at (x, y), return false
func (p *Plateau) PositionIsValid(x int, y int) bool {
	key := p.hash(x, y)
	outOfBounds := false
	if x > p.Width || x < 0 || y > p.Height || y < 0 {
		outOfBounds = true
	}

	return p.Topology[key] == nil && !outOfBounds
}

// Given an (x, y) coordinate, add a rover to Topology(x, y)
// If a Rover doesn't exist at (x, y) spot, add it and return true.
// If a Rover does exist at (x, y), don't add it and return false
func (p *Plateau) Add(r *rover.Rover) bool {
	key := p.hash(r.Position.X, r.Position.Y)
	p.Topology[key] = r
	return true
}

// Remove a Rover from the Topology
func (p *Plateau) Remove(r *rover.Rover) bool {
	key := p.hash(r.Position.X, r.Position.Y)
	p.Topology[key] = nil
	return true
}

func (p *Plateau) Move(r *rover.Rover, instruction string) bool {
	// First we need to check that the instruction is valid
	// Rules:
	// A Rover is not allowed to move off the Topology
	// Not two Rovers can occupy the same location in the Topology

	// We only care about the "M" case for special cases
	if instruction == "M" {
		newX := r.Position.X
		newY := r.Position.Y

		if r.Position.Cardinal == "N" {
			newY++
		} else if r.Position.Cardinal == "E" {
			newX++
		} else if r.Position.Cardinal == "S" {
			newY--
		} else if r.Position.Cardinal == "W" {
			newX--
		}

		if p.PositionIsValid(newX, newY) {

			// Remove old reference using old coordinates
			p.Remove(r)

			// Move the Rover
			r.Move(instruction)

			// Read it to the Topology
			p.Add(r)

			return true
		} else {
			return false
		}
	} else if instruction == "L" || instruction == "R" {
		return r.Move(instruction)
	}

	return false

}
