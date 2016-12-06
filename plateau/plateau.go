package plateau

import (
	"log"
	"strconv"

	"github.com/robmadden/mars/rover"
)

type Plateau struct {
	Grid   map[string]*rover.Rover
	Width  int
	Height int
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
		Grid:   make(map[string]*rover.Rover),
		Width:  w,
		Height: h,
	}
}

// Given an (x, y) coordinate, check if the coordinate is empty
// If a Rover doesn't exist at (x, y) spot, return true.
// If a Rover does exist at (x, y), return false
func (p *Plateau) IsEmptyAt(x int, y int) bool {
	key := p.hash(x, y)
	if p.Grid[key] == nil {
		return true
	} else {
		return false
	}
}

// Given an (x, y) coordinate, add a rover to Grid(x, y)
// If a Rover doesn't exist at (x, y) spot, add it and return true.
// If a Rover does exist at (x, y), don't add it and return false
func (p *Plateau) Add(r *rover.Rover) bool {
	if p.IsEmptyAt(r.Position.X, r.Position.Y) {
		key := p.hash(r.Position.X, r.Position.Y)
		p.Grid[key] = r
		return true
	} else {
		return false
	}
}
