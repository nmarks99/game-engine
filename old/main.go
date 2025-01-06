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

type Paddle struct {
    Rectangle rl.Rectangle
    Color rl.Color
}

func main() {
	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "Physics Simulation")
	rl.SetTargetFPS(TARGET_FPS)

	// const g float32 = 9.81 * 250.0 // scale to pixels/s/s
	engine := NewEngine()

	var particle_radius float32 = 20.0
	var particle_mass float32 = 1.0
	var particleX0 float32 = float32(SCREEN_WIDTH) / 8.0
	var particleY0 float32 = float32(SCREEN_HEIGHT) / 2.0
    p1 := NewParticle(rl.NewVector2(particleX0, particleY0), particle_radius, particle_mass, rl.Red)
    p1.Velocity = getRandomVector(300)
    engine.AddParticle(&p1)

    b1 := NewBlock(rl.NewVector2(300, 300), rl.NewVector2(50, 30), 1.0, rl.Blue)
    engine.AddBlock(&b1)

    // rl.HideCursor()
    // paddle := Paddle {rl.NewRectangle(0, 0, 20, 60), rl.Black}

	for !rl.WindowShouldClose() {
		mousePos := rl.GetMousePosition()
		engine.Update(mousePos)

        // paddle.Rectangle.X = mousePos.X
        // paddle.Rectangle.Y = mousePos.Y
        // if rl.CheckCollisionCircleRec(engine.particles[0].Position, engine.particles[0].Radius, paddle.Rectangle) {
            // engine.particles[0].Velocity.X *= -1
        // }

		// --------------------- DRAWING ---------------------
		rl.BeginDrawing()
		rl.ClearBackground(rl.LightGray)

		engine.Draw()

		rl.EndDrawing()
		// ---------------------------------------------------
	}

	rl.CloseWindow()
}
