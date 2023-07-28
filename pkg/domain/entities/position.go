package entities

import "fmt"

type position struct {
	direction Direction
	X         int
	Y         int
}

// UpdateDirection updates the current facing direction from an int value(1 or -1)
// since each direction is represented by an int value only the values defined in the constants section are valid.
// There is two validations, if the domain is facing north(1) and turns left(-1), it will be facing west(4)
// the other one is if the domain is facing west(4) and turns right(1), then it will be facing north(1)
func (p *position) UpdateDirection(i int) {
	currentDirection := int(p.direction)
	if (currentDirection + i) > 4 {
		currentDirection = 0
		p.direction = Direction(currentDirection + i)
		return
	}

	if (currentDirection + i) < 1 {
		currentDirection = 5
		p.direction = Direction(currentDirection + i)
		return
	}

	p.direction = Direction(currentDirection + i)

}

// GetDirection returns a string value of the direction.
func (p *position) GetDirection() string {
	switch p.direction {
	case North:
		return "N"
	case East:
		return "E"
	case South:
		return "S"
	case West:
		return "W"
	default:
		return "unknown"
	}
}

// GetCoordinates output the current x and y positions and the facing direction.
func (p *position) GetCoordinates() string {
	return fmt.Sprintf("%v %v %s", p.X, p.Y, p.GetDirection())
}
