package main

import (
	// "fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jakecoffman/cp/v2"
)

const (
	SCREEN_WIDTH  int32 = 600
	SCREEN_HEIGHT int32 = 600
	TARGET_FPS    int32 = 60
)

func main() {
	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "Chipmunk")
	rl.SetTargetFPS(TARGET_FPS)

	// Create Chipmunk space
	space := cp.NewSpace()
	space.Iterations = 20

	// Create two circle bodies and shapes
	mass := 1.0
	radius := 20.0

	var shape *cp.Shape

	// First circle
	body1 := space.AddBody(cp.NewBody(mass, cp.MomentForCircle(mass, 0.0, radius, cp.Vector{})))
	body1.SetPosition(cp.Vector{X: 100, Y: 310})
	body1.SetVelocity(600, 0)
	shape1 := space.AddShape(cp.NewCircle(body1, radius, cp.Vector{}))
	shape1.SetElasticity(0.9)
	shape1.SetFriction(0.7)

	// Second circle
	body2 := space.AddBody(cp.NewBody(mass, cp.MomentForCircle(mass, 0.0, radius, cp.Vector{})))
	body2.SetPosition(cp.Vector{X: 300, Y: 300})
	shape2 := space.AddShape(cp.NewCircle(body2, radius, cp.Vector{}))
	shape2.SetElasticity(0.9)
	shape2.SetFriction(0.7)

	// box
    boxW := radius*2
    boxH := boxW
	body3 := space.AddBody(cp.NewBody(mass, cp.MomentForBox(mass, boxW, boxH)))
	body3.SetPosition(cp.Vector{X: 300, Y: 350})
	shape3 := space.AddShape(cp.NewBox(body3, boxW, boxH, 0))
	shape3.SetElasticity(0.9)
	shape3.SetFriction(0.7)

	// create perimeter wall
	wallThickness := 10.0
	// Create right wall (line segment)
	rightWallBody := cp.NewStaticBody()
	rightWallv1 := cp.Vector{X: float64(SCREEN_WIDTH), Y: 0}
	rightWallv2 := cp.Vector{X: float64(SCREEN_WIDTH), Y: float64(SCREEN_HEIGHT)}
	shape = space.AddShape(cp.NewSegment(rightWallBody, rightWallv1, rightWallv2, wallThickness))
	shape.SetElasticity(1)
	shape.SetFriction(1)

	// Create left wall (line segment)
	leftWallBody := cp.NewStaticBody()
	leftWallv1 := cp.Vector{X: 0, Y: 0}
	leftWallv2 := cp.Vector{X: 0, Y: float64(SCREEN_HEIGHT)}
	shape = space.AddShape(cp.NewSegment(leftWallBody, leftWallv1, leftWallv2, wallThickness))
	shape.SetElasticity(1)
	shape.SetFriction(1)

	// Create bottom wall (line segment)
	bottomWallBody := cp.NewStaticBody()
	bottomWallv1 := cp.Vector{X: 0, Y: float64(SCREEN_HEIGHT)}
	bottomWallv2 := cp.Vector{X: float64(SCREEN_WIDTH), Y: float64(SCREEN_HEIGHT)}
	shape = space.AddShape(cp.NewSegment(bottomWallBody, bottomWallv1, bottomWallv2, wallThickness))
	shape.SetElasticity(1)
	shape.SetFriction(1)

	// Create bottom wall (line segment)
	topWallBody := cp.NewStaticBody()
	topWallv1 := cp.Vector{X: 0, Y: 0}
	topWallv2 := cp.Vector{X: float64(SCREEN_WIDTH), Y: 0}
	shape = space.AddShape(cp.NewSegment(topWallBody, topWallv1, topWallv2, wallThickness))
	shape.SetElasticity(1)
	shape.SetFriction(1)

	var boxRect rl.Rectangle

	for !rl.WindowShouldClose() {
		// Step the physics simulation
		space.Step(1.0 / float64(TARGET_FPS))

		// --------------------- DRAWING ---------------------
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		// Draw the first circle
		lineLength := radius * 0.8 // Length of the line (80% of the radius)
		pos1 := body1.Position()
		rl.DrawCircle(int32(pos1.X), int32(pos1.Y), float32(radius), rl.Blue)
		indicatorX := pos1.X + lineLength*(math.Cos(float64(body1.Angle())))
		indicatorY := pos1.Y + lineLength*(math.Sin(float64(body1.Angle())))
		rl.DrawLine(int32(pos1.X), int32(pos1.Y), int32(indicatorX), int32(indicatorY), rl.White)

		// Draw the second circle
		pos2 := body2.Position()
		rl.DrawCircle(int32(pos2.X), int32(pos2.Y), float32(radius), rl.Red)
		indicatorX = pos2.X + lineLength*(math.Cos(float64(body2.Angle())))
		indicatorY = pos2.Y + lineLength*(math.Sin(float64(body2.Angle())))
		rl.DrawLine(int32(pos2.X), int32(pos2.Y), int32(indicatorX), int32(indicatorY), rl.White)

		// Draw the box
		pos3 := body3.Position()
		angle := float32(body3.Angle()) * 180.0 / math.Pi
		boxRect = rl.NewRectangle(float32(pos3.X), float32(pos3.Y), float32(boxW), float32(boxH))
		rl.DrawRectanglePro(boxRect, rl.NewVector2(boxRect.Width/2, boxRect.Height/2), angle, rl.DarkGreen)

		// Draw the walls (line segments)
		rl.DrawLineEx(CpToRlVector2(rightWallv1), CpToRlVector2(rightWallv2), float32(wallThickness), rl.Black)
		rl.DrawLineEx(CpToRlVector2(leftWallv1), CpToRlVector2(leftWallv2), float32(wallThickness), rl.Black)
		rl.DrawLineEx(CpToRlVector2(bottomWallv1), CpToRlVector2(bottomWallv2), float32(wallThickness), rl.Black)
		rl.DrawLineEx(CpToRlVector2(topWallv1), CpToRlVector2(topWallv2), float32(wallThickness), rl.Black)

		rl.EndDrawing()
		// ---------------------------------------------------
	}

	rl.CloseWindow()
}

func CpToRlVector2(v cp.Vector) rl.Vector2 {
	return rl.NewVector2(float32(v.X), float32(v.Y))
}
