package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	SCREEN_WIDTH  int32 = 600
	SCREEN_HEIGHT int32 = 600
	TARGET_FPS    int32 = 60
)

func main() {
	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "Simple Game Engine")
	rl.SetTargetFPS(TARGET_FPS)

	game := NewGame()

	// Create some objects
	p1 := NewParticle(NewVector2(50.0, 200.0), 20.0, 1.0, rl.Green)
	p1.Velocity.X = 200
	p2 := NewParticle(NewVector2(200.0, 190.0), 20.0, 1.0, rl.Red)
	b1 := NewBox(NewVector2(200, 400), 100.0, 20.0, 1.0, rl.Purple)
	game.AddEntity(&p1)
	game.AddEntity(&p2)
	game.AddEntity(&b1)

    // Create a permiter wall
    wallWidth := 20.0
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


	for !rl.WindowShouldClose() {

		game.Update()

		// ---------- Drawing ----------
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		game.Draw()
		rl.EndDrawing()
		// -----------------------------
	}

	rl.CloseWindow()

}
