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
			r1 = rover.New(100, 100, 0, 0, "N")
			r2 = rover.New(100, 100, 0, 0, "N")
		})

		g.It("Should create a new Plateau with the proper boundaries ", func() {
			g.Assert(p.Grid != nil).IsTrue()
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
			r1 = rover.New(100, 100, 0, 0, "N")
			r2 = rover.New(100, 100, 0, 0, "N")
		})

		g.It("Should add a Rover to Grid(x, y) if no Rover exists at (x, y) ", func() {
			added := p.Add(r1)
			g.Assert(added).Equal(true)
			g.Assert(p.Grid["00"]).Equal(r1)
		})

		g.It("Should not add a Rover to Grid(x, y) if a Rover exists there already ", func() {
			p.Add(r1)
			added := p.Add(r2)
			g.Assert(added).Equal(false)
			g.Assert(p.Grid["00"]).Equal(r1)
		})
	})

	g.Describe("isEmptyAt", func() {
		var p *plateau.Plateau
		var r1 *rover.Rover
		var r2 *rover.Rover

		g.BeforeEach(func() {
			p = plateau.New(100, 100)
			r1 = rover.New(100, 100, 0, 0, "N")
			r2 = rover.New(100, 100, 0, 0, "N")
		})

		g.It("Should return true if the plateau is empty at (x, y) ", func() {
			empty := p.IsEmptyAt(0, 0)
			g.Assert(empty).Equal(true)
		})

		g.It("Should return false if the plateau is not empty at (x, y) ", func() {
			p.Add(r1)
			empty := p.IsEmptyAt(0, 0)
			g.Assert(empty).Equal(false)
		})
	})

}
