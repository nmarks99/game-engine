package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
    . "engine"
    "math"
)

func main() {
	const (
		SCREEN_WIDTH  int32 = 1000
		SCREEN_HEIGHT int32 = 800
		TARGET_FPS    int32 = 60
	)

	// Create a new game instance
	game := NewGame(SCREEN_WIDTH, SCREEN_HEIGHT, TARGET_FPS)
    // game.SetGravity(NewVector2(0.0, 800.0))
    game.SetDamping(0.9)

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
    ballRadius := 20.0
	ball := NewParticle(NewVector2(250.0, 50.0), ballRadius, 1.0, rl.DarkGreen)
    game.AddEntity(&ball)    

    // add custom draw function for ball to draw line for angle
    ball_line_draw_cbk := func(p *Particle) {
        DefaultParticleDrawFunc(p)
        x0 := p.Position().X - p.Radius() * math.Cos(p.Angle())
        y0 := p.Position().Y - p.Radius() * math.Sin(p.Angle())
        x1 := p.Position().X + p.Radius() * math.Cos(p.Angle())
        y1 := p.Position().Y + p.Radius() * math.Sin(p.Angle())
        rl.DrawLineEx(NewVector2(x0, y0).ToRaylib(), NewVector2(x1, y1).ToRaylib(), 3, rl.Yellow) 
    }
    ball.SetDrawCallback(ball_line_draw_cbk)

    // Custom game callback
    game.SetUpdateCallback(func(game *Game) {
        mousePos := rl.GetMousePosition()
        mouseVel := rl.Vector2Scale(rl.GetMouseDelta(), 50.0)

        if rl.IsMouseButtonReleased(rl.MouseButtonRight) {
	        new_ball := NewParticle(NewVector2(float64(mousePos.X), float64(mousePos.Y)), ballRadius, 1.0, rl.DarkGreen)
            new_ball.SetDrawCallback(ball_line_draw_cbk)
            new_ball.SetVelocity(float64(mouseVel.X), float64(mouseVel.Y))
            game.AddEntity(&new_ball)
        }

        if rl.IsMouseButtonReleased(rl.MouseButtonMiddle) {
            const boxWidth float64 = 50.0
            new_box := NewBox(NewVector2(float64(mousePos.X), float64(mousePos.Y)), boxWidth, boxWidth, 1.0, rl.Red)
            new_box.SetVelocity(float64(mouseVel.X), float64(mouseVel.Y))
            game.AddEntity(&new_box)
        }
    })
    
	// Run the main game loop
	game.Run()

}
