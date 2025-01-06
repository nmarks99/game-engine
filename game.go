package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jakecoffman/cp/v2"
)

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

type Game struct {
	space    *cp.Space
	bodies   []*cp.Body
	entities []interface{}
}

func NewGame() Game {
	space := cp.NewSpace()
	return Game{
		space: space,
	}
}

func (game *Game) Update() {
	const TARGET_FPS float64 = 60.0
	game.space.Step(1.0 / TARGET_FPS)
}

func (game Game) Draw() {

	for i, v := range game.entities {
		switch entity := v.(type) {
		case *Particle:
			pos := game.bodies[i].Position()
			rl.DrawCircle(int32(pos.X), int32(pos.Y), float32(entity.Radius), entity.Color)
		default:
			fmt.Println("Unknown entity type")
		}
	}

}

type Particle struct {
	Position   Vector2
	Velocity   Vector2
	Radius     float64
	Mass       float64
	Elasticity float64
	Friction   float64
	Color      rl.Color
	id         uint64
	cpBody     *cp.Body
}

func NewParticle(position Vector2, radius float64, mass float64, color rl.Color) Particle {
	return Particle{
		Position:   position,
		Radius:     radius,
		Mass:       mass,
		Color:      color,
		Elasticity: 1.0,
		Friction:   1.0,
	}
}

func (p *Particle) SetVelocity(x float64, y float64) {
	p.Velocity.X = x
	p.Velocity.Y = y

	if p.cpBody != nil {
		p.cpBody.SetVelocity(x, y)
	}
}

func (p *Particle) SetPosition(x float64, y float64) {
	p.Position.X = x
	p.Position.Y = y

	if p.cpBody != nil {
		p.cpBody.SetPosition(p.Position.ToChipmunk())
	}
}

func (game *Game) AddParticle(p *Particle) {
	var body *cp.Body
	var shape *cp.Shape

	body = game.space.AddBody(cp.NewBody(p.Mass, cp.MomentForCircle(p.Mass, 0.0, p.Radius, cp.Vector{})))
	body.SetPosition(cp.Vector{X: p.Position.X, Y: p.Position.Y})
	body.SetVelocity(p.Velocity.X, p.Velocity.Y)
	shape = game.space.AddShape(cp.NewCircle(body, p.Radius, cp.Vector{}))
	shape.SetElasticity(p.Elasticity)
	shape.SetFriction(p.Friction)

	p.id = uint64(len(game.bodies))
	p.cpBody = body

	game.bodies = append(game.bodies, body)
	game.entities = append(game.entities, p)
}
