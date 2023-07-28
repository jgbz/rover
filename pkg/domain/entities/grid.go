package entities

type Grid []int

// NewGrid returns a new Grid
func NewGrid(x, y int) Grid {
	return []int{x, y}
}

// GetX returns the x position of the grid
func (g Grid) GetX() int {
	return g[0]
}

// GetY returns the y position of the grid
func (g Grid) GetY() int {
	return g[1]
}
