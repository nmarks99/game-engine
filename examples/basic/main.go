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
	game.SetUpdateCallback(func(g *Game) {
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
