package engine

import (
    "math"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jakecoffman/cp/v2"
	// "fmt"
)

type Entity interface{}

type Particle struct {
	position       Vector2
	velocity       Vector2
	Radius         float64
	Mass           float64
	Elasticity     float64
	Friction       float64
	Color          rl.Color
	updateCallback func(*Particle)
	drawCallback   func(*Particle)
	id             uint64
	cpBody         *cp.Body
	cpShape        *cp.Shape
}

func (p *Particle) SetUpdateCallback(callback func(*Particle)) {
	p.updateCallback = callback
}

func (p *Particle) SetDrawCallback(callback func(*Particle)) {
	p.drawCallback = callback
}

func (p *Particle) Fix() {
    if p.cpBody != nil {
        p.SetVelocity(0, 0)
        p.cpBody.SetMass(math.Inf(1))
    }
}

func (p *Particle) Unfix() {
    if p.cpBody != nil {
        p.cpBody.SetMass(p.Mass)
    }
}

func (p Particle) Position() Vector2 {
	if p.cpBody != nil {
		return Vector2{X: p.cpBody.Position().X, Y: p.cpBody.Position().Y}
	} else {
		return Vector2{X: p.Position().X, Y: p.Position().Y}
	}
}

func (p Particle) Velocity() Vector2 {
	if p.cpBody != nil {
		return Vector2{X: p.cpBody.Velocity().X, Y: p.cpBody.Velocity().Y}
	} else {
		return Vector2{X: p.Velocity().X, Y: p.Velocity().Y}
	}
}

func (p *Particle) SetVelocity(x float64, y float64) {
	p.velocity.X = x
	p.velocity.Y = y

	if p.cpBody != nil {
		p.cpBody.SetVelocity(x, y)
	}
}

func (p *Particle) SetPosition(x float64, y float64) {
	p.position.X = x
	p.position.Y = y

	if p.cpBody != nil {
		p.cpBody.SetPosition(p.Position().ToChipmunk())
	}
}

func defaultDrawCallback(p *Particle) {
	pos := p.Position()
	rl.DrawCircle(int32(pos.X), int32(pos.Y), float32(p.Radius), p.Color)
}

func NewParticle(position Vector2, radius float64, mass float64, color rl.Color) Particle {
	pOut := Particle{
		position:   position,
		Radius:     radius,
		Mass:       mass,
		Color:      color,
		Elasticity: 1.0,
		Friction:   1.0,
	}
	pOut.SetDrawCallback(defaultDrawCallback)
	return pOut
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
