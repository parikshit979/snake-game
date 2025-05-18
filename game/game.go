package game

import (
	"image"
	colorUtil "image/color"
	"image/png"
	"os"
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
	lastUpdate time.Time
	// The game over message.
	gameOverMessage string
	// The game over color.
	gameOverColor colorUtil.Color
	// The game over font size.
	gameOverFontSize int
	// The game over font source.
	mPlusFaceSource *text.GoTextFaceSource
	// captureScreenshot flag to capture the screenshot.
	captureScreenshot bool
	// gameStarted flag to check if the game has started.
	gameStarted bool
}

func NewGame(snake *Snake, food *Food, gameSpeed int,
	mplusFaceSource *text.GoTextFaceSource) *Game {
	return &Game{
		snake:             snake,
		food:              food,
		gameSpeed:         time.Second / time.Duration(gameSpeed),
		mPlusFaceSource:   mplusFaceSource,
		gameOverMessage:   "Game Over!",
		gameOverColor:     colorUtil.White,
		gameOverFontSize:  50,
		captureScreenshot: false,
		gameStarted:       false,
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

// CaptureScreenshot captures the current screen and saves it as a PNG file.
func (g *Game) CaptureScreenshot(screen *ebiten.Image) {
	// Create an empty image with the same dimensions as the screen.
	img := image.NewRGBA(image.Rect(0, 0, ScreenWidth, ScreenHeight))

	// Copy pixel data from the screen to the image.
	for y := 0; y < ScreenHeight; y++ {
		for x := 0; x < ScreenWidth; x++ {
			img.Set(x, y, screen.At(x, y))
		}
	}

	// Create a file to save the screenshot.
	filename := "screenshot_" + time.Now().Format("20060102_150405") + ".png"
	file, err := os.Create(filename)
	if err != nil {
		// Handle file creation error.
		ebitenutil.DebugPrint(screen, "Failed to save screenshot!")
		return
	}
	defer file.Close()

	// Encode the image as PNG and save it to the file.
	if err := png.Encode(file, img); err != nil {
		// Handle encoding error.
		ebitenutil.DebugPrint(screen, "Failed to encode screenshot!")
		return
	}

	// Notify the user that the screenshot was saved.
	ebitenutil.DebugPrint(screen, "Screenshot saved: "+filename)
}

// Update is called every frame. It should return the next game state.
func (g *Game) Update() error {

	// Check for the Enter key to start the game.
	if ebiten.IsKeyPressed(ebiten.KeyEnter) && !g.IsGameStarted() {
		// Start the game if the Enter key is pressed.
		g.gameStarted = true
	}

	// Check if the game has started.
	if !g.IsGameStarted() {
		return nil
	}

	// Check for the P key to capture a screenshot.
	if ebiten.IsKeyPressed(ebiten.KeyP) {
		// Set a flag to capture the screenshot in the Draw method.
		g.captureScreenshot = true
	}

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
	if time.Since(g.lastUpdate) < g.gameSpeed {
		return nil
	}
	g.lastUpdate = time.Now()

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

func (g *Game) AddScreenMessage(screen *ebiten.Image, message string) {
	// Draw the start message if the game has not started.
	face := &text.GoTextFace{
		Source: g.mPlusFaceSource,
		Size:   float64(g.gameOverFontSize),
	}
	w, h := text.Measure(message, face, face.Size)
	drawOptions := &text.DrawOptions{}
	screenWidth := (ScreenWidth - w) / 2
	screenHeight := (ScreenHeight - h) / 2
	drawOptions.GeoM.Translate(float64(screenWidth), float64(screenHeight))
	drawOptions.ColorScale.ScaleWithColor(g.gameOverColor)
	text.Draw(screen, message, face, drawOptions)

	// Draw the game over message if the game is over.
	if g.GetSnake().IsGameOver() {
		ebitenutil.DebugPrintAt(screen, "Press R to restart",
			int(screenWidth)+80, int(screenHeight)+80)
	}
}

// Draw is called every frame. It should draw the game state to the screen.
func (g *Game) Draw(screen *ebiten.Image) {
	if !g.IsGameStarted() {
		// Draw the start message if the game has not started.
		g.AddScreenMessage(screen, "Press Enter to start")
		return
	}
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
		// Draw the game over message.
		g.AddScreenMessage(screen, g.gameOverMessage)
	}

	// Capture the screenshot if the flag is set.
	if g.captureScreenshot {
		g.CaptureScreenshot(screen)
		g.captureScreenshot = false
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

// IsGameStarted returns true if the game has started.
func (g *Game) IsGameStarted() bool {
	return g.gameStarted
}
