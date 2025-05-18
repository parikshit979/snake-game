package game

import (
	colorUtil "image/color"
	"math/rand"
)

// Food represents a food item in the game.
// It is defined by its position, color, size, type, score, and maximum score.
type Food struct {
	// Position of the food in the game.
	Position *Point
	// Color of the food.
	Color colorUtil.Color
	// Type of the food.
	Type string
	// Score value of the food.
	Score int
}

// NewFood creates a new food instance.
func NewFood() *Food {
	return &Food{}
}

// Move moves the food to a new position.
// It updates the position of the food in the game.
func (f *Food) SpawnFood() {
	point := NewPoint(rand.Intn(ScreenWidth/GridSize),
		rand.Intn(ScreenHeight/GridSize))
	// Generate a new random position for the food
	f.Position = &point

	// Set the color, size, type, and score of the food
	f.Color = colorUtil.RGBA{255, 0, 0, 255} // Red color
	f.Type = "Apple"                         // Type of food
	f.Score = 1
}

// ChangeColor changes the color of the food.
// It updates the color of the food in the game.
func (f *Food) ChangeColor(newColor colorUtil.Color) {
	f.Color = newColor
}

// ChangeType changes the type of the food.
// It updates the type of the food in the game.
func (f *Food) ChangeType(newType string) {
	f.Type = newType
}

// ChangeScore changes the score value of the food.
// It updates the score value of the food in the game.
func (f *Food) ChangeScore(newScore int) {
	f.Score = newScore
}

// IsEaten checks if the food is eaten by the snake.
// It returns true if the food is eaten, false otherwise.
func (f *Food) IsEaten(snake Snake) bool {
	// Check if the snake's head is at the same position as the food
	return f.GetPosition().X == snake.GetHead().X &&
		f.GetPosition().Y == snake.GetHead().Y
}

// GetPosition returns the position of the food.
// It returns the position of the food in the game.
func (f *Food) GetPosition() *Point {
	return f.Position
}

// GetScore returns the score value of the food.
// It returns the score value of the food in the game.
func (f *Food) GetScore() int {
	return f.Score
}

// GetColor returns the color of the food.
// It returns the color of the food in the game.
func (f *Food) GetColor() colorUtil.Color {
	return f.Color
}
