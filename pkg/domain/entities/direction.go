package entities

import "fmt"

// Each direction is represented by an int value
const (
	North Direction = iota + 1
	East
	South
	West
)

type Direction int

// NewDirection return a new direction or an error based on the inputted direction
func NewDirection(direction string) (Direction, error) {
	switch direction {
	case "N":
		return North, nil
	case "E":
		return East, nil
	case "S":
		return South, nil
	case "W":
		return West, nil
	}
	return Direction(0), fmt.Errorf("invalid initial direction: %s", direction)
}
