package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
    . "engine"
)

func main() {
	const (
		SCREEN_WIDTH  int32 = 1000
		SCREEN_HEIGHT int32 = 600
		TARGET_FPS    int32 = 60
	)

	// Create a new game instance
	game := NewGame(SCREEN_WIDTH, SCREEN_HEIGHT, TARGET_FPS)

	// Create some objects
	p1 := NewParticle(NewVector2(50.0, 200.0), 20.0, 1.0, rl.Green)
	p2 := NewParticle(NewVector2(200.0, 185.0), 20.0, 1.0, rl.Red)
	p3 := NewParticle(NewVector2(300.0, 185.0), 20.0, 1.0, rl.Black)
	b1 := NewBox(NewVector2(200.0, 300.0), 100.0, 20.0, 0.2, rl.Purple)
	b2 := NewBox(NewVector2(220.0, 350.0), 50.0, 50.0, 0.2, rl.DarkBlue)
	b3 := NewBox(NewVector2(400.0, 330.0), 50.0, 70.0, 0.2, rl.Orange)
	b4 := NewBox(NewVector2(600.0, 400.0), 50.0, 70.0, 0.2, rl.DarkGreen)
	p1.Velocity.X = 400
	p2.Velocity.X = -400
	game.AddEntity(&p1)
	game.AddEntity(&p2)
	game.AddEntity(&p3)
	game.AddEntity(&b1)
	game.AddEntity(&b2)
	game.AddEntity(&b3)
	game.AddEntity(&b4)

	// Create a permiter wall
	wallWidth := 5.0
	v_tl := NewVector2(0, 0)
	v_tr := NewVector2(float64(SCREEN_WIDTH), 0)
	v_br := NewVector2(float64(SCREEN_WIDTH), float64(SCREEN_HEIGHT))
	v_bl := NewVector2(0, float64(SCREEN_HEIGHT))
	wallTop := NewWall(v_tl, v_tr, wallWidth, rl.Black)
	wallRight := NewWall(v_tr, v_br, wallWidth, rl.Black)
	wallBot := NewWall(v_bl, v_br, wallWidth, rl.Black)
	wallLeft := NewWall(v_tl, v_bl, wallWidth, rl.Black)
	game.AddEntity(&wallTop)
	game.AddEntity(&wallRight)
	game.AddEntity(&wallBot)
	game.AddEntity(&wallLeft)

	// Add custom update and draw functions
	var mousePos rl.Vector2
	var mousePosText string
	game.SetUpdateCallback(func(game *Game) {
		mousePos = rl.GetMousePosition()
		mousePosText = fmt.Sprintf("%.2f, %.2f", mousePos.X, mousePos.Y)
	})

	game.SetDrawCallback(func(game *Game) {
		rl.DrawText(mousePosText, 10, 10, 18, rl.DarkBlue)
	})

	// Run the main game loop
	game.Run()

}
