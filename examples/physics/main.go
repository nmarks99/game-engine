package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	. "raychip"
)

func main() {
	const (
		SCREEN_WIDTH  int32 = 1000
		SCREEN_HEIGHT int32 = 800
		TARGET_FPS    int32 = 60
	)

	// Create a new game instance
	game := NewGame(SCREEN_WIDTH, SCREEN_HEIGHT, TARGET_FPS)
	game.SetDamping(0.9)
	game.SetGravity(NewVector2(0.0, 1000.0))

	// Create a permiter wall (invisible)
	game.AddPerimiterWall(1, rl.NewColor(0, 0, 0, 0))

	// Cursor texture, hide default cursor
	cursorTexture := rl.LoadTexture("./assets/cursors/Tiles/tile_0026.png")
	rl.HideCursor()

	// Texture for the circles
	ballTexture := rl.LoadTexture("./assets/planets/Terran.png")

	var mousePos Vector2
	var mouseVel Vector2
	var numEntities int
	const ballRadius float64 = 20.0
	const boxWidth float64 = 50.0

	// Custom game update function
	game.SetUpdateCallback(func(game *Game) {
		mousePos = Vector2FromRaylib(rl.GetMousePosition())
		mouseVel = Vector2FromRaylib(rl.Vector2Scale(rl.GetMouseDelta(), 50.0))
		numEntities = game.EntitiesCount()

		// lift click adds a ball with the velocity of the mouse
		if rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
			new_ball := NewPhysicalCircle(mousePos.X, mousePos.Y, ballRadius, 1.0, rl.DarkGreen)
			new_ball.SetElasticity(0.7)
			new_ball.SetFriction(0.7)
			new_ball.SetTexture(ballTexture)
			new_ball.SetVelocity(mouseVel.X, mouseVel.Y)
			game.AddEntity(&new_ball)
		}

		// right click adds a square with the velocity of the mouse
		if rl.IsMouseButtonReleased(rl.MouseButtonRight) {
			new_box := NewPhysicalBox(mousePos.X, mousePos.Y, boxWidth, boxWidth, 1.0, rl.NewColor(39, 81, 130, 255))
			new_box.SetElasticity(0.7)
			new_box.SetFriction(0.7)
			new_box.SetVelocity(mouseVel.X, mouseVel.Y)
			game.AddEntity(&new_box)
		}

		// scroll wheel click adds a non physical circle
		if rl.IsMouseButtonReleased(rl.MouseButtonMiddle) {
			new_box := NewBox(mousePos.X, mousePos.Y, boxWidth, boxWidth, rl.Red)
			game.AddEntity(&new_box)
		}
	})

	// Custom game draw function
	game.SetDrawCallback(func(game *Game) {
		rl.DrawTextureEx(cursorTexture, mousePos.ToRaylib(), 0.0, 2.0, rl.White)
		rl.DrawText(fmt.Sprintf("Entities: %d\n", numEntities-4), 10, 10, 20, rl.Black)
	})

	// Run the main game loop
	game.Run()

}
