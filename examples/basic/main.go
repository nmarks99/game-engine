package main

import (
    rl "github.com/gen2brain/raylib-go/raylib"
	. "raychip"
)

func main() {
    game := NewGame(600, 600, 60)

    box := NewBox(NewVector2(300,300), 50, 50, rl.Red)
    game.AddEntity(&box)

    const inc float64 = 10
    game.SetUpdateCallback(func (g *Game) {
        newPos := box.Position()
        if rl.IsKeyDown(rl.KeyUp) {
            newPos.Y -= inc
        }
        if rl.IsKeyDown(rl.KeyDown) {
            newPos.Y += inc
        }
        if rl.IsKeyDown(rl.KeyRight) {
            newPos.X += inc
        }
        if rl.IsKeyDown(rl.KeyLeft) {
            newPos.X -= inc
        }
        box.SetPosition(newPos.X, newPos.Y)
    })

    game.Run()
}
