package main

import (
	"fmt"
    "math/rand"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func randomFloat(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

const TARGET_FPS int32 = 60
const dt float32 = 1.0 / float32(TARGET_FPS) // simulation timestep

func main() {
	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "Physics Simulation")
	rl.SetTargetFPS(TARGET_FPS)

    const g float32 = 9.81 * 250.0 // scale to pixels/s/s
    engine := NewEngine()
    engine.SetGravity(rl.NewVector2(0.0, 0.0))

    var particle_radius float32 = 20.0
    var particle_mass float32 = 1.0
    var particleX0 float32 = float32(SCREEN_WIDTH) / 2.0
    var particleY0 float32 = 100
	p1 := NewParticle(rl.NewVector2(particleX0, particleY0), particle_radius, particle_mass, rl.Blue)
    p2 := NewParticle(rl.NewVector2(particleX0+particle_radius*2, particleY0), particle_radius, particle_mass, rl.Orange)
    p3 := NewParticle(rl.NewVector2(particleX0+particle_radius*4, particleY0), particle_radius, particle_mass, rl.Purple)
    p4 := NewParticle(rl.NewVector2(particleX0+particle_radius*4, particleY0-particle_radius*2), particle_radius, particle_mass, rl.Yellow)
    p5 := NewParticle(rl.NewVector2(particleX0+particle_radius*4, particleY0-particle_radius*3), particle_radius, particle_mass, rl.Red)
    
    p1.Velocity = rl.NewVector2(randomFloat(-MAX_VELOCITY, MAX_VELOCITY), randomFloat(-MAX_VELOCITY, MAX_VELOCITY))
    p2.Velocity = rl.NewVector2(randomFloat(-MAX_VELOCITY, MAX_VELOCITY), randomFloat(-MAX_VELOCITY, MAX_VELOCITY))
    p3.Velocity = rl.NewVector2(randomFloat(-MAX_VELOCITY, MAX_VELOCITY), randomFloat(-MAX_VELOCITY, MAX_VELOCITY))
    p4.Velocity = rl.NewVector2(randomFloat(-MAX_VELOCITY, MAX_VELOCITY), randomFloat(-MAX_VELOCITY, MAX_VELOCITY))
    p5.Velocity = rl.NewVector2(randomFloat(-MAX_VELOCITY, MAX_VELOCITY), randomFloat(-MAX_VELOCITY, MAX_VELOCITY))

    engine.AddParticle(&p1)
    engine.AddParticle(&p2)
    engine.AddParticle(&p3)
    engine.AddParticle(&p4)
    engine.AddParticle(&p5)

    // t0 := time.Now()

	for !rl.WindowShouldClose() {

		mousePos := rl.GetMousePosition()
        // timestamp := time.Since(t0)

        engine.Update(mousePos)

        // text for displaying various info
        mousePosText := fmt.Sprintf("Mouse: %.2f, %.2f", mousePos.X, mousePos.Y)
        // fmt.Printf("Time: %.2f\n", float32(timestamp.Milliseconds())/1000.0)

		// --------------------- DRAWING ---------------------
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
        rl.DrawText(mousePosText, 30, 30, 20, rl.Black)

        engine.Draw()

		rl.EndDrawing()
		// ---------------------------------------------------
	}

	rl.CloseWindow()
}
