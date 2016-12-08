package simio

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/robmadden/mars/plateau"
)

type SimIO struct {
	reader *bufio.Reader
}

func New(file *os.File) *SimIO {
	return &SimIO{
		reader: bufio.NewReader(file),
	}
}

// Helper method to read a line from the file and trim the trailing newline
func (io *SimIO) ReadLine() (string, error) {
	line, err := io.reader.ReadString('\n')
	line = strings.Trim(line, "\n")
	return line, err
}

// Iterate over the Plateau's Rovers and print their last known whereabouts
func (io *SimIO) PrintOutcome(plateau *plateau.Plateau) {
	for rover := range plateau.Topology {
		r := plateau.Topology[rover]
		if r != nil {
			fmt.Printf("%x %x %s\n", r.Position.X, r.Position.Y, r.Position.Cardinal)
		}
	}
}
