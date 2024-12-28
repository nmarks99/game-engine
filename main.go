package main

import (
	"fmt"
    "time"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const SCREEN_WIDTH int32 = 700
const SCREEN_HEIGHT int32 = 600
const BORDER_WIDTH float32 = float32(SCREEN_WIDTH)*0.95
const BORDER_HEIGHT float32 = float32(SCREEN_HEIGHT)*0.95
const BORDER_X float32 = (float32(SCREEN_WIDTH) - BORDER_WIDTH) / 2.0
const BORDER_Y float32 = (float32(SCREEN_HEIGHT) - BORDER_HEIGHT) / 2.0
const MOVE_INCREMENT float32 = 2

func main() {
	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "Physics Simulation")
	border := rl.NewRectangle(BORDER_X, BORDER_Y, BORDER_WIDTH, BORDER_HEIGHT)

	const particleX0 float32 = float32(SCREEN_WIDTH) / 2.0
	const particleY0 float32 = float32(SCREEN_HEIGHT) / 2.0
	particle := NewParticle(rl.NewVector2(particleX0, particleY0), 20, 1, rl.Blue)

	quitButton := NewButton("Quit",
		func() {
			rl.EndDrawing()
			rl.CloseWindow()
		},
		rl.NewVector2(10, 10),  // position
		rl.NewVector2(60, 20), // size
		rl.Black,
	)

	rl.SetTargetFPS(60)
    
    t0 := time.Now()

	for !rl.WindowShouldClose() {

		// if rl.CheckCollisionCircleLine(particle.Position,
            // particle.Radius,
            // rl.NewVector2(border.X, border.Height+border.Y),
            // rl.NewVector2(border.X + border.Width, border.Height + border.Y),
        // ) {
			// particle.SetX(particleX0)
			// particle.SetY(particleY0)
		// }

		// if not colliding with walls
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

        timestamp := time.Since(t0)

		mousePos := rl.GetMousePosition()
        mousePosText := fmt.Sprintf("Mouse: %.2f, %.2f", mousePos.X, mousePos.Y)
        particlePosText := fmt.Sprintf("Particle: %.2f, %.2f", particle.Position.X, particle.Position.Y)
		timestampText := fmt.Sprintf("Time: %.2f s", float32(timestamp.Milliseconds())/1000.0)

		// --------------------- DRAWING ---------------------
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
        rl.DrawRectangleLinesEx(border, 5.0, rl.Red)
        rl.DrawText(mousePosText, 30, 30, 20, rl.Black)
        rl.DrawText(particlePosText, 30, 50, 20, rl.Brown)
		rl.DrawText(timestampText, 300, 35, 30, rl.Maroon)
		quitButton.Draw(mousePos)
		particle.Draw()
		rl.EndDrawing()
		// ---------------------------------------------------
	}

	rl.CloseWindow()
}
