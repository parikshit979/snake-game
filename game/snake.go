package game

import (
	colorUtil "image/color"
)

// Snake represents the snake in the game.
// It contains the snake's head position, direction, length, body, speed,
// color, score, and game over state.
type Snake struct {
	// The current head position of the snake.
	Head *Point
	// The current direction of the snake.
	Direction Direction
	// The length of the snake.
	Length int
	// The body of the snake, represented as a slice of Points.
	Body []*Point
	// The speed of the snake.
	Speed int
	// The color of the snake.
	Color colorUtil.Color
	// The score of the snake.
	Score int
	// The maximum score of the snake.
	MaxScore int
	// The game over flag.
	GameOver bool
}

// NewSnake creates a new snake with the given head position, direction, length,
// speed, color, score, and maximum score.
func NewSnake(head Point, direction Direction, speed int,
	color colorUtil.Color, score int, maxScore int) *Snake {
	snake := &Snake{
		Head:      &head,
		Direction: direction,
		Speed:     speed,
		Color:     color,
		Score:     score,
		MaxScore:  maxScore,
		GameOver:  false,
	}
	snake.Body = append(snake.Body, &head)
	snake.Length = 1
	return snake
}

// Move moves the snake in the current direction.
// It updates the head position and the body of the snake.
func (s *Snake) Move() {
	// Move the head in the current direction.
	newHead := Point{
		X: s.GetHead().X + s.GetDirection().point.X,
		Y: s.GetHead().Y + s.GetDirection().point.Y,
	}
	// Update the head position.
	s.Head = &newHead

	// Update the body of the snake.
	s.Body = append([]*Point{s.Head}, s.Body[:s.Length-1]...)
}

// CheckCollision checks if the snake has collided with itself or the walls.
// It returns true if the snake has collided, false otherwise.
func (s *Snake) CheckCollision(screenWidth, screenHeight int, direction Direction) bool {
	headX := s.Head.X + direction.point.X
	headY := s.Head.Y + direction.point.Y
	// Check if the snake has collided with itself.
	for _, point := range s.Body {
		if headX == point.X && headY == point.Y {
			s.GameOver = true
			// Collision with self detected.
			return true
		}
	}

	// Check if the snake has collided with the walls.
	if headX < 0 || headX >= screenWidth/GridSize || headY < 0 || headY >= screenHeight/GridSize {
		s.GameOver = true
		// Collision with wall detected.
		return true
	}

	return false
}

// CheckWin checks if the snake has won the game.
// It returns true if the snake has won, false otherwise.
func (s *Snake) CheckWin() bool {
	// Check if the snake has reached the maximum score
	if s.Score >= s.MaxScore {
		s.GameOver = true
		return true
	}
	return false
}

// IncreaseScore increases the score of the snake.
// It updates the score of the snake based on the food eaten.
func (s *Snake) IncreaseScore(score int) {
	s.Score += score
	if s.Score > s.MaxScore {
		s.MaxScore = s.Score
	}
}

// GetBody returns the body of the snake.
// It returns the body of the snake as a slice of Points.
func (s *Snake) GetBody() []*Point {
	return s.Body
}

// GetHead returns the head of the snake.
// It returns the head of the snake as a Point.
func (s *Snake) GetHead() *Point {
	return s.Head
}

// UpdateHead updates the head of the snake.
// It updates the head of the snake based on the input.
func (s *Snake) UpdateHead(newHead *Point) {
	s.Head = newHead
}

// GetColor returns the color of the snake.
// It returns the color of the snake as a colorUtil.Color.
func (s *Snake) GetColor() colorUtil.Color {
	return s.Color
}

// SetDirection sets the direction of the snake.
// It sets the direction of the snake based on the input.
func (s *Snake) SetDirection(direction Direction) {
	s.Direction = direction
}

// GetDirection returns the direction of the snake.
// It returns the direction of the snake as a Direction.
func (s *Snake) GetDirection() Direction {
	return s.Direction
}

// IncreaseLength increases the length of the snake.
// It increases the length of the snake by 1.
func (s *Snake) IncreaseLength(point *Point) {
	s.Length++
	// Add a new point to the body of the snake
	s.Body = append([]*Point{point}, s.Body...)
}

// IsGameOver checks if the game is over.
// It returns true if the game is over, false otherwise.
func (s *Snake) IsGameOver() bool {
	return s.GameOver
}

// Restart restarts the game.
// It resets the snake's position, direction, length, and score.
func (s *Snake) Restart() {
	// Set snake start position to the center of the screen.
	s.Head = &Point{X: ScreenWidth / GridSize / 2,
		Y: ScreenHeight / GridSize / 2}
	s.Length = 1
	s.Body = []*Point{s.Head}
	s.Score = 0
	s.MaxScore = 20
	s.GameOver = false
}

// GetScore returns the score of the snake.
// It returns the score of the snake as an int.
func (s *Snake) GetScore() int {
	return s.Score
}

// GetMaxScore returns the maximum score of the snake.
// It returns the maximum score of the snake as an int.
func (s *Snake) GetMaxScore() int {
	return s.MaxScore
}

func (s *Snake) IsOppositeDirection(direction Direction) bool {
	return (s.Direction.point.X == -direction.point.X &&
		s.Direction.point.Y == -direction.point.Y)
}

func (s *Snake) IsSameDirection(direction Direction) bool {
	return (s.Direction.point.X == direction.point.X &&
		s.Direction.point.Y == direction.point.Y)
}
