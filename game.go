package engine

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jakecoffman/cp/v2"
	"math"
)

type Game struct {
	screenWidth     int32
	screenHeight    int32
	targetFPS       int32
	windowName      string
	space           *cp.Space
	entities        []Entity
	backgroundColor rl.Color
	updateCallback  func(*Game)
	drawCallback    func(*Game)
}

func NewGame(screenWidth int32, screenHeight int32, targetFPS int32) Game {
	space := cp.NewSpace()
	game := Game{
		screenWidth:     screenWidth,
		screenHeight:    screenHeight,
		targetFPS:       targetFPS,
		windowName:      "Game",
		backgroundColor: rl.RayWhite,
		space:           space,
	}
	rl.InitWindow(game.screenWidth, game.screenHeight, game.windowName)
	rl.SetTargetFPS(game.targetFPS)
	return game
}

func (game Game) Dt() float64 {
	return 1.0 / float64(game.targetFPS)
}

func (game *Game) SetBackgroundColor(color rl.Color) {
    game.backgroundColor = color
}

func (game *Game) SetUpdateCallback(callback func(*Game)) {
	game.updateCallback = callback
}

func (game *Game) SetDrawCallback(callback func(*Game)) {
	game.drawCallback = callback
}

func (game *Game) SetGravity(v Vector2) {
	if game.space != nil {
		game.space.SetGravity(v.ToChipmunk())
	}
}

func (game *Game) SetDamping(d float64) {
	if game.space != nil {
		game.space.SetDamping(d)
	}
}

func (game *Game) Update() {
	game.space.Step(game.Dt())

	for _, ev := range game.entities {
		switch entity := ev.(type) {
		case *Particle:
			if entity.updateCallback != nil {
				entity.updateCallback(entity)
			}
		case *Box:
		case *Wall:
		default:
			fmt.Println("Unknown entity type")
		}
	}
}

func (game Game) Draw() {

	for _, v := range game.entities {
		switch entity := v.(type) {
		case *Particle:
			entity.drawCallback(entity)
		case *Box:
			angle := entity.cpBody.Angle() * 180.0 / math.Pi
			pos := entity.cpBody.Position()
			boxRect := rl.NewRectangle(float32(pos.X), float32(pos.Y), float32(entity.Width), float32(entity.Height))
			rl.DrawRectanglePro(boxRect, rl.NewVector2(boxRect.Width/2, boxRect.Height/2), float32(angle), entity.Color)
		case *Wall:
			if entity.Visible {
				rl.DrawLineEx(entity.Vertex1.ToRaylib(), entity.Vertex2.ToRaylib(), float32(entity.Width), entity.Color)
			}
		default:
			fmt.Println("Unknown entity type")
		}
	}

}

func (game *Game) Run() {

	for !rl.WindowShouldClose() {
		game.Update()
		if game.updateCallback != nil {
			game.updateCallback(game)
		}

		// ---------- Drawing ----------
		rl.BeginDrawing()
		rl.ClearBackground(game.backgroundColor)
		game.Draw()
		if game.drawCallback != nil {
			game.drawCallback(game)
		}
		rl.EndDrawing()
		// -----------------------------
	}

	rl.CloseWindow()
}
// Limit velocity to the specified max speed
func (p Particle) limitVelocity(body *cp.Body, gravity cp.Vector, damping float64, dt float64) {
    maxSpeed := p.velocityMax // Maximum speed (pixels/second)
	cp.BodyUpdateVelocity(body, gravity, damping, dt)
	velocity := body.Velocity()
	speed := math.Sqrt(velocity.X*velocity.X + velocity.Y*velocity.Y)
	if speed > maxSpeed {
		scale := maxSpeed / speed
		body.SetVelocity(velocity.X*scale, velocity.Y*scale)
	}
}

func (game *Game) AddEntity(entity Entity) {
	var body *cp.Body
	var shape *cp.Shape

	switch e := entity.(type) {
	case *Particle:
		body = game.space.AddBody(cp.NewBody(e.mass, cp.MomentForCircle(e.mass, 0.0, e.radius, cp.Vector{})))
		body.SetType(cp.BODY_DYNAMIC)
		body.SetPosition(cp.Vector{X: e.position.X, Y: e.position.Y})
		body.SetVelocity(e.velocity.X, e.velocity.Y)
		shape = game.space.AddShape(cp.NewCircle(body, e.radius, cp.Vector{}))
		shape.SetElasticity(e.elasticity)
		shape.SetFriction(e.friction)

        body.SetVelocityUpdateFunc(e.limitVelocity)

		e.id = uint64(len(game.entities))
		e.cpBody = body
		e.cpShape = shape
		game.entities = append(game.entities, e)
	case *Box:
		body = game.space.AddBody(cp.NewBody(e.Mass, cp.MomentForBox(e.Mass, e.Width, e.Height)))
		body.SetPosition(cp.Vector{X: e.position.X, Y: e.position.Y})
		body.SetVelocity(e.velocity.X, e.velocity.Y)
		shape = game.space.AddShape(cp.NewBox(body, e.Width, e.Height, 0))
		shape.SetElasticity(e.elasticity)
		shape.SetFriction(e.friction)

		e.id = uint64(len(game.entities))
		e.cpBody = body
		e.cpShape = shape
		game.entities = append(game.entities, e)
	case *Wall:
		body = cp.NewStaticBody()
		shape = game.space.AddShape(cp.NewSegment(body, cp.Vector{X: e.Vertex1.X, Y: e.Vertex1.Y}, cp.Vector{X: e.Vertex2.X, Y: e.Vertex2.Y}, e.Width/2))
		shape.SetElasticity(1)
		shape.SetFriction(1)

		e.id = uint64(len(game.entities))
		e.cpBody = body
		// e.cpShape = shape
		game.entities = append(game.entities, e)
	default:
		fmt.Println("Unknown entity type")
		return
	}
}

type Vector2 struct {
	X float64
	Y float64
}

func NewVector2(x float64, y float64) Vector2 {
	return Vector2{X: x, Y: y}
}

func (v Vector2) ToRaylib() rl.Vector2 {
	return rl.NewVector2(float32(v.X), float32(v.Y))
}

func (v Vector2) ToChipmunk() cp.Vector {
	return cp.Vector{X: float64(v.X), Y: float64(v.Y)}
}
