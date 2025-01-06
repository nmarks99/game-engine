package engine

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jakecoffman/cp/v2"
	"math"
)

type Entity interface{}

type Game struct {
	screenWidth    int32
	screenHeight   int32
	targetFPS      int32
	windowName     string
	space          *cp.Space
	bodies         []*cp.Body
	entities       []Entity
	updateCallback func(*Game)
	drawCallback   func(*Game)
}

func NewGame(screenWidth int32, screenHeight int32, targetFPS int32) Game {
	space := cp.NewSpace()
	game := Game{
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
		targetFPS:    targetFPS,
		windowName:   "Game",
		space:        space,
	}
	rl.InitWindow(game.screenWidth, game.screenHeight, game.windowName)
	rl.SetTargetFPS(game.targetFPS)
	return game
}

func (game *Game) SetUpdateCallback(callback func(*Game)) {
	game.updateCallback = callback
}

func (game *Game) SetDrawCallback(callback func(*Game)) {
	game.drawCallback = callback
}

func (game *Game) Update() {
	game.space.Step(1.0 / float64(game.targetFPS))
}

func (game Game) Draw() {

	for i, v := range game.entities {
		pos := game.bodies[i].Position()
		switch entity := v.(type) {
		case *Particle:
			rl.DrawCircle(int32(pos.X), int32(pos.Y), float32(entity.Radius), entity.Color)
		case *Box:
			angle := game.bodies[i].Angle() * 180.0 / math.Pi
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
		rl.ClearBackground(rl.RayWhite)
		game.Draw()
		if game.drawCallback != nil {
			game.drawCallback(game)
		}
		rl.EndDrawing()
		// -----------------------------
	}

	rl.CloseWindow()
}

func (game *Game) AddEntity(entity Entity) {
	var body *cp.Body
	var shape *cp.Shape

	switch e := entity.(type) {
	case *Particle:
		body = game.space.AddBody(cp.NewBody(e.Mass, cp.MomentForCircle(e.Mass, 0.0, e.Radius, cp.Vector{})))
		body.SetPosition(cp.Vector{X: e.Position.X, Y: e.Position.Y})
		body.SetVelocity(e.Velocity.X, e.Velocity.Y)
		shape = game.space.AddShape(cp.NewCircle(body, e.Radius, cp.Vector{}))
		shape.SetElasticity(e.Elasticity)
		shape.SetFriction(e.Friction)

		e.id = uint64(len(game.bodies))
		e.cpBody = body
		game.bodies = append(game.bodies, body)
		game.entities = append(game.entities, e)
	case *Box:
		body = game.space.AddBody(cp.NewBody(e.Mass, cp.MomentForBox(e.Mass, e.Width, e.Height)))
		body.SetPosition(cp.Vector{X: e.Position.X, Y: e.Position.Y})
		shape = game.space.AddShape(cp.NewBox(body, e.Width, e.Height, 0))
		shape.SetElasticity(e.Elasticity)
		shape.SetFriction(e.Friction)

		e.id = uint64(len(game.bodies))
		e.cpBody = body
		game.bodies = append(game.bodies, body)
		game.entities = append(game.entities, e)
	case *Wall:
		body = cp.NewStaticBody()
		shape = game.space.AddShape(cp.NewSegment(body, cp.Vector{X: e.Vertex1.X, Y: e.Vertex1.Y}, cp.Vector{X: e.Vertex2.X, Y: e.Vertex2.Y}, e.Width/2))
		shape.SetElasticity(1)
		shape.SetFriction(1)

		e.id = uint64(len(game.bodies))
		e.cpBody = body
		game.bodies = append(game.bodies, body)
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
