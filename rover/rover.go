package rover

import (
	"github.com/robmadden/mars/position"
)

// This struct represents a rover
type Rover struct {
	plateauWidth  int
	plateauHeight int
	Position      *position.Position
}

func New(width int, height int, x int, y int, cardinal string) *Rover {
	return &Rover{
		plateauWidth:  width,
		plateauHeight: height,
		Position:      position.New(x, y, cardinal),
	}
}

func (r *Rover) Move(direction string) {
	if direction == "L" {
		r.rotateLeft()
	} else if direction == "R" {
		r.rotateRight()
	} else if direction == "M" {
		r.maintain()
	}
}

// Given a cardinal direction, rotate 90 degrees "left" (counter clockwise)
func (r *Rover) rotateLeft() {
	if r.Position.Cardinal == "N" {
		r.Position.Cardinal = "W"
	} else if r.Position.Cardinal == "E" {
		r.Position.Cardinal = "N"
	} else if r.Position.Cardinal == "S" {
		r.Position.Cardinal = "E"
	} else if r.Position.Cardinal == "W" {
		r.Position.Cardinal = "S"
	}
}

// Given a cardinal direction, rotate 90 degrees "right" (clockwise)
func (r *Rover) rotateRight() {
	if r.Position.Cardinal == "N" {
		r.Position.Cardinal = "E"
	} else if r.Position.Cardinal == "E" {
		r.Position.Cardinal = "S"
	} else if r.Position.Cardinal == "S" {
		r.Position.Cardinal = "W"
	} else if r.Position.Cardinal == "W" {
		r.Position.Cardinal = "N"
	}
}

func (r *Rover) maintain() {
	if r.Position.Cardinal == "N" {
		if r.Position.Y < r.plateauHeight {
			r.Position.Y++
		}
	} else if r.Position.Cardinal == "E" {
		if r.Position.X < r.plateauWidth {
			r.Position.X++
		}
	} else if r.Position.Cardinal == "S" {
		if r.Position.Y > 0 {
			r.Position.Y--
		}
	} else if r.Position.Cardinal == "W" {
		if r.Position.X > 0 {
			r.Position.X--
		}
	}
}
