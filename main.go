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
	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "Simple Game Engine - Test")
	rl.SetTargetFPS(TARGET_FPS)
    
    game := NewGame()

    p1 := NewParticle(NewVector2(50.0, 200.0), 20.0, 1.0, rl.Green)
    p1.Velocity.X = 200
    p2 := NewParticle(NewVector2(200.0, 190.0), 20.0, 1.0, rl.Red)
    game.AddParticle(&p1)
    game.AddParticle(&p2)

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
