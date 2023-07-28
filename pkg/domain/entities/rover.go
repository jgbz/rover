package entities

type Rover struct {
	Pos          *position
	Instructions []rune
}

// NewRover returns a new Rover entity
func NewRover(initialX, initialY int, direction string, instructions []rune) (*Rover, error) {
	dir, err := NewDirection(direction)
	if err != nil {
		return nil, err
	}

	return &Rover{
		Pos: &position{
			X:         initialX,
			Y:         initialY,
			direction: dir,
		},
		Instructions: instructions,
	}, nil
}

// Explore runs all the inputted instructions from a single Rover and outputs it's final coordinates
func (r *Rover) Explore(_ Grid) (string, error) {

	for _, v := range r.Instructions {
		switch v {
		case 'L':
			r.Pos.UpdateDirection(-1)
		case 'R':
			r.Pos.UpdateDirection(1)
		case 'M':
			switch r.Pos.GetDirection() {
			case "N":
				r.Pos.Y = r.Pos.Y + 1
			case "E":
				r.Pos.X = r.Pos.X + 1
			case "S":
				r.Pos.Y = r.Pos.Y - 1
			case "W":
				r.Pos.X = r.Pos.X - 1
			}
		}
	}

	return r.Pos.GetCoordinates(), nil

}
