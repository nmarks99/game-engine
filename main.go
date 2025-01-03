package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math/rand"
)

func randomFloat(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func getRandomVector(max float32) rl.Vector2 {
    return rl.NewVector2(randomFloat(-max, max), randomFloat(-max, max))
}

const SCREEN_WIDTH int32 = 600
const SCREEN_HEIGHT int32 = 600
const TARGET_FPS int32 = 60
const dt float32 = 1.0 / float32(TARGET_FPS) // simulation timestep

func main() {
	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "Physics Simulation")
	rl.SetTargetFPS(TARGET_FPS)

	const g float32 = 9.81 * 250.0 // scale to pixels/s/s
	engine := NewEngine()
	engine.SetGravity(rl.NewVector2(0.0, 0.0))

	var particle_radius float32 = 30.0
	var particle_mass float32 = 1.0
	var particleX0 float32 = float32(SCREEN_WIDTH) / 8.0
	var particleY0 float32 = float32(SCREEN_HEIGHT) / 2.0
    var velInit float32 = 700
    p1 := NewParticle(rl.NewVector2(particleX0, particleY0), particle_radius, particle_mass, rl.Red)
    p2 := NewParticle(rl.NewVector2(particleX0+particle_radius*3, particleY0), particle_radius, particle_mass, rl.Orange)
    p3 := NewParticle(rl.NewVector2(particleX0+particle_radius*6, particleY0), particle_radius, particle_mass, rl.DarkBrown)
    p4 := NewParticle(rl.NewVector2(particleX0+particle_radius*9, particleY0), particle_radius, particle_mass, rl.DarkGreen)
    p5 := NewParticle(rl.NewVector2(particleX0+particle_radius*12, particleY0), particle_radius, particle_mass, rl.Blue)
    p6 := NewParticle(rl.NewVector2(particleX0, particleY0+100), particle_radius, particle_mass, rl.DarkBlue)
    p7 := NewParticle(rl.NewVector2(particleX0+particle_radius*3, particleY0+100), particle_radius, particle_mass, rl.Purple)
    p8 := NewParticle(rl.NewVector2(particleX0+particle_radius*6, particleY0+100), particle_radius, particle_mass, rl.DarkPurple)
    p9 := NewParticle(rl.NewVector2(particleX0+particle_radius*9, particleY0+100), particle_radius, particle_mass, rl.SkyBlue)
    p10 := NewParticle(rl.NewVector2(particleX0+particle_radius*12, particleY0+100), particle_radius, particle_mass, rl.Maroon)

    p1.Velocity = getRandomVector(velInit)
    p2.Velocity = getRandomVector(velInit)
    p3.Velocity = getRandomVector(velInit)
    p4.Velocity = getRandomVector(velInit)
    p5.Velocity = getRandomVector(velInit)
    p6.Velocity = getRandomVector(velInit)
    p7.Velocity = getRandomVector(velInit)
    p8.Velocity = getRandomVector(velInit)
    p9.Velocity = getRandomVector(velInit)
    p10.Velocity = getRandomVector(velInit)

    engine.AddParticle(&p1)
    engine.AddParticle(&p2)
    engine.AddParticle(&p3)
    engine.AddParticle(&p4)
    engine.AddParticle(&p5)
    engine.AddParticle(&p6)
    engine.AddParticle(&p7)
    engine.AddParticle(&p8)
    engine.AddParticle(&p9)
    engine.AddParticle(&p10)
    
	for !rl.WindowShouldClose() {
		mousePos := rl.GetMousePosition()
		engine.Update(mousePos)

		// --------------------- DRAWING ---------------------
		rl.BeginDrawing()
		rl.ClearBackground(rl.LightGray)
		engine.Draw()
		rl.EndDrawing()
		// ---------------------------------------------------
	}

	rl.CloseWindow()
}
