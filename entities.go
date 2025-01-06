package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jakecoffman/cp/v2"
)

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

type Box struct {
	Position   Vector2
	Velocity   Vector2
	Width      float64
	Height     float64
	Mass       float64
	Elasticity float64
	Friction   float64
	Color      rl.Color
	id         uint64
	cpBody     *cp.Body
}

func NewBox(position Vector2, width float64, height float64, mass float64, color rl.Color) Box {
	return Box{
		Position:   position,
		Width:      width,
		Height:     height,
		Mass:       mass,
		Color:      color,
		Elasticity: 1.0,
		Friction:   1.0,
	}
}

type Wall struct {
	Vertex1 Vector2
	Vertex2 Vector2
	Width   float64
	Color   rl.Color
	Visible bool
	id      uint64
	cpBody  *cp.Body
}

func NewWall(vertex1 Vector2, vertex2 Vector2, width float64, color rl.Color) Wall {
	return Wall{
		Vertex1: vertex1,
		Vertex2: vertex2,
		Width:   width,
		Color:   color,
		Visible: true,
	}
}
