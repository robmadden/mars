package position_test

import (
	"testing"

	"github.com/robmadden/mars/position"

	. "github.com/franela/goblin"
)

func TestPosition(t *testing.T) {
	g := Goblin(t)

	g.Describe("New", func() {
		g.It("Should create a new Position with the specified coordinates ", func() {
			p := position.New(0, 0, "N")
			g.Assert(p.X).Equal(0)
			g.Assert(p.Y).Equal(0)
			g.Assert(p.Cardinal).Equal("N")
		})
	})

	g.Describe("Update", func() {
		g.It("Should update an existing Position to the specified coordinates ", func() {
			p := position.New(0, 0, "N")
			g.Assert(p.X).Equal(0)
			g.Assert(p.Y).Equal(0)
			g.Assert(p.Cardinal).Equal("N")

			p.Update(1, 1, "S")
			g.Assert(p.X).Equal(1)
			g.Assert(p.Y).Equal(1)
			g.Assert(p.Cardinal).Equal("S")
		})
	})
}
