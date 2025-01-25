package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	. "raychip"
)

func main() {
	game := NewGame(600, 600, 60)

	box := NewBox(300, 300, 50, 50, rl.Red)
	game.AddEntity(&box)

	const inc float64 = 10
	var mousePos rl.Vector2
    var i int = 0
    colors := []rl.Color{rl.Black, rl.Red}

    game.OnClick(rl.MouseButtonLeft, MouseReleased, func(){
        mousePos = rl.GetMousePosition()
        if rl.CheckCollisionPointCircle(mousePos, box.Position().ToRaylib(), 30) {
            i = i ^ 1
            box.SetColor(colors[i])
        }
    })

	game.SetUpdateCallback(func(g *Game) {
		// move with w,a,s,d
		newPos := box.Position()
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
		box.SetPosition(newPos.X, newPos.Y)
	})

	game.Run()
}
