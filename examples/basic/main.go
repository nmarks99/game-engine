package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	. "raychip"
)

func main() {
	game := NewGame(600, 600, 60)

    ball := NewCircle(300, 300, 30, rl.Red)
    game.AddEntity(&ball)

	const inc float64 = 10
    var i int = 0
    colors := []rl.Color{rl.Black, rl.Red}

    ball.OnClick(rl.MouseButtonLeft, MouseReleased, func(){
        i = i ^ 1
        ball.SetColor(colors[i])
    })

	game.SetUpdateCallback(func(g *Game) {
		// move with w,a,s,d
		newPos := ball.Position()
		if rl.IsKeyDown(rl.KeyW) {
			newPos.Y -= inc
		}
		if rl.IsKeyDown(rl.KeyS) {
			newPos.Y += inc
		}
		if rl.IsKeyDown(rl.KeyD) {
			newPos.X += inc
		}
		if rl.IsKeyDown(rl.KeyA) {
			newPos.X -= inc
		}
		ball.SetPosition(newPos.X, newPos.Y)
	})

	game.Run()
}
