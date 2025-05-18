package game

import (
	colorUtil "image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// Game represents the game state and implements ebiten.Game interface.
// It contains the game logic and rendering code.
type Game struct {
	// The snake instance in the game.
	snake *Snake
	// The food instance in the game.
	food *Food
	// The game speed in ticks per second.
	gameSpeed time.Duration
	// The last update time.
	lasUpdate time.Time
	// The game over message.
	gameOverMessage string
	// The game over color.
	gameOverColor colorUtil.Color
	// The game over font size.
	gameOverFontSize int
	// The game over font source.
	mPlusFaceSource *text.GoTextFaceSource
}

func NewGame(snake *Snake, food *Food, gameSpeed int,
	mplusFaceSource *text.GoTextFaceSource) *Game {
	return &Game{
		snake:            snake,
		food:             food,
		gameSpeed:        time.Second / time.Duration(gameSpeed),
		mPlusFaceSource:  mplusFaceSource,
		gameOverMessage:  "Game Over!",
		gameOverColor:    colorUtil.White,
		gameOverFontSize: 50,
	}
}

func (g *Game) ReadKeyboard() Direction {
	direction := g.GetSnake().GetDirection()

	// Check for keyboard input to change the snake's direction.
	if !g.GetSnake().IsGameOver() {
		// If the game is over, do not change the direction.
		// Check for keyboard input to change the snake's direction.
		if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
			// Move the snake up.
			direction = Up()
		} else if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
			// Move the snake down.
			direction = Down()
		} else if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
			// Move the snake left.
			direction = Left()
		} else if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
			// Move the snake right.
			direction = Right()
		}
	}

	// Check for the R key to restart the game.
	if ebiten.IsKeyPressed(ebiten.KeyR) && g.GetSnake().IsGameOver() {
		// Restart the game if the game is over and the R key is pressed.
		g.GetSnake().Restart()
		g.GetFood().SpawnFood()
	}
	return direction
}

// Update is called every frame. It should return the next game state.
func (g *Game) Update() error {
	// Read keyboard input.
	direction := g.ReadKeyboard()
	if g.GetSnake().IsOppositeDirection(direction) {
		// If the snake is moving in the opposite direction, do not change the direction.
		return nil
	}
	if g.GetSnake().CheckCollision(ScreenWidth, ScreenHeight, direction) {
		return nil
	}
	g.GetSnake().SetDirection(direction)

	// Check game speed and update the game state.
	if time.Since(g.lasUpdate) < g.gameSpeed {
		return nil
	}
	g.lasUpdate = time.Now()

	if !g.GetSnake().IsGameOver() {
		// Move the snake in the current direction.
		g.GetSnake().Move()

		// Spawn food if it is eaten.
		if g.GetFood().IsEaten(*g.GetSnake()) {
			// Increase the snake's length and score.
			point := g.GetFood().GetPosition()
			g.GetSnake().IncreaseLength(point)
			g.GetSnake().IncreaseScore(g.GetFood().GetScore())

			// Spawn new food at a random position.
			g.GetFood().SpawnFood()
		}
	}
	return nil
}

// Draw is called every frame. It should draw the game state to the screen.
func (g *Game) Draw(screen *ebiten.Image) {
	// Draw the game state here
	for _, point := range g.GetSnake().GetBody() {
		vector.DrawFilledRect(screen, float32(point.X*GridSize),
			float32(point.Y*GridSize), GridSize, GridSize,
			g.GetSnake().GetColor(), true)
	}

	vector.DrawFilledRect(screen, float32(g.GetFood().GetPosition().X*GridSize),
		float32(g.GetFood().GetPosition().Y*GridSize), GridSize, GridSize,
		g.GetFood().GetColor(), true)

	// Draw the game over message if the game is over.
	if g.GetSnake().IsGameOver() {
		face := &text.GoTextFace{
			Source: g.mPlusFaceSource,
			Size:   float64(g.gameOverFontSize),
		}

		w, h := text.Measure(g.gameOverMessage, face, face.Size)
		drawOptions := &text.DrawOptions{}
		screenWidth := (ScreenWidth - w) / 2
		screenHeight := (ScreenHeight - h) / 2
		drawOptions.GeoM.Translate(float64(screenWidth), float64(screenHeight))
		drawOptions.ColorScale.ScaleWithColor(g.gameOverColor)
		text.Draw(screen, g.gameOverMessage, face, drawOptions)

		ebitenutil.DebugPrintAt(screen, "Press R to restart",
			int(screenWidth)+80, int(screenHeight)+80)
	}
}

// Layout is called when the window is resized. It should return the new screen size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// Return the screen size here
	return ScreenWidth, ScreenHeight
}

// GetSnake returns the snake instance.
func (g *Game) GetSnake() *Snake {
	return g.snake
}

// GetFood returns the food instance.
func (g *Game) GetFood() *Food {
	return g.food
}
