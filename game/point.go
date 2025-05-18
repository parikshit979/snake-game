package game

// Point represents a point in 2D space.
// It is defined by its x and y coordinates.
type Point struct {
	X int
	Y int
}

// NewPoint creates a new point with the given x and y coordinates.
func NewPoint(x, y int) Point {
	return Point{X: x, Y: y}
}
