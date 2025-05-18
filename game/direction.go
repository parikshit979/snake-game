package game

// Direction represents a direction in the game.
// It is defined by a point with x and y coordinates.
type Direction struct {
	point Point
}

// NewDirection creates a new direction with the given x and y coordinates.
func NewDirection(x, y int) Direction {
	return Direction{NewPoint(x, y)}
}

// Up is the direction for moving up.
func Up() Direction {
	return NewDirection(0, -1)
}

// Down is the direction for moving down.
func Down() Direction {
	return NewDirection(0, 1)
}

// Left is the direction for moving left.
func Left() Direction {
	return NewDirection(-1, 0)
}

// Right is the direction for moving right.
func Right() Direction {
	return NewDirection(1, 0)
}
