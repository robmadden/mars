package rover_test

import (
	"testing"

	. "github.com/franela/goblin"
	"github.com/robmadden/mars/rover"
)

// Input
// 5 5              (upper right coordinates of plateau)
// 1 2 N            (first rover starting position)
// LMLMLMLMM        (first rover instructions)
// 3 3 E            (second rover starting position)
// MMRMMRMRRM       (second rover instructions)

// Output
// 1 3 N            (first rover ending coordinates)
// 5 1 E            (second rover ending coordinates)

func TestRover(t *testing.T) {
	g := Goblin(t)

	g.Describe("New", func() {
		g.It("Should create a new Rover with the specified position ", func() {
			r := rover.New(10, 10, 0, 0, "N")

			g.Assert(r.Position.X).Equal(0)
			g.Assert(r.Position.Y).Equal(0)
			g.Assert(r.Position.Cardinal).Equal("N")
		})
	})

	// The possible letters are 'L', 'R' and 'M'.
	// 'L' and 'R' makes the rover spin 90 degrees left or right respectively,
	// without moving from its current spot.
	// 'M' means move forward one grid point, and maintain the same heading.

	g.Describe("Move('L')", func() {
		g.It("it spins a Rover facing North to face West ", func() {
			r := rover.New(10, 10, 0, 0, "N")
			r.Move("L")
			g.Assert(r.Position.X).Equal(0)
			g.Assert(r.Position.Y).Equal(0)
			g.Assert(r.Position.Cardinal).Equal("W")
		})

		g.It("it spins a Rover facing East to face North ", func() {
			r := rover.New(10, 10, 0, 0, "E")
			r.Move("L")
			g.Assert(r.Position.X).Equal(0)
			g.Assert(r.Position.Y).Equal(0)
			g.Assert(r.Position.Cardinal).Equal("N")
		})

		g.It("it spins a Rover facing South to face East ", func() {
			r := rover.New(10, 10, 0, 0, "S")
			r.Move("L")
			g.Assert(r.Position.X).Equal(0)
			g.Assert(r.Position.Y).Equal(0)
			g.Assert(r.Position.Cardinal).Equal("E")
		})

		g.It("it spins a Rover facing West to face South ", func() {
			r := rover.New(10, 10, 0, 0, "W")
			r.Move("L")
			g.Assert(r.Position.X).Equal(0)
			g.Assert(r.Position.Y).Equal(0)
			g.Assert(r.Position.Cardinal).Equal("S")
		})
	})

	g.Describe("Move('R')", func() {
		g.It("it spins a Rover facing North to face East ", func() {
			r := rover.New(10, 10, 0, 0, "N")
			r.Move("R")
			g.Assert(r.Position.X).Equal(0)
			g.Assert(r.Position.Y).Equal(0)
			g.Assert(r.Position.Cardinal).Equal("E")
		})

		g.It("it spins a Rover facing East to face South ", func() {
			r := rover.New(10, 10, 0, 0, "E")
			r.Move("R")
			g.Assert(r.Position.X).Equal(0)
			g.Assert(r.Position.Y).Equal(0)
			g.Assert(r.Position.Cardinal).Equal("S")
		})

		g.It("it spins a Rover facing South to face West ", func() {
			r := rover.New(10, 10, 0, 0, "S")
			r.Move("R")
			g.Assert(r.Position.X).Equal(0)
			g.Assert(r.Position.Y).Equal(0)
			g.Assert(r.Position.Cardinal).Equal("W")
		})

		g.It("it spins a Rover facing West to face North ", func() {
			r := rover.New(10, 10, 0, 0, "W")
			r.Move("R")
			g.Assert(r.Position.X).Equal(0)
			g.Assert(r.Position.Y).Equal(0)
			g.Assert(r.Position.Cardinal).Equal("N")
		})
	})

	g.Describe("Move('M')", func() {
		g.It("will not let a rover fall off the plateau from (0, 0, W)'", func() {
			r := rover.New(10, 10, 0, 0, "W")
			r.Move("M")
			g.Assert(r.Position.X).Equal(0)
			g.Assert(r.Position.Y).Equal(0)
			g.Assert(r.Position.Cardinal).Equal("W")
		})

		g.It("will not let a rover fall off the plateau from (0, 0, S)'", func() {
			r := rover.New(10, 10, 0, 0, "S")
			r.Move("M")
			g.Assert(r.Position.X).Equal(0)
			g.Assert(r.Position.Y).Equal(0)
			g.Assert(r.Position.Cardinal).Equal("S")
		})

		g.It("will not let a rover fall off the plateau from (0, Y, W)'", func() {
			r := rover.New(10, 10, 0, 10, "W")
			r.Move("M")
			g.Assert(r.Position.X).Equal(0)
			g.Assert(r.Position.Y).Equal(10)
			g.Assert(r.Position.Cardinal).Equal("W")
		})

		g.It("will not let a rover fall off the plateau from (0, Y, N)'", func() {
			r := rover.New(10, 10, 0, 10, "N")
			r.Move("M")
			g.Assert(r.Position.X).Equal(0)
			g.Assert(r.Position.Y).Equal(10)
			g.Assert(r.Position.Cardinal).Equal("N")
		})

		g.It("will not let a rover fall off the plateau from (X, 0, E)'", func() {
			r := rover.New(10, 10, 10, 0, "E")
			r.Move("M")
			g.Assert(r.Position.X).Equal(10)
			g.Assert(r.Position.Y).Equal(0)
			g.Assert(r.Position.Cardinal).Equal("E")
		})

		g.It("will not let a rover fall off the plateau from (X, 0, S)'", func() {
			r := rover.New(10, 10, 10, 0, "S")
			r.Move("M")
			g.Assert(r.Position.X).Equal(10)
			g.Assert(r.Position.Y).Equal(0)
			g.Assert(r.Position.Cardinal).Equal("S")
		})

		g.It("will not let a rover fall off the plateau from (X, Y, E)'", func() {
			r := rover.New(10, 10, 10, 10, "E")
			r.Move("M")
			g.Assert(r.Position.X).Equal(10)
			g.Assert(r.Position.Y).Equal(10)
			g.Assert(r.Position.Cardinal).Equal("E")
		})

		g.It("will not let a rover fall off the plateau from (X, Y, N)'", func() {
			r := rover.New(10, 10, 10, 10, "N")
			r.Move("M")
			g.Assert(r.Position.X).Equal(10)
			g.Assert(r.Position.Y).Equal(10)
			g.Assert(r.Position.Cardinal).Equal("N")
		})

		g.It("it maintains a heading of North and moves it forward one grid point ", func() {
			r := rover.New(10, 10, 1, 1, "N")
			r.Move("M")
			g.Assert(r.Position.X).Equal(1)
			g.Assert(r.Position.Y).Equal(2)
			g.Assert(r.Position.Cardinal).Equal("N")
		})

		g.It("it maintains a heading of East and moves it forward one grid point ", func() {
			r := rover.New(10, 10, 1, 1, "E")
			r.Move("M")
			g.Assert(r.Position.X).Equal(2)
			g.Assert(r.Position.Y).Equal(1)
			g.Assert(r.Position.Cardinal).Equal("E")
		})

		g.It("it maintains a heading of South and moves it forward one grid point ", func() {
			r := rover.New(10, 10, 1, 1, "S")
			r.Move("M")
			g.Assert(r.Position.X).Equal(1)
			g.Assert(r.Position.Y).Equal(0)
			g.Assert(r.Position.Cardinal).Equal("S")
		})

		g.It("it maintains a heading of West and moves it forward one grid point ", func() {
			r := rover.New(10, 10, 1, 1, "W")
			r.Move("M")
			g.Assert(r.Position.X).Equal(0)
			g.Assert(r.Position.Y).Equal(1)
			g.Assert(r.Position.Cardinal).Equal("W")
		})
	})
}
