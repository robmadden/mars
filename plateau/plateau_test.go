package plateau_test

import (
	"testing"

	. "github.com/franela/goblin"
	"github.com/robmadden/mars/plateau"
	"github.com/robmadden/mars/rover"
)

func TestPlateau(t *testing.T) {
	g := Goblin(t)

	g.Describe("New", func() {
		var p *plateau.Plateau
		var r1 *rover.Rover
		var r2 *rover.Rover

		g.BeforeEach(func() {
			p = plateau.New(100, 100)
			r1 = rover.New(0, 0, "N")
			r2 = rover.New(0, 0, "N")
		})

		g.It("Should create a new Plateau with the proper boundaries ", func() {
			g.Assert(p.Topology != nil).IsTrue()
			g.Assert(p.Width).Equal(100)
			g.Assert(p.Height).Equal(100)
		})
	})

	g.Describe("Add", func() {
		var p *plateau.Plateau
		var r1 *rover.Rover
		var r2 *rover.Rover

		g.BeforeEach(func() {
			p = plateau.New(100, 100)
			r1 = rover.New(0, 1, "N")
			r2 = rover.New(0, 0, "N")
		})

		g.It("Should add a Rover to Topology(x, y) if no Rover exists at (x, y) ", func() {
			added := p.Add(r1)
			g.Assert(added).Equal(true)
			g.Assert(p.Topology["01"]).Equal(r1)
		})
	})

	g.Describe("Remove", func() {
		var p *plateau.Plateau
		var r1 *rover.Rover

		g.BeforeEach(func() {
			p = plateau.New(100, 100)
			r1 = rover.New(0, 0, "N")
			p.Add(r1)
		})

		g.It("Should remove a Rover from Topology(x, y) ", func() {
			removed := p.Remove(r1)
			g.Assert(removed).Equal(true)
		})
	})

	g.Describe("PositionIsValid", func() {
		var p *plateau.Plateau
		var r1 *rover.Rover
		var r2 *rover.Rover

		g.BeforeEach(func() {
			p = plateau.New(100, 100)
			r1 = rover.New(0, 0, "N")
			r2 = rover.New(0, 0, "N")
		})

		g.It("Should return true if the plateau is empty at (x, y) ", func() {
			empty := p.PositionIsValid(0, 0)
			g.Assert(empty).Equal(true)
		})

		g.It("Should return false if the plateau is not empty at (x, y) ", func() {
			p.Add(r1)
			empty := p.PositionIsValid(0, 0)
			g.Assert(empty).Equal(false)
		})
	})

	g.Describe("Move", func() {
		var p *plateau.Plateau
		g.BeforeEach(func() { p = plateau.New(10, 10) })

		g.It("will not let a rover fall off the plateau from (0, 0, W)'", func() {
			r := rover.New(0, 0, "W")
			p.Move(r, "M")
			g.Assert(r.Position.X).Equal(0)
			g.Assert(r.Position.Y).Equal(0)
			g.Assert(r.Position.Cardinal).Equal("W")
		})

		g.It("will not let a rover fall off the plateau from (0, 0, S)'", func() {
			r := rover.New(0, 0, "S")
			p.Move(r, "M")
			g.Assert(r.Position.X).Equal(0)
			g.Assert(r.Position.Y).Equal(0)
			g.Assert(r.Position.Cardinal).Equal("S")
		})

		g.It("will not let a rover fall off the plateau from (0, Y, W)'", func() {
			r := rover.New(0, 10, "W")
			p.Move(r, "M")
			g.Assert(r.Position.X).Equal(0)
			g.Assert(r.Position.Y).Equal(10)
			g.Assert(r.Position.Cardinal).Equal("W")
		})

		g.It("will not let a rover fall off the plateau from (0, Y, N)'", func() {
			r := rover.New(0, 10, "N")
			p.Move(r, "M")
			g.Assert(r.Position.X).Equal(0)
			g.Assert(r.Position.Y).Equal(10)
			g.Assert(r.Position.Cardinal).Equal("N")
		})

		g.It("will not let a rover fall off the plateau from (X, 0, E)'", func() {
			r := rover.New(10, 0, "E")
			p.Move(r, "M")
			g.Assert(r.Position.X).Equal(10)
			g.Assert(r.Position.Y).Equal(0)
			g.Assert(r.Position.Cardinal).Equal("E")
		})

		g.It("will not let a rover fall off the plateau from (X, 0, S)'", func() {
			r := rover.New(10, 0, "S")
			p.Move(r, "M")
			g.Assert(r.Position.X).Equal(10)
			g.Assert(r.Position.Y).Equal(0)
			g.Assert(r.Position.Cardinal).Equal("S")
		})

		g.It("will not let a rover fall off the plateau from (X, Y, E)'", func() {
			r := rover.New(10, 10, "E")
			p.Move(r, "M")
			g.Assert(r.Position.X).Equal(10)
			g.Assert(r.Position.Y).Equal(10)
			g.Assert(r.Position.Cardinal).Equal("E")
		})

		g.It("will not let a rover fall off the plateau from (X, Y, N)'", func() {
			r := rover.New(10, 10, "N")
			p.Move(r, "M")
			g.Assert(r.Position.X).Equal(10)
			g.Assert(r.Position.Y).Equal(10)
			g.Assert(r.Position.Cardinal).Equal("N")
		})

		g.It("Should not add a Rover to Topology(x, y) if a Rover exists there already ", func() {
			r1 := rover.New(0, 1, "N")
			r2 := rover.New(0, 0, "N")
			p.Add(r1)
			p.Add(r2)
			p.Move(r2, "M")
			g.Assert(r2.Position.X).Equal(0)
			g.Assert(r2.Position.Y).Equal(0)
			g.Assert(r2.Position.Cardinal).Equal("N")
		})
	})
}
