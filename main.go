package main

import (
	"bytes"
	"log"

	colorUtil "image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"

	"github.com/snake-game/game"
)

func main() {
	// Set snake start position to the center of the screen.
	snakeStartX := game.ScreenWidth / game.GridSize / 2
	snakeStartY := game.ScreenHeight / game.GridSize / 2

	// Create a new snake instance. The snake starts at the center of the screen,
	// facing right, with a length of 1, speed of 10, and color white.
	snake := game.NewSnake(game.NewPoint(snakeStartX, snakeStartY),
		game.NewDirection(1, 0), 10, colorUtil.White, 0, 20)

	// Create a new food instance. The food is spawned at a random position.
	food := game.NewFood()
	food.SpawnFood()

	// Create a new font source for rendering text. The font is loaded from the
	// MPlus1pRegular_ttf font file. The font is used for rendering the game over
	// message.
	mplusFaceSource, err := text.NewGoTextFaceSource(
		bytes.NewReader(
			fonts.MPlus1pRegular_ttf,
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Create a new game instance. This is where the game logic and rendering code is defined.
	// The Game struct should implement the ebiten.Game interface.
	gameInstance := game.NewGame(snake, food, game.GameSpeed, mplusFaceSource)

	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(game.ScreenWidth, game.ScreenHeight)
	ebiten.SetWindowTitle("Snake Game")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(gameInstance); err != nil {
		log.Fatal(err)
	}
}
