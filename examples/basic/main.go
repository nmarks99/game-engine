package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	. "raychip"
)

func main() {
	game := NewGame(600, 600, 60)

	ball := NewCircle(300, 300, 30, rl.Black)
	game.AddEntity(&ball)
	box := NewBox(500, 500, 50, 50, rl.Black)
	game.AddEntity(&box)

	bus := NewEventBus()
	mousePosPub := bus.CreatePublisher("mouse.position", Vector2{})
	bus.CreateSubscription("mouse.position", Vector2{}, func(pos Vector2) {
		if rl.CheckCollisionPointCircle(rl.GetMousePosition(), ball.Position().ToRaylib(), float32(ball.Radius())) {
			fmt.Println("Mouse is on ball!")
		}
	})
	bus.CreateSubscription("mouse.position", Vector2{}, func(pos Vector2) {
		if rl.CheckCollisionPointCircle(rl.GetMousePosition(), box.Position().ToRaylib(), float32(box.Width())) {
			fmt.Println("Mouse is on box!")
		}
	})

	var ballSelect int = 0
	ballColors := []rl.Color{rl.Black, rl.Red}
	ball.OnClick(&game, rl.MouseButtonLeft, MouseReleased, func() {
		ballSelect = ballSelect ^ 1
		ball.SetColor(ballColors[ballSelect])
	})

	var boxSelect int = 0
	boxColors := []rl.Color{rl.Black, rl.Red}
	box.OnClick(&game, rl.MouseButtonLeft, MouseReleased, func() {
		boxSelect = boxSelect ^ 1
		box.SetColor(boxColors[boxSelect])
	})

	const inc float64 = 10

	game.SetUpdateCallback(func(g *Game) {

		mousePosPub.Publish(Vector2FromRaylib(rl.GetMousePosition()))

		// move with w,a,s,d if selected
		newBallPos := ball.Position()
		newBoxPos := box.Position()

		if rl.IsKeyDown(rl.KeyW) {
			if ballSelect == 1 {
				newBallPos.Y -= inc
			}
			if boxSelect == 1 {
				newBoxPos.Y -= inc
			}
		}
		if rl.IsKeyDown(rl.KeyS) {
			if ballSelect == 1 {
				newBallPos.Y += inc
			}
			if boxSelect == 1 {
				newBoxPos.Y += inc
			}
		}
		if rl.IsKeyDown(rl.KeyD) {
			if ballSelect == 1 {
				newBallPos.X += inc
			}
			if boxSelect == 1 {
				newBoxPos.X += inc
			}
		}
		if rl.IsKeyDown(rl.KeyA) {
			if ballSelect == 1 {
				newBallPos.X -= inc
			}
			if boxSelect == 1 {
				newBoxPos.X -= inc
			}
		}
		ball.SetPosition(newBallPos.X, newBallPos.Y)
		box.SetPosition(newBoxPos.X, newBoxPos.Y)
	})

	game.Run()
}
