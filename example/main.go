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

    ballTexture := rl.LoadTexture("./assets/planets/Terran.png")
    
    // add custom draw function for ball to draw line for angle
    ball_line_draw_cbk := func(p *Particle) {

        pos := p.Position()
       
        ballTextureW := float32(ballTexture.Width)
        ballTextureH := float32(ballTexture.Height)
        srcRect := rl.NewRectangle(0, 0, ballTextureW, ballTextureH)
        destRect := rl.NewRectangle(float32(pos.X), float32(pos.Y), ballTextureW, ballTextureH)
        origin := rl.NewVector2(ballTextureW/2, ballTextureH/2)

        angle := float32(p.Angle()*180.0/math.Pi)
        rl.DrawTexturePro(ballTexture, srcRect, destRect, origin, float32(angle), rl.White)

        // ball with line
        // DefaultParticleDrawFunc(p)
        // x0 := p.Position().X - p.Radius() * math.Cos(p.Angle())
        // y0 := p.Position().Y - p.Radius() * math.Sin(p.Angle())
        // x1 := p.Position().X + p.Radius() * math.Cos(p.Angle())
        // y1 := p.Position().Y + p.Radius() * math.Sin(p.Angle())
        // rl.DrawLineEx(NewVector2(x0, y0).ToRaylib(), NewVector2(x1, y1).ToRaylib(), 3, rl.Yellow)
    }

    // Cursor texture, hide default cursor
    cursor := rl.LoadTexture("./assets/cursors/Tiles/tile_0026.png")
    rl.HideCursor()

    var mousePos rl.Vector2

    game.SetUpdateCallback(func(game *Game) {
        mousePos = rl.GetMousePosition()
        mouseVel := rl.Vector2Scale(rl.GetMouseDelta(), 50.0)

        // lift click adds a ball
        if rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
            const ballRadius float64 = 20.0
	        new_ball := NewParticle(NewVector2(float64(mousePos.X), float64(mousePos.Y)), ballRadius, 1.0, rl.DarkGreen)
            new_ball.SetDrawCallback(ball_line_draw_cbk)
            new_ball.SetVelocity(float64(mouseVel.X), float64(mouseVel.Y))
            game.AddEntity(&new_ball)
        }

        // right click adds a block
        if rl.IsMouseButtonReleased(rl.MouseButtonRight) {
            const boxWidth float64 = 50.0
            new_box := NewBox(NewVector2(float64(mousePos.X), float64(mousePos.Y)), boxWidth, boxWidth, 1.0, rl.NewColor(39, 81, 130, 255))
            new_box.SetVelocity(float64(mouseVel.X), float64(mouseVel.Y))
            game.AddEntity(&new_box)
        }
    })

    game.SetDrawCallback(func(game *Game){
        rl.DrawTextureV(cursor, mousePos, rl.White)
    })
    
	// Run the main game loop
	game.Run()

}
