package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const SCREEN_WIDTH int32 = 600
const SCREEN_HEIGHT int32 = 400
const MOVE_INCREMENT float32 = 5


func main() {
	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "raylib dev")
    
    particle := NewParticle(rl.NewVector2(100, 100), 10, rl.Blue)

    quitButton := NewButton(
        "Quit",
        func() {
            rl.EndDrawing()
            rl.CloseWindow()
        },
        rl.NewVector2(10,10),
        rl.NewVector2(100, 20),
        rl.Black)


	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

        if rl.IsKeyDown(rl.KeyW) {
            particle.IncrementY(-MOVE_INCREMENT)
        }
        if rl.IsKeyDown(rl.KeyS) {
            particle.IncrementY(MOVE_INCREMENT)
        }
        if rl.IsKeyDown(rl.KeyD) {
            particle.IncrementX(MOVE_INCREMENT)
        }
        if rl.IsKeyDown(rl.KeyA) {
            particle.IncrementX(-MOVE_INCREMENT)
        }

        mousePos := rl.GetMousePosition()

        // --------------------- DRAWING ---------------------
        rl.BeginDrawing()
        rl.ClearBackground(rl.RayWhite)
        quitButton.Draw(mousePos)
        particle.Draw()
        rl.EndDrawing()
        // ---------------------------------------------------
	}

    rl.CloseWindow()
}
