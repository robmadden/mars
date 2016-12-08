package rover

import (
	"github.com/robmadden/mars/position"
)

// This struct represents a rover
type Rover struct {
	Position *position.Position
}

// A Rover itself is dumb, it only knows it's location and
// Has no concept of other things in it's environment
func New(x int, y int, cardinal string) *Rover {
	return &Rover{
		Position: position.New(x, y, cardinal),
	}
}

func (r *Rover) Move(direction string) bool {
	if direction == "L" {
		return r.rotateLeft()
	} else if direction == "R" {
		return r.rotateRight()
	} else if direction == "M" {
		return r.maintain()
	}

	return false
}

// Given a cardinal direction, rotate 90 degrees "left" (counter clockwise)
func (r *Rover) rotateLeft() bool {
	if r.Position.Cardinal == "N" {
		r.Position.Cardinal = "W"
		return true
	} else if r.Position.Cardinal == "E" {
		r.Position.Cardinal = "N"
		return true
	} else if r.Position.Cardinal == "S" {
		r.Position.Cardinal = "E"
		return true
	} else if r.Position.Cardinal == "W" {
		r.Position.Cardinal = "S"
		return true
	}

	return false
}

// Given a cardinal direction, rotate 90 degrees "right" (clockwise)
func (r *Rover) rotateRight() bool {
	if r.Position.Cardinal == "N" {
		r.Position.Cardinal = "E"
		return true
	} else if r.Position.Cardinal == "E" {
		r.Position.Cardinal = "S"
		return true
	} else if r.Position.Cardinal == "S" {
		r.Position.Cardinal = "W"
		return true
	} else if r.Position.Cardinal == "W" {
		r.Position.Cardinal = "N"
		return true
	}

	return false
}

func (r *Rover) maintain() bool {
	if r.Position.Cardinal == "N" {
		r.Position.Y++
		return true
	} else if r.Position.Cardinal == "E" {
		r.Position.X++
		return true
	} else if r.Position.Cardinal == "S" {
		r.Position.Y--
		return true
	} else if r.Position.Cardinal == "W" {
		r.Position.X--
		return true
	}

	return false
}
