package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
    . "engine"
    "math"
)

func main() {
	const (
		SCREEN_WIDTH  int32 = 500
		SCREEN_HEIGHT int32 = 500
		TARGET_FPS    int32 = 60
	)

	// Create a new game instance
	game := NewGame(SCREEN_WIDTH, SCREEN_HEIGHT, TARGET_FPS)
    game.SetGravity(NewVector2(0.0, 400.0))
    // game.SetDamping(1.0)

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

    // create ball
	ball := NewParticle(NewVector2(250.0, 50.0), 20.0, 1.0, rl.DarkBlue)
    game.AddEntity(&ball)    
    ball.SetElasticity(1.0)
    ball.SetFriction(1.0)

    // add custom draw function for ball to draw line for angle
    ball.SetDrawCallback(func(p *Particle) {
        DefaultParticleDrawFunc(&ball)
        x := ball.Position().X + ball.Radius() * math.Cos(ball.Angle())
        y := ball.Position().Y + ball.Radius() * math.Sin(ball.Angle())
        rl.DrawLineEx(ball.Position().ToRaylib(), NewVector2(x, y).ToRaylib(), 3, rl.White) 
    })

    // Create paddle
    paddle := NewBox(NewVector2(100,100), 100, 10, math.Inf(1), rl.Red) 
    game.AddEntity(&paddle)
    paddle.SetKinematic()
    paddle.SetElasticity(1.0)
    paddle.SetFriction(1.0)

    game.SetUpdateCallback(func(game *Game) {
        mousePos := rl.GetMousePosition()
        mouseVel := rl.GetMouseDelta()
        paddle.SetPosition(float64(mousePos.X), float64(mousePos.Y))
        paddle.SetVelocity(float64(mouseVel.X)*50.0, float64(mouseVel.Y))
    })
    
	// Run the main game loop
	game.Run()

}
