package position

type Position struct {
	X        int
	Y        int
	Cardinal string
}

func New(x int, y int, cardinal string) *Position {
	return &Position{
		X:        x,
		Y:        y,
		Cardinal: cardinal,
	}
}

func (p *Position) Update(x int, y int, cardinal string) {
	p.X = x
	p.Y = y
	p.Cardinal = cardinal
}
