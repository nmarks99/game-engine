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
	rl.SetTargetFPS(TARGET_FPS)

	border := rl.NewRectangle(BORDER_X, BORDER_Y, BORDER_WIDTH, BORDER_HEIGHT)

    const g float32 = 9.81 * 250.0 // scale to pixels/s/s
    world := NewWorld()
    world.SetAcceleration(rl.NewVector2(0.0, g))

    floorY := border.Height + border.Y
    var particle_radius float32 = 20.0
    var particle_mass float32 = 1.0
    particleX0 := float32(SCREEN_WIDTH) / 2.0
    particleY0 := float32(floorY) + float32(particle_radius)
	p1 := NewParticle(rl.NewVector2(particleX0, particleY0), particle_radius, particle_mass, rl.Blue)
	p2 := NewParticle(rl.NewVector2(particleX0+particle_radius*2, particleY0), particle_radius, particle_mass, rl.Orange)
	p3 := NewParticle(rl.NewVector2(particleX0+particle_radius*4, particleY0), particle_radius, particle_mass, rl.Purple)

    world.AddRigidBody(&p1)
    world.AddRigidBody(&p2)
    world.AddRigidBody(&p3)

	quitButton := NewButton("Quit",
		func() {
			rl.EndDrawing()
			rl.CloseWindow()
		},
		rl.NewVector2(10, 10),  // position
		rl.NewVector2(60, 20), // size
		rl.Black,
	)
    
    t0 := time.Now()

	for !rl.WindowShouldClose() {

		mousePos := rl.GetMousePosition()
        timestamp := time.Since(t0)

        world.Step(mousePos)

        // text for displaying various info
        mousePosText := fmt.Sprintf("Mouse: %.2f, %.2f", mousePos.X, mousePos.Y)
        fmt.Printf("Time: %.2f\n", float32(timestamp.Milliseconds())/1000.0)

		// --------------------- DRAWING ---------------------
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
        rl.DrawRectangleLinesEx(border, 5.0, rl.Red)
        rl.DrawText(mousePosText, 30, 30, 20, rl.Black)
        quitButton.Draw(mousePos)

        world.Draw()

		rl.EndDrawing()
		// ---------------------------------------------------
	}

	rl.CloseWindow()
}
