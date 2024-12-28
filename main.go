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
const TARGET_FPS int32 = 60
const dt float32 = 1.0 / float32(TARGET_FPS) // simulation timestep


func main() {
	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "Physics Simulation")
	border := rl.NewRectangle(BORDER_X, BORDER_Y, BORDER_WIDTH, BORDER_HEIGHT)

	const particleX0 float32 = float32(SCREEN_WIDTH) / 6.0
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

	rl.SetTargetFPS(TARGET_FPS)
    
    t0 := time.Now()
    var g float32 = 9.81 * 100.0 // scale to 100 pixels/meter
    accel := rl.NewVector2(0.0, g)

    floorY := border.Height + border.Y
    
	for !rl.WindowShouldClose() {

		mousePos := rl.GetMousePosition()
        timestamp := time.Since(t0)

        // update position and velocity
        newVel := rl.Vector2Add(particle.Velocity, rl.Vector2Scale(accel, dt))
        newPos := rl.Vector2Add(particle.Position, rl.Vector2Scale(newVel, dt))
        if newPos.Y >= (floorY - particle.Radius) {
            newPos.Y = floorY - particle.Radius
            newVel.Y = 0.0
            newVel.X = 0.0
        }
        particle.Velocity = newVel
        particle.Position = newPos

        if rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
            particle.Position = mousePos
            particle.Velocity.X = 0.0
            particle.Velocity.Y = 0.0
        }

        // text for displaying various info
        mousePosText := fmt.Sprintf("Mouse: %.2f, %.2f", mousePos.X, mousePos.Y)
        particlePosText := fmt.Sprintf("Position: %.2f, %.2f", particle.Position.X, particle.Position.Y)
        particleVelText := fmt.Sprintf("Velocity: %.2f, %.2f", particle.Velocity.X, particle.Velocity.Y)
        fmt.Printf("Time: %.2f\n", float32(timestamp.Milliseconds())/1000.0)

		// --------------------- DRAWING ---------------------
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
        rl.DrawRectangleLinesEx(border, 5.0, rl.Red)
        rl.DrawText(particlePosText, 30, 30, 20, rl.Black)
        rl.DrawText(particleVelText, 30, 50, 20, rl.Black)
        rl.DrawText(mousePosText, 30, 70, 20, rl.Black)
		quitButton.Draw(mousePos)
		particle.Draw()
		rl.EndDrawing()
		// ---------------------------------------------------
	}

	rl.CloseWindow()
}
